package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
)

type ProvinceHandler struct {
	Handler
}

func (handler ProvinceHandler) BrowseAll(ctx echo.Context) error {
	uc := usecase.ProvinceUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll()

	return handler.SendResponse(ctx, res, nil, err)
}
