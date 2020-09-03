package usecase

import (
	"database/sql"
	"errors"
	"github.com/skilld-labs/go-odoo"
	"math/rand"
	"os"
	queue "qibla-backend/helpers/amqp"
	"qibla-backend/helpers/aws"
	"qibla-backend/helpers/jwe"
	"qibla-backend/helpers/jwt"
	"qibla-backend/helpers/mailing"
	"qibla-backend/helpers/messages"
	"qibla-backend/helpers/pusher"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"

	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/streadway/amqp"
	"qibla-backend/helpers/redis"
)

const (
	defaultLimit      = 10
	maxLimit          = 50
	defaultOrderBy    = "id"
	defaultSort       = "asc"
	PasswordLength    = 6
	defaultLastPage   = 0
	OtpLifeTime       = "3m"
	MaxOtpSubmitRetry = 3
	StaticBaseUrl     = "/statics"
	fasPayBaseUrl     = "https://dev.faspay.co.id/cvr"
)

//globalsmscounter
var GlobalSmsCounter int

// AmqpConnection ...
var AmqpConnection *amqp.Connection

// AmqpChannel ...
var AmqpChannel *amqp.Channel

//X-Request-ID
var xRequestID interface{}

type UcContract struct {
	E            *echo.Echo
	DB           *sql.DB
	TX           *sql.Tx
	RedisClient  redis.RedisClient
	Jwe          jwe.Credential
	Validate     *validator.Validate
	Translator   ut.Translator
	JwtConfig    middleware.JWTConfig
	JwtCred      jwt.JwtCredential
	Odoo         *odoo.Client
	AWSS3        aws.AWSS3
	Pusher       pusher.Credential
	GoMailConfig mailing.GoMailConfig
}

func (uc UcContract) setPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	if order == "" {
		order = defaultOrderBy
	}

	if sort == "" {
		sort = defaultSort
	}

	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc UcContract) setPaginationResponse(page, limit, total int) (paginationResponse viewmodel.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	paginationResponse = viewmodel.PaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}

	return paginationResponse
}

func (uc UcContract) GetRandomString(length int) string {
	if length == 0 {
		length = PasswordLength
	}

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	password := b.String()

	return password
}

func (uc UcContract) LimitRetryByKey(key string, limit float64) (err error) {
	var count float64
	res := map[string]interface{}{}

	err = uc.RedisClient.GetFromRedis(key, &res)
	if err != nil {
		err = nil
		res = map[string]interface{}{
			"counter": count,
		}
	}
	count = res["counter"].(float64) + 1
	if count > limit {
		uc.RedisClient.RemoveFromRedis(key)

		return errors.New(messages.MaxRetryKey)
	}

	res["counter"] = count
	uc.RedisClient.StoreToRedistWithExpired(key, res, "24h")

	return err
}

func (uc UcContract) SetXRequestID(ctx echo.Context) {
	xRequestID = ctx.Get(echo.HeaderXRequestID)
}

func (uc UcContract) GetXRequestID() interface{} {
	return xRequestID
}

func (uc UcContract) PushToQueue(queueBody map[string]interface{}, queueType, deadLetterType string) (err error) {
	mqueue := queue.NewQueue(AmqpConnection, AmqpChannel)

	_, _, err = mqueue.PushQueueReconnect(os.Getenv("AMQP_URL"), queueBody, queueType, deadLetterType)
	if err != nil {
		return err
	}

	return err
}

func (uc UcContract) Read(object string, criteria *odoo.Criteria, options *odoo.Options, res interface{}) (err error) {
	err = uc.Odoo.SearchRead(object, criteria, odoo.NewOptions().Limit(1), res)
	if err != nil {
		return err
	}

	return err
}
