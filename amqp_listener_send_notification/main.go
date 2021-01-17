package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"runtime"
	"strconv"

	"qibla-backend/db"
	amqpPkg "qibla-backend/pkg/amqp"
	"qibla-backend/pkg/amqpconsumer"
	awsHelper "qibla-backend/pkg/aws"
	"qibla-backend/pkg/fcm"
	"qibla-backend/pkg/google"
	"qibla-backend/pkg/interfacepkg"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/mailing"
	"qibla-backend/pkg/pusher"
	redisHelper "qibla-backend/pkg/redis"
	"qibla-backend/pkg/str"
	"qibla-backend/usecase"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"github.com/skilld-labs/go-odoo"
	"github.com/streadway/amqp"
	"google.golang.org/api/youtube/v3"
)

var (
	uri          *string
	formURL      = flag.String("form_url", "http://localhost", "The URL that requests are sent to")
	logFile      = flag.String("log_file", "system.log", "The file where errors are logged")
	threads      = flag.Int("threads", 1, "The max amount of go routines that you would like the process to use")
	maxprocs     = flag.Int("max_procs", 1, "The max amount of processors that your application should use")
	paymentsKey  = flag.String("payments_key", "secret", "Access key")
	exchange     = flag.String("exchange", amqpPkg.SendNotificationExchange, "The exchange we will be binding to")
	exchangeType = flag.String("exchange_type", "direct", "Type of exchange we are binding to | topic | direct| etc..")
	queue        = flag.String("queue", amqpPkg.SendNotification, "Name of the queue that you would like to connect to")
	routingKey   = flag.String("routing_key", amqpPkg.SendNotificationDeadLetter, "queue to route messages to")
	workerName   = flag.String("worker_name", "worker.name", "name to identify worker by")
	verbosity    = flag.Bool("verbos", false, "Set true if you would like to log EVERYTHING")

	// Hold consumer so our go routine can listen to
	// it's done error chan and trigger reconnects
	// if it's ever returned
	conn *amqpconsumer.Consumer
)

func init() {
	flag.Parse()
	runtime.GOMAXPROCS(*maxprocs)
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	uri = flag.String("uri", os.Getenv("AMQP_URL"), "The rabbitmq endpoint")
}

func main() {
	file := false
	if file {
		f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		log.SetOutput(f)
	}

	conn := amqpconsumer.NewConsumer(*workerName, *uri, *exchange, *exchangeType, *queue)

	if err := conn.Connect(); err != nil {
		panic(err)
	}

	deliveries, err := conn.AnnounceQueue(*queue, *routingKey)
	if err != nil {
		panic(err)
	}

	//setup redis
	redisOption := &redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}
	redisClient := redisHelper.RedisClient{Client: redis.NewClient(redisOption)}

	//setup db connection
	dbInfo := db.Connection{
		Host:     os.Getenv("DB_HOST"),
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
		MaxConnection:         str.StringToInt(os.Getenv("DB_MAX_CONNECTION")),
		MaxIdleConnection:     str.StringToInt(os.Getenv("DB_MAX_IDLE_CONNECTION")),
		MaxLifeTimeConnection: str.StringToInt(os.Getenv("DB_MAX_LIFETIME_CONNECTION")),
	}

	database, err := dbInfo.DbConnect()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	_, err = redisClient.Client.Ping().Result()
	if err != nil {
		panic(err)
	}

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

	ucContract := usecase.UcContract{
		DB:             database,
		RedisClient:    redisClient,
		Odoo:           c,
		AWSS3:          awsS3,
		Pusher:         pusherCredential,
		GoMailConfig:   goMailConfig,
		YoutubeService: youtubeService,
		Fcm:            fcmConnection,
	}

	conn.Handle(deliveries, handler, *threads, *queue, *routingKey, ucContract)
}

func handler(deliveries <-chan amqp.Delivery, uc *usecase.UcContract) {
	ctx := "SendNotificationListener"
	for d := range deliveries {
		var formData map[string]interface{}

		err := json.Unmarshal(d.Body, &formData)
		if err != nil {
			log.Printf("Error unmarshaling data: %s", err.Error())
		}

		logruslogger.Log(logruslogger.InfoLevel, interfacepkg.Marshall(formData), ctx, "begin", formData["qid"].(string))

		uc.ReqID = formData["qid"].(string)
		fcmUc := usecase.FcmUseCase{UcContract: uc}
		res, err := fcmUc.Send(
			[]string{formData["fcm_device_token"].(string)}, formData["title"].(string), formData["message"].(string),
			map[string]interface{}{},
		)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "err", formData["qid"].(string))

			// Get fail counter from redis
			failCounter := amqpconsumer.FailCounter{}
			err = uc.RedisClient.GetFromRedis("amqpFail"+formData["qid"].(string), &failCounter)
			if err != nil {
				failCounter = amqpconsumer.FailCounter{
					Counter: 1,
				}
			}

			if failCounter.Counter > amqpconsumer.MaxFailCounter {
				logruslogger.Log(logruslogger.WarnLevel, strconv.Itoa(failCounter.Counter), ctx, "rejected", formData["qid"].(string))
				d.Reject(false)
			} else {
				// Save the new counter to redis
				failCounter.Counter++
				err = uc.RedisClient.StoreToRedistWithExpired("amqpFail"+formData["qid"].(string), failCounter, "10m")

				logruslogger.Log(logruslogger.WarnLevel, strconv.Itoa(failCounter.Counter), ctx, "failed", formData["qid"].(string))
				d.Nack(false, true)
			}
		} else {
			logruslogger.Log(logruslogger.InfoLevel, interfacepkg.Marshall(res), ctx, "success", formData["qid"].(string))
			d.Ack(false)
		}
	}

	return
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
