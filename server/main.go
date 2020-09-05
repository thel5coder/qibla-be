package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	redis "github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/skilld-labs/go-odoo"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
	"qibla-backend/db"
	awsHelper "qibla-backend/helpers/aws"
	"qibla-backend/helpers/jwe"
	"qibla-backend/helpers/jwt"
	"qibla-backend/helpers/mailing"
	"qibla-backend/helpers/pusher"
	redisHelper "qibla-backend/helpers/redis"
	"qibla-backend/helpers/str"
	"qibla-backend/server/bootstrap"
	"qibla-backend/server/middleware"
	"qibla-backend/usecase"
)

var (
	validatorDriver *validator.Validate
	uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error loading ..env file")
	}

	//jwe
	jweCredential := jwe.Credential{
		KeyLocation: os.Getenv("PRIVATE_KEY"),
		Passphrase:  os.Getenv("PASSPHRASE"),
	}

	//setup redis
	redisOption := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	redisClient := redisHelper.RedisClient{Client: redis.NewClient(redisOption)}

	//jwtconfig
	jwtConfig := echoMiddleware.JWTConfig{
		Claims:     &jwt.CustomClaims{},
		SigningKey: []byte(os.Getenv("SECRET")),
	}

	//jwt credential
	jwtCred := jwt.JwtCredential{
		TokenSecret:         os.Getenv("SECRET"),
		ExpiredToken:        str.StringToInt(os.Getenv("TOKEN_EXP_TIME")),
		RefreshTokenSecret:  os.Getenv("SECRET_REFRESH_TOKEN"),
		ExpiredRefreshToken: str.StringToInt(os.Getenv("REFRESH_TOKEN_EXP_TIME")),
	}

	//setup db connection
	dbInfo := db.Connection{
		Host:     os.Getenv("DB_HOST"),
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}

	database, err := dbInfo.DbConnect()
	if err != nil {
		panic(err)
	}

	pong, err := redisClient.Client.Ping().Result()
	fmt.Println("Redis ping status: "+pong, err)

	//odo
	c, err := odoo.NewClient(&odoo.ClientConfig{
		Admin:    "admin",
		Password: "admin",
		Database: "him",
		URL:      "https://demo.garudea.com",
	})

	//aws setup
	awsAccessKey := os.Getenv("S3_ACCESS_KEY")
	awsSecretKey := os.Getenv("S3_SECRET_KEY")
	awsBucket := os.Getenv("S3_BUCKET")
	awsDirectory := os.Getenv("S3_DIRECTORY")
	s3EndPoint := os.Getenv("S3_BASE_URL")
	awsConfig := aws.Config{
		Endpoint:    &s3EndPoint,
		Region:      aws.String(os.Getenv("S3_LOCATION")),
		Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
	}
	awsS3 := awsHelper.AWSS3{
		AWSConfig: awsConfig,
		Bucket:    awsBucket,
		Directory: awsDirectory,
		AccessKey: awsAccessKey,
		SecretKey: awsSecretKey,
	}

	//pusher
	pusherCredential := pusher.Credential{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
	}

	//gomail config
	goMailConfig := mailing.GoMailConfig{
		SMTPHost: os.Getenv("SMTP_HOST"),
		SMTPPort: str.StringToInt(os.Getenv("SMTP_PORT")),
		Sender:   os.Getenv("MAIL_SENDER"),
		Password: os.Getenv("PASSWORD"),
	}

	//youtube
	client := &http.Client{
		Transport: &transport.APIKey{Key:"AIzaSyCJ6UN03_evvGzIGfMKK0yRyB-L9JluE0k"},
	}

	youtubeClient, err := youtube.New(client)
	if err != nil {
		log.Fatal(err.Error())
	}

	//init validator
	validatorInit()

	e := echo.New()

	ucContract := usecase.UcContract{
		E:             e,
		DB:            database,
		RedisClient:   redisClient,
		Jwe:           jweCredential,
		Validate:      validatorDriver,
		Translator:    translator,
		JwtConfig:     jwtConfig,
		JwtCred:       jwtCred,
		Odoo:          c,
		AWSS3:         awsS3,
		Pusher:        pusherCredential,
		GoMailConfig:  goMailConfig,
		YoutubeClient: youtubeClient,
	}

	bootApp := bootstrap.Bootstrap{
		E:               e,
		UseCaseContract: ucContract,
		Jwe:             jweCredential,
		Translator:      translator,
		Validator:       validatorDriver,
		JwtConfig:       jwtConfig,
		JwtCred:         jwtCred,
		Odoo:            c,
	}

	bootApp.E.Static("/", "")
	//bootApp.E.Use(echoMiddleware.Static("statics"))
	bootApp.E.Use(echoMiddleware.Logger())
	bootApp.E.Use(echoMiddleware.Recover())
	bootApp.E.Use(echoMiddleware.CORS())
	bootApp.E.Use(middleware.HeaderXRequestID())

	bootApp.RegisterRouters()
	bootApp.E.Logger.Fatal(bootApp.E.Start(os.Getenv("APP_HOST_SERVER")))
}

func validatorInit() {
	en := en.New()
	id := id.New()
	uni = ut.New(en, id)

	transEN, _ := uni.GetTranslator("en")
	transID, _ := uni.GetTranslator("id")

	validatorDriver = validator.New()

	enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	idTranslations.RegisterDefaultTranslations(validatorDriver, transID)

	switch os.Getenv("APP_LOCALE") {
	case "en":
		translator = transEN
	case "id":
		translator = transID
	}
}
