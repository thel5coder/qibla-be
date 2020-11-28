package handlers

import (
	"qibla-backend/usecase"
	"strconv"

	"github.com/labstack/echo"
)

// InvoiceHandler ...
type InvoiceHandler struct {
	Handler
}

// BrowseInvoicesHandler ...
func (handler InvoiceHandler) BrowseInvoicesHandler(ctx echo.Context) error {

	name := ctx.QueryParam("name")
	transactionType := ctx.QueryParam("transaction_type")
	invoiceNumber := ctx.QueryParam("invoice_number")
	wumberOfWorshipers := ctx.QueryParam("wumber_of_worshipers")
	transactionDate := ctx.QueryParam("transaction_date")
	feeQibla := ctx.QueryParam("fee_qibla")
	dueDate := ctx.QueryParam("due_date")
	status := ctx.QueryParam("status")
	dueDatePeriod := ctx.QueryParam("due_date_period")
	total := ctx.QueryParam("total")
	paymentStatus := ctx.QueryParam("payment_status")

	filters := make(map[string]interface{})

	if name != "" {
		filters["name"] = name
	}
	if transactionType != "" {
		filters["transaction_type"] = transactionType
	}
	if invoiceNumber != "" {
		filters["invoice_number"] = invoiceNumber
	}
	if wumberOfWorshipers != "" {
		filters["wumber_of_worshipers"] = wumberOfWorshipers
	}
	if transactionDate != "" {
		filters["transaction_date"] = transactionDate
	}
	if feeQibla != "" {
		filters["fee_qibla"] = feeQibla
	}
	if dueDate != "" {
		filters["due_date"] = dueDate
	}
	if status != "" {
		filters["status"] = status
	}
	if dueDatePeriod != "" {
		filters["due_date_period"] = dueDatePeriod
	}
	if total != "" {
		filters["total"] = total
	}
	if paymentStatus != "" {
		filters["payment_status"] = paymentStatus
	}

	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.BrowseInvoices(filters, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}
