package usecase

import (
	"errors"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/enums"
	"qibla-backend/pkg/interfacepkg"
	"qibla-backend/pkg/logruslogger"
	timepkg "qibla-backend/pkg/time"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

// DisbursementUseCase ...
type DisbursementUseCase struct {
	*UcContract
}

// Browse ...
func (uc DisbursementUseCase) Browse(filters map[string]interface{}, order, sort string, page, limit int) (res []viewmodel.DisbursementVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewDisbursementRepository(uc.DB)

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)
	disbursements, count, err := repository.Browse(filters, order, sort, limit, offset)
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
func (uc DisbursementUseCase) BrowseAll(status string) (res []viewmodel.DisbursementVm, err error) {
	repository := actions.NewDisbursementRepository(uc.DB)

	disbursements, err := repository.BrowseAll(status)
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
	res, err = uc.ReadBy("def.id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

// ReadByPaymentID ...
func (uc DisbursementUseCase) ReadByPaymentID(paymentID int) (res viewmodel.DisbursementVm, err error) {
	repository := actions.NewDisbursementRepository(uc.DB)
	disbursement, err := repository.ReadByPaymentID(paymentID)
	if err != nil {
		return res, err
	}

	res = uc.buildBody(&disbursement)

	return res, err
}

// Add ...
func (uc DisbursementUseCase) Add(input *requests.DisbursementRequest) (res viewmodel.DisbursementVm, err error) {
	ctx := "DisbursementUseCase.Add"

	now := time.Now().UTC()
	res = viewmodel.DisbursementVm{
		ContactID:             input.ContactID,
		TransactionID:         input.TransactionID,
		Total:                 input.Total,
		Status:                enums.KeyPaymentStatus1,
		DisbursementType:      input.DisbursementType,
		StartPeriod:           input.StartPeriod,
		EndPeriod:             input.EndPeriod,
		DisburseAt:            input.DisburseAt,
		AccountNumber:         input.AccountNumber,
		AccountName:           input.AccountName,
		AccountBankName:       input.AccountBankName,
		AccountBankCode:       input.AccountBankCode,
		OriginAccountNumber:   input.OriginAccountNumber,
		OriginAccountName:     input.OriginAccountName,
		OriginAccountBankName: input.OriginAccountBankName,
		OriginAccountBankCode: input.OriginAccountBankCode,
		CreatedAt:             now.Format(time.RFC3339),
		UpdatedAt:             now.Format(time.RFC3339),
	}
	repository := actions.NewDisbursementRepository(uc.DB)
	res.ID, err = repository.Add(res, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	disbursementDetailUc := DisbursementDetailUseCase{UcContract: uc.UcContract}
	err = disbursementDetailUc.AddBulk(res.ID, &input.Details)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "add_detail", uc.ReqID)
		return res, err
	}

	return res, err
}

func (uc DisbursementUseCase) checkZakatInput(input *requests.DisbursementRequest, transaction *[]viewmodel.TransactionVm) (err error) {
	ctx := "DisbursementUseCase.checkZakatInput"

	for _, d := range *transaction {
		input.Total += float64(d.Total) - float64(d.FeeQibla)

		date, err := timepkg.Parse(d.TransactionDate, time.RFC3339, DefaultLocation)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "parse_transaction_date", uc.ReqID)
			return err
		}

		if input.StartPeriod == "" {
			input.StartPeriod = d.TransactionDate
		} else {
			startDate, err := timepkg.Parse(input.StartPeriod, time.RFC3339, DefaultLocation)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "parse_start_date", uc.ReqID)
				return err
			}
			if startDate.After(date) {
				input.StartPeriod = d.TransactionDate
			}
		}

		if input.EndPeriod == "" {
			input.EndPeriod = d.TransactionDate
		} else {
			endDate, err := timepkg.Parse(input.EndPeriod, time.RFC3339, DefaultLocation)
			if err != nil {
				logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "parse_end_date", uc.ReqID)
				return err
			}
			if endDate.Before(date) {
				input.EndPeriod = d.TransactionDate
			}
		}
	}

	return err
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
	if len(transaction) == 0 {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "empty_transaction", uc.ReqID)
		return errors.New("Empty")
	}

	var bodyDetails []requests.DisbursementDetailRequest
	for _, t := range transaction {
		bodyDetails = append(bodyDetails, requests.DisbursementDetailRequest{
			TransactionID: t.ID,
		})
	}

	body := requests.DisbursementRequest{
		ContactID:             contact.ID,
		DisbursementType:      enums.KeyTransactionType1,
		AccountNumber:         contact.AccountNumber,
		AccountName:           contact.AccountName,
		AccountBankName:       contact.AccountBankName,
		AccountBankCode:       contact.AccountBankCode,
		OriginAccountNumber:   DefaultOriginAccountNumber,
		OriginAccountName:     DefaultOriginAccountName,
		OriginAccountBankName: DefaultOriginAccountBankName,
		OriginAccountBankCode: DefaultOriginAccountBankCode,
		Details:               bodyDetails,
	}
	err = uc.checkZakatInput(&body, &transaction)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "check_input", uc.ReqID)
		return err
	}

	// Add record to zakat
	_, err = uc.Add(&body)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "add", uc.ReqID)
		return err
	}

	// Change transaction is disburse
	for _, t := range transaction {
		err = transactionUc.EditIsDisburse(t.ID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "edit_transaction", uc.ReqID)
			return err
		}
	}

	return err
}

