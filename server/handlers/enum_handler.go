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

func (handler EnumHandler) GetPromotionPackage(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetPromotionPackageEnum()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetPlatform(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetPlatformEnum()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetPositionPromotion(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetPositionPromotionEnum()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetSubscription(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetSubscriptionEnum()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetPriceUnit(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetPriceUnitEnum()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetDiscountType(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetDiscountTypeEnum()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetComplaintStatus(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetStatusComplaint()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetTypeZakat(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetTypeZakat()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetRemember(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetRememberOptions()

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler EnumHandler) GetSex(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetSexEnum()

	return handler.SendResponse(ctx, res, nil, nil)
}

// GetMaritalStatus ...
func (handler EnumHandler) GetMaritalStatus(ctx echo.Context) error {
	uc := usecase.EnumOptionsUseCase{UcContract: handler.UseCaseContract}
	res := uc.GetMaritalStatusEnum()

	return handler.SendResponse(ctx, res, nil, nil)
}
