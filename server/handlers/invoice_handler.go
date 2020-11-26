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
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.BrowseInvoices(order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}
