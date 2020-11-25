package usecase

import (
	functionCaller "qibla-backend/helpers/functioncaller"
	"qibla-backend/helpers/logruslogger"
)

// FlipUseCase ...
type FlipUseCase struct {
	*UcContract
}

// GetBank ...
func (uc FlipUseCase) GetBank() (res map[string]interface{}, err error) {
	res, err = uc.Flip.GetBank()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "get_bank")
		return res, err
	}

	return res, err
}

// Disbursement ...
func (uc FlipUseCase) Disbursement(id, accountNumber, bankCode string, amount float64, remark, recipientCity string) (res map[string]interface{}, err error) {
	res, err = uc.Flip.Disbursement(id, accountNumber, bankCode, amount, remark, recipientCity)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "disbursement")
		return res, err
	}

	return res, err
}
