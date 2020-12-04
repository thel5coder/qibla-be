package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
)

// FlipHandler ...
type FlipHandler struct {
	Handler
}

// GetBank ...
func (handler FlipHandler) GetBank(ctx echo.Context) error {
	uc := usecase.FlipUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.GetBank()

	return handler.SendResponse(ctx, res, nil, err)
}

// GetBankByCode ...
func (handler FlipHandler) GetBankByCode(ctx echo.Context) error {
	code := ctx.Param("code")

	uc := usecase.FlipUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.GetBankByCode(code)

	return handler.SendResponse(ctx, res, nil, err)
}

// DisbursementCallback ...
func (handler FlipHandler) DisbursementCallback(ctx echo.Context) error {
	input := new(requests.FlipDisbursementCallbackRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	if input.Token != handler.UseCaseContract.Flip.ValidationToken {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, "Invalid Token")
	}

	uc := usecase.FlipUseCase{UcContract: handler.UseCaseContract}
	err := uc.DisbursementCallbackQueue(&input.Data)
	if err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	return handler.SendResponse(ctx, nil, nil, err)
}
