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

func (handler OdooMasterPackageHandler) ReadByPk(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.OdooMasterPackageUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("m.id", ID, "=")

	return handler.SendResponse(ctx, res, nil, err)
}
