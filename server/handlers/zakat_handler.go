package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
)

type ZakatHandler struct {
	Handler
}

func (handler ZakatHandler) BrowseAll(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	uc := usecase.ContactUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAllZakatPlace(search)

	return handler.SendResponse(ctx, res, nil, err)
}
