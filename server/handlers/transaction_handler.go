package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
)

type TransactionHandler struct {
	Handler
}

func (handler TransactionHandler) GetInvoiceCount(ctx echo.Context) error {
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.GetInvoiceNumber()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler TransactionHandler) ReadByTrxID(ctx echo.Context) error {
	trxID := ctx.Param("trxId")
	uc := usecase.TransactionUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("t.trx_id", trxID, "=")

	return handler.SendResponse(ctx, res, nil, err)
}
