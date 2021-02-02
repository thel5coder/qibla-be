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
	"github.com/rs/xid"
	"github.com/skilld-labs/go-odoo"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
	"qibla-backend/db"
	"qibla-backend/pkg/aes"
	"qibla-backend/pkg/amqp"
	awsHelper "qibla-backend/pkg/aws"
	"qibla-backend/pkg/env"
	"qibla-backend/pkg/fcm"
	"qibla-backend/pkg/flip"
	"qibla-backend/pkg/google"
	"qibla-backend/pkg/jwe"
	"qibla-backend/pkg/jwt"
	"qibla-backend/pkg/mailing"
	"qibla-backend/pkg/pusher"
	redisHelper "qibla-backend/pkg/redis"
	"qibla-backend/pkg/str"
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

	// Load env into map
	envConfig := env.NewEnvConfig("../.env")

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
		Host:                  os.Getenv("DB_HOST"),
		DbName:                os.Getenv("DB_NAME"),
		User:                  os.Getenv("DB_USER"),
		Password:              os.Getenv("DB_PASSWORD"),
		Port:                  os.Getenv("DB_PORT"),
		SslMode:               os.Getenv("DB_SSL_MODE"),
		MaxConnection:         str.StringToInt(os.Getenv("DB_MAX_CONNECTION")),
		MaxIdleConnection:     str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		MaxLifeTimeConnection: str.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")),
	}
	database, err := dbInfo.DbConnect()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	pong, err := redisClient.Client.Ping().Result()
	fmt.Println("Redis ping status: "+pong, err)

	//odo
	c, err := odoo.NewClient(&odoo.ClientConfig{
		Admin:    "admin",
		Password: "admin",
		Database: "Demo",
		URL:      "https://staging.qibla.co.id",
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
	youtubeCred := google.YoutubeCred{
		TokenFile:  os.Getenv("YOUTUBE_TOKEN_FILE"),
		SecretFile: os.Getenv("YOUTUBE_SECRET_FILE"),
		Scope:      youtube.YoutubeReadonlyScope,
	}
	youtubeService, err := youtubeCred.GetYoutubeService()

	// FCM connection
	fcmConnection := fcm.Connection{
		APIKey: os.Getenv("FMC_KEY"),
	}

	// Mqueue connection
	amqpInfo := amqp.Connection{
		URL: os.Getenv("AMQP_URL"),
	}
	amqpConn, amqpChannel, err := amqpInfo.Connect()
	if err != nil {
		panic(err)
	}
	usecase.AmqpConnection = amqpConn
	usecase.AmqpChannel = amqpChannel

	// Flip
	flipCredential := flip.Credential{
		BaseURL:         os.Getenv("FLIP_BASE_URL"),
		SecretKey:       os.Getenv("FLIP_SECRET_KEY"),
		ValidationToken: envConfig["FLIP_VALIDATION_TOKEN"],
	}

	//aes
	aesCred := aes.Credential{Key: os.Getenv("AES_KEY")}

	//init validator
	validatorInit()

	e := echo.New()

	ucContract := usecase.UcContract{
		ReqID:          xid.New().String(),
		E:              e,
		TX:             nil,
		DB:             database,
		RedisClient:    redisClient,
		Jwe:            jweCredential,
		Validate:       validatorDriver,
		Translator:     translator,
		JwtConfig:      jwtConfig,
		JwtCred:        jwtCred,
		Odoo:           c,
		AWSS3:          awsS3,
		Pusher:         pusherCredential,
		GoMailConfig:   goMailConfig,
		YoutubeService: youtubeService,
		Fcm:            fcmConnection,
		AmqpConn:       amqpConn,
		AmqpChannel:    amqpChannel,
		Flip:           flipCredential,
		AES:            aesCred,
	}

	bootApp := bootstrap.Bootstrap{
		E:               e,
		Db:              database,
		UseCaseContract: ucContract,
		Jwe:             jweCredential,
		Translator:      translator,
		Validator:       validatorDriver,
		JwtConfig:       jwtConfig,
		JwtCred:         jwtCred,
		Odoo:            c,
	}

	bootApp.E.Static("/", "")
	bootApp.E.Use(echoMiddleware.Static("statics"))
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
