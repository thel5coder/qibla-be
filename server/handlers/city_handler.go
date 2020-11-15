package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
)

type CityHandler struct {
	Handler
}

func (handler CityHandler) BrowseAllByProvince(ctx echo.Context) error {
	provinceID := ctx.QueryParam("province_id")
	uc := usecase.CityUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAllByProvince(provinceID)

	return handler.SendResponse(ctx, res, nil, err)
}
