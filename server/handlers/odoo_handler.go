package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
	"qibla-backend/usecase/viewmodel"
	"strconv"
)

type OdooHandler struct {
	Handler
}

func (handler OdooHandler) GetField(ctx echo.Context) error {
	objectName := ctx.Param("objectName")
	uc := usecase.OdooUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.GetField(objectName)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler OdooHandler) Browse(ctx echo.Context) error {
	objectName := ctx.Param("objectName")
	uc := usecase.OdooUseCase{UcContract: handler.UseCaseContract}
	var res []viewmodel.TravelPackageOdooVm

	err := uc.Browse(objectName, "", 0, 0, &res)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler OdooHandler) Read(ctx echo.Context) error{
	objectName := ctx.Param("objectName")
	ID,_ := strconv.Atoi(ctx.Param("id"))

	uc := usecase.OdooUseCase{UcContract: handler.UseCaseContract}
	var res []viewmodel.TravelPackageOdooVm

	err := uc.Read(objectName, int64(ID), &res)

	return handler.SendResponse(ctx, res, nil, err)
}
