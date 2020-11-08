package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/enums"
	"qibla-backend/helpers/interfacepkg"
	"qibla-backend/helpers/logruslogger"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

// DisbursementUseCase ...
type DisbursementUseCase struct {
	*UcContract
}

// Browse ...
func (uc DisbursementUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.DisbursementVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewDisbursementRepository(uc.DB)

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)
	disbursements, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, disbursement := range disbursements {
		res = append(res, uc.buildBody(&disbursement))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

// BrowseAll ...
func (uc DisbursementUseCase) BrowseAll() (res []viewmodel.DisbursementVm, err error) {
	repository := actions.NewDisbursementRepository(uc.DB)

	disbursements, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, disbursement := range disbursements {
		res = append(res, uc.buildBody(&disbursement))
	}

	return res, err
}

// BrowseBy ...
func (uc DisbursementUseCase) BrowseBy(column, value, operator string) (res []viewmodel.DisbursementVm, err error) {
	repository := actions.NewDisbursementRepository(uc.DB)
	disbursements, err := repository.BrowseBy(column, value, operator)

	for _, disbursement := range disbursements {
		res = append(res, uc.buildBody(&disbursement))
	}

	return res, err
}

// ReadBy ...
func (uc DisbursementUseCase) ReadBy(column, value string) (res viewmodel.DisbursementVm, err error) {
	repository := actions.NewDisbursementRepository(uc.DB)
	disbursement, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = uc.buildBody(&disbursement)

	return res, err
}

// ReadByPk ...
func (uc DisbursementUseCase) ReadByPk(ID string) (res viewmodel.DisbursementVm, err error) {
	res, err = uc.ReadBy("uz.id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc DisbursementUseCase) checkInput(input *requests.DisbursementRequest) (err error) {
	var details []string
	for _, d := range input.Details {
		details = append(details, d.TransactionID)
	}

	return err
}

// Add ...
func (uc DisbursementUseCase) Add(input *requests.DisbursementRequest) (res viewmodel.DisbursementVm, err error) {
	err = uc.checkInput(input)
	if err != nil {
		return res, err
	}

	// transactionUseCase := TransactionUseCase{UcContract: uc.UcContract}
	// transaction, err := transactionUseCase.AddTransactionZakat(input)
	// if err != nil {
	// 	return res, err
	// }

	now := time.Now().UTC()
	res = viewmodel.DisbursementVm{
		ContactID: input.ContactID,
		// TransactionID:    transaction.ID,
		Total:            input.Total,
		Status:           enums.KeyPaymentStatus1,
		DisbursementType: input.DisbursementType,
		StartPeriod:      input.StartPeriod,
		EndPeriod:        input.EndPeriod,
		DisburseAt:       input.DisburseAt,
		AccountNumber:    input.AccountNumber,
		AccountName:      input.AccountName,
		AccountBankName:  input.AccountBankName,
		AccountBankCode:  input.AccountBankCode,
		CreatedAt:        now.Format(time.RFC3339),
		UpdatedAt:        now.Format(time.RFC3339),
	}
	repository := actions.NewDisbursementRepository(uc.DB)
	res.ID, err = repository.Add(res, uc.TX)
	if err != nil {
		return res, err
	}

	disbursementDetailUc := DisbursementDetailUseCase{UcContract: uc.UcContract}
	err = disbursementDetailUc.AddBulk(res.ID, &input.Details)
	if err != nil {
		return res, err
	}

	return res, err
}

// AddZakatByContact ...
func (uc DisbursementUseCase) AddZakatByContact(contact *viewmodel.ContactVm) (err error) {
	ctx := "DisbursementUseCase.AddZakatByContact"

	// Get all transaction
	transactionUc := TransactionUseCase{UcContract: uc.UcContract}
	transaction, err := transactionUc.BrowseAllZakatDisbursement(contact.ID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "transaction", uc.ReqID)
		return err
	}

	logruslogger.Log(logruslogger.InfoLevel, interfacepkg.Marshall(transaction), ctx, "transaction", uc.ReqID)

	return err
}

// Delete ...
func (uc DisbursementUseCase) Delete(ID string) (err error) {
	now := time.Now().UTC().Format(time.RFC3339)
	repository := actions.NewDisbursementRepository(uc.DB)
	err = repository.Delete(ID, now, now, uc.TX)
	if err != nil {
		return err
	}

	return err
}

func (uc DisbursementUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewDisbursementRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc DisbursementUseCase) buildBody(data *models.Disbursement) (res viewmodel.DisbursementVm) {
	return viewmodel.DisbursementVm{
		ID:                           data.ID,
		TransactionID:                data.TransactionID,
		TransactionInvoiceNumber:     data.Transaction.InvoiceNumber.String,
		TransactionPaymentMethodCode: data.Transaction.PaymentMethodCode.Int32,
		TransactionPaymentStatus:     data.Transaction.PaymentStatus.String,
		TransactionDueDate:           data.Transaction.DueDate.String,
		TransactionVaNumber:          data.Transaction.VaNumber.String,
		TransactionBankName:          data.Transaction.BankName.String,
		Total:                        data.Total.Float64,
		Status:                       data.Status.String,
		DisbursementType:             data.DisbursementType.String,
		StartPeriod:                  data.StartPeriod.String,
		EndPeriod:                    data.EndPeriod.String,
		DisburseAt:                   data.DisburseAt.String,
		AccountNumber:                data.AccountNumber.String,
		AccountName:                  data.AccountName.String,
		AccountBankName:              data.AccountBankName.String,
		AccountBankCode:              data.AccountBankCode.String,
		CreatedAt:                    data.CreatedAt,
		UpdatedAt:                    data.UpdatedAt,
		DeletedAt:                    data.DeletedAt.String,
	}
}
