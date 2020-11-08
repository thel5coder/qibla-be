package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

// DisbursementDetailUseCase ...
type DisbursementDetailUseCase struct {
	*UcContract
}

// BrowseAll ...
func (uc DisbursementDetailUseCase) BrowseAll(disbursementID string) (res []viewmodel.DisbursementDetailVm, err error) {
	repository := actions.NewDisbursementDetailRepository(uc.DB)

	disbursements, err := repository.BrowseAll(disbursementID)
	if err != nil {
		return res, err
	}

	for _, disbursement := range disbursements {
		res = append(res, uc.buildBody(&disbursement))
	}

	return res, err
}

// Add ...
func (uc DisbursementDetailUseCase) Add(input *requests.DisbursementDetailRequest) (res viewmodel.DisbursementDetailVm, err error) {
	res = viewmodel.DisbursementDetailVm{
		DisbursementID: input.DisbursementID,
		TransactionID:  input.TransactionID,
	}
	repository := actions.NewDisbursementDetailRepository(uc.DB)
	res.ID, err = repository.Add(res, uc.TX)
	if err != nil {
		return res, err
	}

	return res, err
}

// AddBulk ...
func (uc DisbursementDetailUseCase) AddBulk(disbursementID string, input *[]requests.DisbursementDetailRequest) (err error) {
	for _, i := range *input {
		i.DisbursementID = disbursementID
		_, err = uc.Add(&i)
		if err != nil {
			return err
		}
	}

	return err
}

// Delete ...
func (uc DisbursementDetailUseCase) Delete(disbursementID string) (err error) {
	now := time.Now().UTC().Format(time.RFC3339)
	repository := actions.NewDisbursementDetailRepository(uc.DB)
	err = repository.Delete(disbursementID, now, now, uc.TX)
	if err != nil {
		return err
	}

	return err
}

func (uc DisbursementDetailUseCase) buildBody(data *models.DisbursementDetail) (res viewmodel.DisbursementDetailVm) {
	return viewmodel.DisbursementDetailVm{
		DisbursementID:               data.DisbursementID,
		TransactionID:                data.TransactionID,
		TransactionInvoiceNumber:     data.Transaction.InvoiceNumber.String,
		TransactionPaymentMethodCode: data.Transaction.PaymentMethodCode.Int32,
		TransactionPaymentStatus:     data.Transaction.PaymentStatus.String,
		TransactionDueDate:           data.Transaction.DueDate.String,
		TransactionVaNumber:          data.Transaction.VaNumber.String,
		TransactionBankName:          data.Transaction.BankName.String,
	}
}
