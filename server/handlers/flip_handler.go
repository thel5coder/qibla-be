package handlers

import (
	"github.com/labstack/echo"
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
