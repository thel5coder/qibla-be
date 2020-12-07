package usecase

import (
	"errors"
	"qibla-backend/pkg/enums"
	functionCaller "qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/interfacepkg"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/messages"
	"qibla-backend/pkg/amqp"
	"os"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
)

// FlipUseCase ...
type FlipUseCase struct {
	*UcContract
}

// GetBank ...
func (uc FlipUseCase) GetBank() (res []viewmodel.BankVM, err error) {
	res, err = uc.Flip.GetBank()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "get_bank")
		return res, err
	}

	return res, err
}

// GetBankByCode ...
func (uc FlipUseCase) GetBankByCode(code string) (res viewmodel.BankVM, err error) {
	data, err := uc.GetBank()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "get_bank")
		return res, err
	}

	for _, d := range data {
		if d.BankCode == code {
			res = d
		}
	}

	if res.BankCode == "" {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "not_found")
		return res, errors.New(messages.DataNotFound)
	}

	return res, err
}

// Disbursement ...
func (uc FlipUseCase) Disbursement(id, accountNumber, bankCode string, amount float64, remark, recipientCity string) (res viewmodel.DisbursementVM, err error) {
	res, err = uc.Flip.Disbursement(id, accountNumber, bankCode, amount, remark, recipientCity)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "disbursement")
		return res, err
	}
	if res.Code != "" {
		logruslogger.Log(logruslogger.WarnLevel, interfacepkg.Marshall(res), functionCaller.PrintFuncName(), "flip_error")
		return res, errors.New(messages.InternalServer)
	}

	return res, err
}

// DisbursementCallbackQueue ...
func (uc FlipUseCase) DisbursementCallbackQueue(data *requests.FlipDisbursementRequest) (err error) {
	mqueue := amqp.NewQueue(AmqpConnection, AmqpChannel)
	queueBody := map[string]interface{}{
		"qid":  uc.UcContract.ReqID,
		"data": data,
	}
	AmqpConnection, AmqpChannel, err = mqueue.PushQueueReconnect(os.Getenv("AMQP_URL"), queueBody, amqp.DisbursementCallback, amqp.DisbursementCallbackDeadLetter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "amqp")
		return err
	}

	return err
}

// DisbursementCallback ...
func (uc FlipUseCase) DisbursementCallback(data *requests.FlipDisbursementRequest) (err error) {
	disbursementUc := DisbursementUseCase{UcContract: uc.UcContract}
	disbursement, err := disbursementUc.ReadByPaymentID(data.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "find_disbursement")
		return err
	}
	if disbursement.Status != enums.KeyPaymentStatus5 {
		logruslogger.Log(logruslogger.WarnLevel, "", functionCaller.PrintFuncName(), "disbursement_status")
		return errors.New(messages.InvalidStatus)
	}

	status := enums.KeyPaymentStatus2
	if data.Status == "DONE" {
		status = enums.KeyPaymentStatus3
	}
	err = disbursementUc.EditPaymentDetails(disbursement.ID, data, status)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "edit_payment")
		return err
	}

	return err
}
