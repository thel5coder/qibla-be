package usecase

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"qibla-backend/helpers/amqp"
	"qibla-backend/helpers/logruslogger"
	timepkg "qibla-backend/helpers/time"
)

// CronjobUseCase ...
type CronjobUseCase struct {
	*UcContract
}

// Test ...
func (uc CronjobUseCase) Test() {
	now := time.Now().UTC()
	date := timepkg.InFormatNoErr(now, DefaultLocation, "2006-01-02")

	fmt.Println(date)
}

// DisbursementMutation ...
func (uc CronjobUseCase) DisbursementMutation() {
	ctx := "CronjobUseCase.DisbursementMutation"
	now := time.Now().UTC()
	date := timepkg.InFormatNoErr(now, DefaultLocation, "2006-01-02")
	success := 0
	failed := 0

	logruslogger.Log(logruslogger.InfoLevel, date, ctx, "begin", uc.ReqID)

	contactUc := ContactUseCase{UcContract: uc.UcContract}
	contact, err := contactUc.BrowseAllZakatDisbursement()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "select_contact", uc.ReqID)
		return
	}

	mqueue := amqp.NewQueue(AmqpConnection, AmqpChannel)
	for _, c := range contact {
		queueBody := map[string]interface{}{
			"qid":     uc.UcContract.ReqID,
			"contact": c,
		}
		AmqpConnection, AmqpChannel, err = mqueue.PushQueueReconnect(os.Getenv("AMQP_URL"), queueBody, amqp.DisbursementMutation, amqp.DisbursementMutationDeadLetter)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "amqp", uc.ReqID)
			failed++
		} else {
			success++
		}
	}

	logruslogger.Log(logruslogger.InfoLevel, "Success: "+strconv.Itoa(success)+", Failed: "+strconv.Itoa(failed), ctx, "finish", uc.ReqID)
}
