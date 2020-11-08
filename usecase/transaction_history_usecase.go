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
		ID:            transactionHistory.ID,
		TransactionID: transactionHistory.TransactionID,
		Status:        transactionHistory.Status,
		Response:      transactionHistory.Response,
		CreatedAt:     transactionHistory.CreatedAt,
		UpdatedAt:     transactionHistory.UpdatedAt,
	}

	return res, err
}

func (uc TransactionHistoryUseCase) Edit(transactionID, status string, response map[string]interface{}) (err error) {
	repository := actions.NewTransactionHistoryRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	responseByte, err := json.Marshal(response)

	_, err = repository.EditByTrxID(transactionID, status, string(responseByte), now)
	if err != nil {
		return err
	}

	return nil
}

func (uc TransactionHistoryUseCase) Add(transactionID, status string, response map[string]interface{}) (err error) {
	repository := actions.NewTransactionHistoryRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	responseByte, err := json.Marshal(response)
	err = repository.Add(transactionID, status, string(responseByte), now, now, uc.TX)
	if err != nil {
		return err
	}

	return nil
}
