package usecase

import (
	"errors"
	"qibla-backend/helpers/flip"
	functionCaller "qibla-backend/helpers/functioncaller"
	"qibla-backend/helpers/logruslogger"
	"qibla-backend/helpers/messages"
	"qibla-backend/helpers/interfacepkg"
)

// FlipUseCase ...
type FlipUseCase struct {
	*UcContract
}

// GetBank ...
func (uc FlipUseCase) GetBank() (res []flip.Bank, err error) {
	res, err = uc.Flip.GetBank()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "get_bank")
		return res, err
	}

	return res, err
}

// GetBankByCode ...
func (uc FlipUseCase) GetBankByCode(code string) (res flip.Bank, err error) {
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
func (uc FlipUseCase) Disbursement(id, accountNumber, bankCode string, amount float64, remark, recipientCity string) (res map[string]interface{}, err error) {
	res, err = uc.Flip.Disbursement(id, accountNumber, bankCode, amount, remark, recipientCity)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "disbursement")
		return res, err
	}
	if res["code"] != nil {
		logruslogger.Log(logruslogger.WarnLevel, interfacepkg.Marshall(res), functionCaller.PrintFuncName(), "flip_error")
		return res, errors.New(messages.InternalServer)
	}

	return res, err
}
