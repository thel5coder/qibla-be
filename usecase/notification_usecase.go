package usecase

import (
	"os"
	"qibla-backend/pkg/amqp"
	"qibla-backend/pkg/logruslogger"
)

// NotificationUseCase ...
type NotificationUseCase struct {
	*UcContract
}

// Send ...
func (uc NotificationUseCase) Send(userID, title, message string) (err error) {
	ctx := "NotificationUseCase.Send"

	userUc := UserUseCase{UcContract: uc.UcContract}
	user, err := userUc.ReadBy("u.id", userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_user", uc.ReqID)
		return err
	}

	mqueue := amqp.NewQueue(AmqpConnection, AmqpChannel)
	queueBody := map[string]interface{}{
		"qid":              uc.UcContract.ReqID,
		"fcm_device_token": user.FcmDeviceToken,
		"title":            title,
		"message":          message,
	}
	AmqpConnection, AmqpChannel, err = mqueue.PushQueueReconnect(os.Getenv("AMQP_URL"), queueBody, amqp.SendNotification, amqp.SendNotificationDeadLetter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "amqp", uc.ReqID)
		return err
	}

	return err
}
