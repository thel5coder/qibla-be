package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
)

type OdooMasterPackageHandler struct {
	Handler
}

func (handler OdooMasterPackageHandler) BrowseAll(ctx echo.Context) error {
	partnerID := ctx.QueryParam("partner_id")
	uc := usecase.OdooMasterPackageUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll(partnerID)

	return handler.SendResponse(ctx, res, nil, err)
}
