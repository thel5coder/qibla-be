package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
)

type VideoKajianHandler struct {
	Handler
}

func (handler VideoKajianHandler) Browse(ctx echo.Context) error {
	uc := usecase.VideoKajianUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Browse()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler VideoKajianHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.VideoKajianUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Read(ID)

	return handler.SendResponse(ctx, res, nil, err)
}
