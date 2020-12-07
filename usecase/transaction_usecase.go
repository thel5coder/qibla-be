package usecase

import (
	"errors"
	"fmt"
	"os"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/enums"
	"qibla-backend/pkg/messages"
	"qibla-backend/pkg/str"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"strconv"
	"time"
)

type TransactionUseCase struct {
	*UcContract
}

var (
	allowOrderInvoice = []string{
		actions.TransactionFieldDate, actions.TransactionFieldInvoice, actions.TransactionFieldTrxID,
		actions.TransactionFieldDueDate, actions.TransactionFieldDueDatePeriod, actions.TransactionFieldPaymentStatus,
		actions.TransactionFieldPaymentMethodCode, actions.TransactionFieldVaNumber, actions.TransactionFieldBankName,
		actions.TransactionFieldTransactionType, actions.TransactionFieldPaidDate, actions.TransactionFieldTransactionDate,
		actions.TransactionFieldTotal, actions.TransactionFieldFeeQibla, actions.TransactionFieldIsDisburse,
	}

	allowSortInvoice = []string{
		actions.DefaultSortAsc, actions.DefaultSortDesc,
	}
)

// BrowseAllZakatDisbursement ...
func (uc TransactionUseCase) BrowseAllZakatDisbursement(contactID string) (res []viewmodel.TransactionVm, err error) {
	repository := actions.NewTransactionRepository(uc.DB)

	data, err := repository.BrowseAllZakatDisbursement(contactID)
	if err != nil {
		return res, err
	}

	for _, d := range data {
		res = append(res, uc.buildBody(d))
	}

	return res, err
}

