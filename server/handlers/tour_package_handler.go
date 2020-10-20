package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
)

type TourPackageHandler struct {
	Handler
}

func (handler TourPackageHandler) Browse(ctx echo.Context) error {
	uc := usecase.TourPackageUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll()

	return handler.SendResponse(ctx, res, nil, err)
}
