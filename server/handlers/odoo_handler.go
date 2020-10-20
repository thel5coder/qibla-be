package handlers

import (
	"fmt"
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

func (handler OdooHandler) BrowseTravelPackage(ctx echo.Context) error {
	uc := usecase.OdooUseCase{UcContract: handler.UseCaseContract}
	var res []viewmodel.TravelPackageOdooVm

	err := uc.Browse("travel.package", "", 0, 0, &res)


	return handler.SendResponse(ctx, res, nil, err)
}

func (handler OdooHandler) BrowseEquipmentPackage(ctx echo.Context) error{
	uc := usecase.OdooUseCase{UcContract: handler.UseCaseContract}
	var res []viewmodel.OdooEquipmentPackageVm

	err := uc.Browse("equipment.package", "", 0, 0, &res)


	return handler.SendResponse(ctx, res, nil, err)
}

func (handler OdooHandler) ReadTravelPackage(ctx echo.Context) error{
	ID,_ := strconv.Atoi(ctx.Param("id"))

	uc := usecase.OdooUseCase{UcContract: handler.UseCaseContract}
	var res []viewmodel.TravelPackageOdooVm

	fmt.Println("ini")
	err := uc.Read("travel.package", int64(ID), &res)
	fmt.Println(res[0].IsActive)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler OdooHandler) ReadEquipmentPackage(ctx echo.Context) error{
	ID,_ := strconv.Atoi(ctx.Param("id"))

	uc := usecase.OdooUseCase{UcContract: handler.UseCaseContract}
	var res []viewmodel.OdooEquipmentPackageVm

	err := uc.Read("equipment.package", int64(ID), &res)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler OdooHandler) ReadProductTemplate(ctx echo.Context) error{
	ID,_ := strconv.Atoi(ctx.Param("id"))

	uc := usecase.OdooUseCase{UcContract: handler.UseCaseContract}
	var res []viewmodel.OdooProductTemplateVm

	err := uc.Read("product.template", int64(ID), &res)
	fmt.Println(res[0].HotelOK)

	return handler.SendResponse(ctx, res, nil, err)
}
