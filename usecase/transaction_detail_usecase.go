package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type TransactionDetailUseCase struct {
	*UcContract
}

func (uc TransactionDetailUseCase) BrowseByTransactionID(transactionID string) (res []viewmodel.TransactionDetailVm, err error) {
	repository := actions.NewTransactionDetailRepository(uc.DB)
	transactionDetails, err := repository.BrowseByTransactionID(transactionID)
	if err != nil {
		return res, err
	}

	for _, transactionDetail := range transactionDetails {
		var subTotal float32
		if transactionDetail.Fee > 0 {
			subTotal = (transactionDetail.Price * float32(transactionDetail.Quantity)) - transactionDetail.Fee
		} else {
			subTotal = transactionDetail.Price * float32(transactionDetail.Quantity)
		}
		res = append(res, viewmodel.TransactionDetailVm{
			ID:       transactionDetail.ID,
			Name:     transactionDetail.Name,
			Fee:      transactionDetail.Fee,
			Price:    transactionDetail.Price,
			Quantity: transactionDetail.Quantity,
			SubTotal: subTotal,
		})
	}

	return res, err
}

func (uc TransactionDetailUseCase) Add(transactionID, name string, fee, price float32, quantity int) (err error) {
	repository := actions.NewTransactionDetailRepository(uc.DB)
	err = repository.Add(transactionID, name, fee, price, quantity, uc.TX)

	return err
}

func (uc TransactionDetailUseCase) DeleteByTransactionID(transactionID string) (err error) {
	repository := actions.NewTransactionDetailRepository(uc.DB)
	err = repository.DeleteByTransactionID(transactionID, uc.TX)

	return err
}

func (uc TransactionDetailUseCase) Store(transactionID string, inputs []viewmodel.TransactionDetailVm) (err error) {
	transactionDetails, err := uc.BrowseByTransactionID(transactionID)
	if err != nil {
		return err
	}

	if len(transactionDetails) > 0 {
		err = uc.DeleteByTransactionID(transactionID)
		if err != nil {
			return err
		}
	}

	for _, input := range inputs {
		err = uc.Add(transactionID, input.Name,input.Fee, input.Price,input.Quantity)
		if err != nil {
			return err
		}
	}

	return nil
}