// BrowseInvoices ...
func (uc TransactionUseCase) BrowseInvoices(filters map[string]interface{}, order, sort string, page, limit int) (res []viewmodel.TransactionVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewTransactionRepository(uc.DB)

	checkOrder := str.StringInSlice(order, allowOrderInvoice)
	if checkOrder == false {
		order = actions.TransactionFieldDate
	}

	checkShort := str.StringInSlice(sort, allowSortInvoice)
	if checkShort == false {
		sort = actions.DefaultSortDesc
	}

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	invoices, count, err := repository.BrowseInvoices(filters, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, invoice := range invoices {
		res = append(res, uc.buildBody(invoice))
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
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

	err = repository.EditStatus(ID, paymentStatus, paidDate, now, uc.TX)

	return err
}

func (uc TransactionUseCase) EditTrxID(ID, trxID string) (err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	err = repository.EditTrxID(ID, trxID, now, uc.TX)

	return err
}

func (uc TransactionUseCase) EditIsDisburse(ID string) (err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	err = repository.EditIsDisburse(ID, now, uc.TX)

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
	res = viewmodel.TransactionVm{
		UserID:            input.UserID,
		InvoiceNumber:     invoiceNumber,
		DueDate:           input.DueDate,
		DueDatePeriod:     input.DueDateAging,
		PaymentStatus:     enums.KeyPaymentStatus1,
		InvoiceStatus:     enums.InvoiceStatusEnum[0],
		PaymentMethodCode: input.PaymentMethodeCode,
		BankName:          input.BankName,
		Direction:         input.TransactionDirection,
		TransactionType:   input.TransactionType,
		TransactionDate:   now,
		UpdatedAt:         now,
		Details:           transactionDetails,
		Total:             input.FaspayBody.Total,
		FeeQibla:          input.FeeQibla,
		IsDisburse:        input.IsDisburse,
		IsDisburseAllowed: input.IsDisburseAllowed,
	}
	res.ID, err = repository.Add(res, uc.TX)
	if err != nil {
		return res, err
	}

	transactionDetailUc := TransactionDetailUseCase{UcContract: uc.UcContract}
	err = transactionDetailUc.Store(res.ID, res.Details)
	if err != nil {
		return res, err
	}

	//post data to faspay_api
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
	err = uc.EditTrxID(res.ID, faspayRes["trx_id"].(string))
	if err != nil {
		return res, err
	}
	res.TrxID = faspayRes["trx_id"].(string)
	res.VaNumber = faspayRes["trx_id"].(string)
	res.FaspayResponse = faspayRes

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
		RequestTransaction:  enums.KeyTransactionType5,
		InvoiceNumber:       "",
		TransactionDate:     now.Format("2006-01-02 15:04:05"),
		DueDate:             dueDateFaspay,
		TransactionDesc:     enums.KeyTransactionType5,
		UserID:              userID,
		CustomerName:        contact.TravelAgentName,
		CustomerEmail:       contact.Email,
		CustomerPhoneNumber: contact.PhoneNumber,
		Total:               total,
		PaymentChannel:      paymentMethodCode,
		Item:                faspayItem,
	}
	transactionInput := requests.TransactionRequest{
		UserID:               userID,
		DueDate:              dueDate,
		DueDateAging:         dueDateAging,
		BankName:             bankName,
		PaymentMethodeCode:   paymentMethodCode,
		TransactionType:      enums.KeyTransactionType5,
		TransactionDirection: enums.KeyTransactionDirection1,
		TransactionDetail:    details,
		FaspayBody:           faspayRequest,
		IsDisburseAllowed:    false,
	}
	res, err = uc.Add(transactionInput)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc TransactionUseCase) Delete(ID string) (err error) {
	repository := actions.NewTransactionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.CountBy("", "id", ID)
	if err != nil {
		return err
	}

	if count > 0 {
		_, err = repository.Delete(ID, now, now)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc TransactionUseCase) AddTransactionZakat(input *requests.UserZakatRequest) (res viewmodel.TransactionVm, err error) {
	userUseCase := UserUseCase{UcContract: uc.UcContract}
	user, err := userUseCase.ReadBy("u.id", uc.UserID)
	if err != nil {
		return res, err
	}

	// Get bank name
	faspayUc := FaspayUseCase{UcContract: uc.UcContract}
	bankName, err := faspayUc.GetLisPaymentMethodsByCode(strconv.Itoa(int(input.PaymentMethodCode)))
	if err != nil {
		return res, err
	}

	now := time.Now().UTC()
	transactionInput := requests.TransactionRequest{
		UserID:               uc.UserID,
		DueDate:              now.AddDate(0, 0, int(defaultInvoiceDueDate)).Format("2006-01-02"),
		DueDateAging:         defaultInvoiceDueDate,
		BankName:             bankName,
		PaymentMethodeCode:   input.PaymentMethodCode,
		TransactionType:      enums.KeyTransactionType1,
		TransactionDirection: enums.KeyTransactionDirection1,
		TransactionDetail:    []requests.TransactionDetailRequest{},
		FeeQibla:             defaultFeeZakat,
		IsDisburseAllowed:    true,
		IsDisburse:           false,
		FaspayBody: requests.FaspayPostRequest{
			RequestTransaction:  enums.KeyTransactionType1,
			TransactionDate:     now.Format("2006-01-02 15:04:05"),
			DueDate:             now.AddDate(0, 0, int(defaultInvoiceDueDate)).Format("2006-01-02 15:04:05"),
			TransactionDesc:     enums.KeyTransactionType1,
			UserID:              uc.UserID,
			CustomerName:        user.Name,
			CustomerEmail:       user.Email,
			CustomerPhoneNumber: user.MobilePhone,
			Total:               float32(input.Total),
			PaymentChannel:      input.PaymentMethodCode,
			Item: []requests.FaspayItemRequest{
				{
					Product:     enums.KeyTransactionType1,
					Amount:      int(input.Total),
					Qty:         1,
					PaymentPlan: defaultFaspayPaymentPlan,
					Tenor:       defaultFaspayTenor,
					MerchantID:  os.Getenv("FASPAY_MERCHANT_ID"),
				},
			},
		},
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
	var status string
	var name string

	if model.PaymentStatus.String == "pending" {
		status = "open"
	} else if model.PaymentStatus.String == "finish" {
		status = "close"
	}

	if model.DueDate.String != "" {
		dueDate, _ := time.Parse(time.RFC3339, model.DueDate.String)
		transactionDate, _ := time.Parse(time.RFC3339, model.TransactionDate)

		if transactionDate.After(dueDate) {
			status = "overdue"
		}
	}

	if model.TransactionType == enums.KeyTransactionType2 {
		name = model.PartnerName.String
	} else if model.TransactionType == enums.KeyTransactionType1 {
		name = ""
	} else {
		name = model.TravelAgentName.String
	}

	res = viewmodel.TransactionVm{
		ID:                 model.ID,
		UserID:             model.UserID,
		InvoiceNumber:      model.InvoiceNumber.String,
		TrxID:              model.TrxID.String,
		DueDate:            model.DueDate.String,
		DueDatePeriod:      model.DueDatePeriod.Int32,
		PaymentStatus:      model.PaymentStatus.String,
		PaymentMethodCode:  model.PaymentMethodCode.Int32,
		VaNumber:           model.VaNumber.String,
		BankName:           model.BankName.String,
		Direction:          model.Direction,
		TransactionType:    model.TransactionType,
		PaidDate:           model.PaidDate.String,
		TransactionDate:    model.TransactionDate,
		UpdatedAt:          model.UpdatedAt,
		Total:              float32(model.Total.Float64),
		FeeQibla:           float32(model.FeeQibla.Float64),
		IsDisburse:         model.IsDisburse.Bool,
		IsDisburseAllowed:  model.IsDisburseAllowed.Bool,
		Details:            nil,
		FaspayResponse:     nil,
		InvoiceStatus:      status,
		NumberOfWorshipers: 0,
		Name:               name,
	}

	return res
}