// EditPaymentDetails ...
func (uc DisbursementUseCase) EditPaymentDetails(ID string, paymentDetails interface{}, status string) (err error) {
	now := time.Now().UTC().Format(time.RFC3339)
	repository := actions.NewDisbursementRepository(uc.DB)
	body := viewmodel.DisbursementVm{
		ID:             ID,
		PaymentDetails: paymentDetails,
		Status:         status,
		UpdatedAt:      now,
	}
	err = repository.EditPaymentDetails(body, uc.TX)
	if err != nil {
		return err
	}

	return err
}

// EditStatus ...
func (uc DisbursementUseCase) EditStatus(ID, status string) (err error) {
	now := time.Now().UTC().Format(time.RFC3339)
	repository := actions.NewDisbursementRepository(uc.DB)
	body := viewmodel.DisbursementVm{
		ID:        ID,
		Status:    status,
		UpdatedAt: now,
	}
	err = repository.EditStatus(body, uc.TX)
	if err != nil {
		return err
	}

	return err
}

// DisbursementReq ...
func (uc DisbursementUseCase) DisbursementReq(data *requests.DisbursementReqRequest) (err error) {
	for _, d := range data.Data {
		disbursement, err := uc.ReadByPk(d.ID)
		if err != nil {
			return err
		}
		if disbursement.Status != enums.KeyPaymentStatus1 {
			return errors.New("Invalid status")
		}

		err = uc.EditStatus(d.ID, enums.KeyPaymentStatus4)
		if err != nil {
			return err
		}
	}

	return err
}

// DisbursementFlip ...
func (uc DisbursementUseCase) DisbursementFlip(id string) (err error) {
	disbursement, err := uc.ReadByPk(id)
	if err != nil {
		return err
	}
	if disbursement.Status != enums.KeyPaymentStatus4 {
		return errors.New("Invalid status")
	}

	flipUc := FlipUseCase{UcContract: uc.UcContract}
	res, err := flipUc.Disbursement(
		id, disbursement.AccountNumber, disbursement.AccountBankCode, disbursement.Total,
		disbursement.DisbursementType, "",
	)
	if err != nil {
		return err
	}

	err = uc.EditPaymentDetails(id, res, enums.KeyPaymentStatus5)
	if err != nil {
		return err
	}

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
	interfacepkg.UnmarshallCb(data.PaymentDetails.String, &res.PaymentDetails)

	return viewmodel.DisbursementVm{
		ID:                           data.ID,
		ContactID:                    data.ContactID,
		ContactBranchName:            data.Contact.BranchName.String,
		ContactTravelAgentName:       data.Contact.TravelAgentName.String,
		ContactAddress: data.Contact.Address.String,
		ContactPhoneNumber: data.Contact.PhoneNumber,
		TransactionID:                data.TransactionID.String,
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
		OriginAccountNumber:          data.OriginAccountNumber.String,
		OriginAccountName:            data.OriginAccountName.String,
		OriginAccountBankName:        data.OriginAccountBankName.String,
		OriginAccountBankCode:        data.OriginAccountBankCode.String,
		PaymentDetails:               res.PaymentDetails,
		CreatedAt:                    data.CreatedAt,
		UpdatedAt:                    data.UpdatedAt,
		DeletedAt:                    data.DeletedAt.String,
	}
}
