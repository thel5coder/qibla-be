package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
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

func (handler FasPayHandler) CheckPaymentNotification(ctx echo.Context) error {
	invoiceID := ctx.Param("invoiceId")
	uc := usecase.FaspayUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.CheckPaymentStatus(invoiceID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler FasPayHandler) PaymentNotification(ctx echo.Context) error {
	input := new(requests.PaymentNotificationRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.FaspayUseCase{UcContract: handler.UseCaseContract}
	res, _, code := uc.PaymentNotification(input)

	return ctx.JSON(code, res)
}
