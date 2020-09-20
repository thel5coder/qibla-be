package usecase

import (
	"errors"
	"fmt"
	"os"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/enums"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"strconv"
	"time"
)

type TransactionUseCase struct {
	*UcContract
}

func (uc TransactionUseCase) ReadBy(column, value, operator string) (res viewmodel.TransactionVm, err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	transaction, err := repository.ReadBy(column, value, operator)
	if err != nil {
		return res, err
	}
	res = uc.buildBody(transaction)

	return res, err
}

func (uc TransactionUseCase) EditStatus(ID, paymentStatus, paidDate string) (err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	fmt.Println("uc")
	fmt.Println(now)

	err = repository.EditStatus(ID, paymentStatus, paidDate, now, uc.TX)

	return err
}

func (uc TransactionUseCase) AddTransactionRegisterPartner(userID, invoiceNumber, bankName string, paymentMethodID, dueDateAging int, extraProducts []requests.ExtraProductRequest, contact viewmodel.ContactVm) (res viewmodel.TransactionVm, err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	now := time.Now().UTC()
	dueDate := now.AddDate(0, 0, dueDateAging).Format("2006-01-02")
	dueDateFaspay := now.AddDate(0, 0, dueDateAging).Format("2006-01-02 15:04:05")
	var total float32
	var faspayItem []requests.FaspayItemRequest

	var details []viewmodel.TransactionDetailVm
	for _, extraProduct := range extraProducts {
		subTotal := float32(extraProduct.Price * 1)
		total = total + subTotal
		details = append(details, viewmodel.TransactionDetailVm{
			ID:       "",
			Name:     extraProduct.Name,
			Fee:      0,
			Price:    float32(extraProduct.Price),
			Quantity: 1,
			SubTotal: float32(extraProduct.Price * 1),
		})
		faspayItem = append(faspayItem, requests.FaspayItemRequest{
			Product:     extraProduct.Name,
			Amount:      extraProduct.Price,
			Qty:         1,
			PaymentPlan: defaultFaspayPaymentPlan,
			Tenor:       defaultFaspayTenor,
			MerchantID:  os.Getenv("FASPAY_MERCHANT_ID"),
		})
	}

	faspayUc := FaspayUseCase{UcContract: uc.UcContract}
	faspayRequest := requests.FaspayPostRequest{
		RequestTransaction: enums.KeyTransactionType5,
		InvoiceNumber:      invoiceNumber,
		TransactionDate:    now.UTC().Format("2006-01-02 15:04:05"),
		DueDate:            dueDateFaspay,
		TransactionDesc:    enums.KeyTransactionType5,
		UserID:             userID,
		Total:              total,
		PaymentChannel:     paymentMethodID,
		Item:               faspayItem,
	}
	faspayRes, err := faspayUc.PostData(faspayRequest, contact)
	if err != nil {
		return res, errors.New(messages.PaymentFailed)
	}

	body := viewmodel.TransactionVm{
		ID:                "",
		UserID:            userID,
		InvoiceNumber:     invoiceNumber,
		TrxID:             faspayRes["trx_id"].(string),
		DueDate:           dueDate,
		DueDatePeriod:     int32(dueDateAging),
		PaymentStatus:     enums.KeyPaymentStatus1,
		PaymentMethodCode: int32(paymentMethodID),
		VaNumber:          faspayRes["trx_id"].(string),
		BankName:          bankName,
		Direction:         enums.KeyTransactionDirection1,
		TransactionType:   enums.KeyTransactionType3,
		PaidDate:          "",
		TransactionDate:   now.Format(time.RFC3339),
		UpdatedAt:         now.Format(time.RFC3339),
		Details:           details,
	}
	body.ID, err = repository.Add(body, uc.TX)
	if err != nil {
		return res, err
	}

	transactionDetailUc := TransactionDetailUseCase{UcContract: uc.UcContract}
	err = transactionDetailUc.Store(body.ID, body.Details)
	if err != nil {
		return res, err
	}

	transactionHistoryUc := TransactionHistoryUseCase{UcContract: uc.UcContract}
	err = transactionHistoryUc.Add(faspayRes["trx_id"].(string), enums.KeyPaymentStatus1, faspayRes)
	if err != nil {
		return res, err
	}
	body.FaspayResponse = faspayRes
	res = body

	return res, err
}

func (uc TransactionUseCase) GetInvoiceNumber() (res string, err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	year, month, _ := time.Now().UTC().Date()
	count, err := repository.GetInvoiceCount(int(month))
	if err != nil {
		return res, err
	}
	res = `INV-QBL/` + strconv.Itoa(int(year)) + `/` + strconv.Itoa(int(month))
	res += `/` + fmt.Sprintf("%05d", count+1)

	return res, err
}

func (uc TransactionUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc TransactionUseCase) buildBody(model models.Transaction) (res viewmodel.TransactionVm) {
	res = viewmodel.TransactionVm{
		ID:                model.ID,
		UserID:            model.UserID,
		InvoiceNumber:     model.InvoiceNumber.String,
		TrxID:             model.TrxID.String,
		DueDate:           model.DueDate,
		DueDatePeriod:     model.DueDatePeriod.Int32,
		PaymentStatus:     model.PaymentStatus.String,
		PaymentMethodCode: model.PaymentMethodCode.Int32,
		VaNumber:          model.VaNumber.String,
		BankName:          model.BankName.String,
		Direction:         model.Direction,
		TransactionType:   model.TransactionType,
		PaidDate:          model.PaidDate.String,
		TransactionDate:   model.TransactionDate,
		UpdatedAt:         model.UpdatedAt,
		Total:             model.Total,
		Details:           nil,
		FaspayResponse:    nil,
	}

	return res
}
