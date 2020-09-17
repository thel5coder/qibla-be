package usecase

import (
	"encoding/json"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type TransactionHistoryUseCase struct {
	*UcContract
}

func (uc TransactionHistoryUseCase) ReadBy(column, value string) (res viewmodel.TransactionHistoryVm, err error) {
	repository := actions.NewTransactionHistoryRepository(uc.DB)
	transactionHistory, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.TransactionHistoryVm{
		ID:        transactionHistory.ID,
		TrxID:     transactionHistory.TrxID,
		Status:    transactionHistory.Status,
		Response:  transactionHistory.Response,
		CreatedAt: transactionHistory.CreatedAt,
		UpdatedAt: transactionHistory.UpdatedAt,
	}

	return res, err
}

func (uc TransactionHistoryUseCase) Edit(trxID, status string, response map[string]interface{}) (err error) {
	repository := actions.NewTransactionHistoryRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	responseByte, err := json.Marshal(response)

	_, err = repository.EditByTrxID(trxID, status, string(responseByte), now)
	if err != nil {
		return err
	}

	return nil
}

func (uc TransactionHistoryUseCase) Add(trxID, status string, response map[string]interface{}) (err error) {
	repository := actions.NewTransactionHistoryRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	responseByte, err := json.Marshal(response)
	err = repository.Add(trxID, status, string(responseByte), now, now,uc.TX)
	if err != nil {
		return err
	}

	return nil
}
