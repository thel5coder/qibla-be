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

func (uc TransactionUseCase) EditTrxID(ID, trxID string) (err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	err = repository.EditTrxID(ID, trxID, now, uc.TX)

	return err
}

func (uc TransactionUseCase) Add(input requests.TransactionRequest) (res viewmodel.TransactionVm, err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	//get invoice Number
	invoiceNumber, err := uc.GetInvoiceNumber()
	if err != nil {
		return res, err
	}

	//add table transaction
	var transactionDetails []viewmodel.TransactionDetailVm
	for _, transactionDetail := range input.TransactionDetail {
		transactionDetails = append(transactionDetails, viewmodel.TransactionDetailVm{
			Name:     transactionDetail.Name,
			Fee:      transactionDetail.Fee,
			Price:    transactionDetail.Price,
			Quantity: transactionDetail.Quantity,
			SubTotal: float32(transactionDetail.Quantity) * transactionDetail.Price,
		})
	}
	body := viewmodel.TransactionVm{
		UserID:            input.UserID,
		InvoiceNumber:     invoiceNumber,
		DueDate:           input.DueDate,
		DueDatePeriod:     input.DueDateAging,
		PaymentStatus:     enums.KeyPaymentStatus1,
		PaymentMethodCode: input.PaymentMethodeCode,
		BankName:          input.BankName,
		Direction:         enums.KeyTransactionDirection1,
		TransactionType:   enums.KeyTransactionType3,
		TransactionDate:   now,
		UpdatedAt:         now,
		Details:           transactionDetails,
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

	//post data to faspay
	faspayUc := FaspayUseCase{UcContract: uc.UcContract}
	input.FaspayBody.InvoiceNumber = invoiceNumber
	faspayRes, err := faspayUc.PostData(input.FaspayBody)
	if err != nil {
		return res, errors.New(messages.PaymentFailed)
	}

	transactionHistoryUc := TransactionHistoryUseCase{UcContract: uc.UcContract}
	err = transactionHistoryUc.Add(faspayRes["trx_id"].(string), enums.KeyPaymentStatus1, faspayRes)
	if err != nil {
		return res, err
	}

	//update trxid or va number
	err = uc.EditTrxID(body.ID,faspayRes["trx_id"].(string))
	if err != nil {
		return res,err
	}

	return res, err
}

func (uc TransactionUseCase) AddTransactionRegisterPartner(userID, bankName string, paymentMethodCode, dueDateAging int32, extraProducts []requests.ExtraProductRequest, contact viewmodel.ContactVm) (res viewmodel.TransactionVm, err error) {
	now := time.Now().UTC()
	dueDate := now.AddDate(0, 0, int(dueDateAging)).Format("2006-01-02")
	dueDateFaspay := now.AddDate(0, 0, int(dueDateAging)).Format("2006-01-02 15:04:05")
	var total float32
	var faspayItem []requests.FaspayItemRequest
	var details []requests.TransactionDetailRequest

	for _, extraProduct := range extraProducts {
		subTotal := float32(extraProduct.Price * 1)
		total = total + subTotal
		details = append(details, requests.TransactionDetailRequest{
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

	//add to table transaction
	faspayRequest := requests.FaspayPostRequest{
		RequestTransaction:  enums.KeyTransactionType3,
		InvoiceNumber:       "",
		TransactionDate:     now.Format("2006-01-02 15:04:05"),
		DueDate:             dueDateFaspay,
		TransactionDesc:     enums.KeyTransactionType5,
		UserID:              userID,
		CustomerName:        contact.TravelAgentName,
		CustomerEmail:       contact.Email,
		CustomerPhoneNumber: contact.PhoneNumber,
		Total:               0,
		PaymentChannel:      paymentMethodCode,
		Item:                faspayItem,
	}
	transactionInput := requests.TransactionRequest{
		UserID:             userID,
		DueDate:            dueDate,
		DueDateAging:       dueDateAging,
		BankName:           bankName,
		PaymentMethodeCode: paymentMethodCode,
		TransactionDetail:  details,
		FaspayBody:         faspayRequest,
	}
	res, err = uc.Add(transactionInput)
	if err != nil {
		return res, err
	}

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
		PaymentStatus:     model.PaymentStatus,
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
