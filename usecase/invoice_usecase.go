package usecase

import (
	"errors"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/messages"
	"qibla-backend/helpers/str"
	"qibla-backend/usecase/viewmodel"
	"time"
)

// InvoiceUseCase ...
type InvoiceUseCase struct {
	*UcContract
}

var (
	allowOrderInvoice = []string{
		actions.InvoiceFieldDate,
	}

	allowSortInvoice = []string{
		actions.DefaultSortAsc, actions.DefaultSortDesc,
	}
)

// Browse ...
func (uc InvoiceUseCase) Browse(order, sort string, page, limit int) (res []viewmodel.InvoiceVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewInvoiceRepository(uc.DB)

	checkOrder := str.StringInSlice(order, allowOrderInvoice)
	if checkOrder == false {
		return res, pagination, errors.New(messages.OrderNotMatch)
	}

	checkShort := str.StringInSlice(sort, allowSortInvoice)
	if checkShort == false {
		return res, pagination, errors.New(messages.SortNotMatch)
	}

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	invoices, count, err := repository.Browse(order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, invoice := range invoices {
		res = append(res, uc.buildBody(invoice))
	}

	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

// buildBody ...
func (uc InvoiceUseCase) buildBody(data models.Invoice) (res viewmodel.InvoiceVm) {

	var status string

	if data.PaymentStatus.String == "pending" {
		status = "open"
	} else if data.PaymentStatus.String == "finish" {
		status = "close"

	}

	if data.DueDate.String != "" {
		dueDate, _ := time.Parse(time.RFC3339, data.DueDate.String)
		transactionDate, _ := time.Parse(time.RFC3339, data.TransactionDate)

		if transactionDate.After(dueDate) {
			status = "overdue"
		}
	}

	return viewmodel.InvoiceVm{
		ID:              data.ID,
		Name:            data.Name.String,
		TransactionType: data.TransactionType,
		InvoiceNumber:   data.InvoiceNumber,
		FeeQibla:        data.FeeQibla.Float64,
		Total:           data.Total.Float64,
		DueDate:         data.DueDate.String,
		BillingStatus:   status,
		DueDatePeriod:   data.DueDatePeriod.Int32,
		PaymentStatus:   data.PaymentStatus.String,
		PaidDate:        data.PaidDate.String,
		InvoiceStatus:   data.InvoiceStatus.String,
		Direction:       data.Direction,
		TransactionDate: data.TransactionDate,
		UpdatedAt:       data.UpdatedAt,
	}
}
