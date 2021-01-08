package usecase

import (
	"database/sql"
	"errors"
	"github.com/skilld-labs/go-odoo"
	"google.golang.org/api/youtube/v3"
	"math/rand"
	"os"
	"qibla-backend/pkg/aes"
	queue "qibla-backend/pkg/amqp"
	"qibla-backend/pkg/aws"
	"qibla-backend/pkg/fcm"
	"qibla-backend/pkg/flip"
	"qibla-backend/pkg/jwe"
	"qibla-backend/pkg/jwt"
	"qibla-backend/pkg/mailing"
	"qibla-backend/pkg/messages"
	"qibla-backend/pkg/pusher"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"

	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/streadway/amqp"
	"qibla-backend/pkg/redis"
)

const (
	defaultLimit   = 10
	maxLimit       = 50
	defaultOrderBy = "id"
	defaultSort    = "asc"
	// PasswordLength ...
	PasswordLength  = 6
	defaultLastPage = 0
	// OtpLifeTime ...
	OtpLifeTime = "3m"
	// MaxOtpSubmitRetry ...
	MaxOtpSubmitRetry = 3
	// StaticBaseURL ...
	StaticBaseURL               = "/statics"
	defaultMaxResultYoutubeData = 3
	defaultYoutubeSearchType    = "video"
	defaultYoutubeOrder         = "date"
	defaultFaspayTerminal       = 10
	defaultFaspayPayType        = 1
	defaultFaspayTenor          = "00"
	defaultFaspayPaymentPlan    = "01"
	defaultFaspayCurrency       = "IDR"
	defaultInvoiceDueDate       = 7
	defaultFeeZakat             = 5000

	// DefaultOriginAccountNumber ...
	DefaultOriginAccountNumber = "111"
	// DefaultOriginAccountName ...
	DefaultOriginAccountName = "111"
	// DefaultOriginAccountBankName ...
	DefaultOriginAccountBankName = "111"
	// DefaultOriginAccountBankCode ...
	DefaultOriginAccountBankCode = "111"

	// DefaultLocation ...
	DefaultLocation = "Asia/Jakarta"
)

// GlobalSmsCounter ...
var GlobalSmsCounter int

// AmqpConnection ...
var AmqpConnection *amqp.Connection

// AmqpChannel ...
var AmqpChannel *amqp.Channel

//X-Request-ID
var xRequestID interface{}

// UcContract ...
type UcContract struct {
	ReqID          string
	E              *echo.Echo
	DB             *sql.DB
	TX             *sql.Tx
	AmqpConn       *amqp.Connection
	AmqpChannel    *amqp.Channel
	RedisClient    redis.RedisClient
	Jwe            jwe.Credential
	Validate       *validator.Validate
	Translator     ut.Translator
	JwtConfig      middleware.JWTConfig
	JwtCred        jwt.JwtCredential
	Odoo           *odoo.Client
	AWSS3          aws.AWSS3
	Pusher         pusher.Credential
	GoMailConfig   mailing.GoMailConfig
	YoutubeService *youtube.Service
	UserID         string
	Fcm            fcm.Connection
	OdooDBConn     *sql.DB
	Flip           flip.Credential
	AES            aes.Credential
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

// GetRandomString ...
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

// LimitRetryByKey ...
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

// SetXRequestID ...
func (uc UcContract) SetXRequestID(ctx echo.Context) {
	xRequestID = ctx.Get(echo.HeaderXRequestID)
}

// GetXRequestID ...
func (uc UcContract) GetXRequestID() interface{} {
	return xRequestID
}

// PushToQueue ...
func (uc UcContract) PushToQueue(queueBody map[string]interface{}, queueType, deadLetterType string) (err error) {
	mqueue := queue.NewQueue(AmqpConnection, AmqpChannel)

	_, _, err = mqueue.PushQueueReconnect(os.Getenv("AMQP_URL"), queueBody, queueType, deadLetterType)
	if err != nil {
		return err
	}

	return err
}

// Read ...
func (uc UcContract) Read(object string, criteria *odoo.Criteria, options *odoo.Options, res interface{}) (err error) {
	err = uc.Odoo.SearchRead(object, criteria, odoo.NewOptions().Limit(1), res)
	if err != nil {
		return err
	}

	return err
}

// InitDBTransaction ...
func (uc UcContract) InitDBTransaction() (err error) {
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	return nil
}

//code from aes
func (uc UcContract) GenerateKeyFromAES(email string) (res string, err error) {
	res, err = uc.AES.Encrypt(email)
	if err != nil {
		return res, err
	}

	redisBody := map[string]interface{}{
		"email": email,
	}
	err = uc.RedisClient.StoreToRedistWithExpired("code-"+res, redisBody, "24h")
	if err != nil {
		return res, err
	}

	return res, nil
}
