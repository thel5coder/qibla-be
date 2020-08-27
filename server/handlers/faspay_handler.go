package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
)

type FasPayHandler struct {
	Handler
}

func (handler FasPayHandler) GetLisPaymentMethods(ctx echo.Context) error {
	uc := usecase.FaspayUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.GetLisPaymentMethods()

	return handler.SendResponse(ctx, res, nil, err)
}
