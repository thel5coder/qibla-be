package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
)

type EnumHandler struct {
	Handler
}

func (handler EnumHandler) GetMenuPermissions(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetMenuPermissionsEnums()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetWebComprofCategori(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetWebComprofCategoryEnums()

	return handler.SendResponse(ctx, res, nil, nil)
}
