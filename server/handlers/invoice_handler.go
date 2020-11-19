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

// Browse ...
func (handler InvoiceHandler) Browse(ctx echo.Context) error {
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.InvoiceUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}
