package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/helper/master_zakat_helper"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

type MasterZakatHandler struct {
	Handler
}

func (handler MasterZakatHandler) Browse(ctx echo.Context) error {
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	filters := master_zakat_helper.SetFilterParams(ctx)

	uc := usecase.MasterZakatUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(filters, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler MasterZakatHandler) BrowseAll(ctx echo.Context) error {
	uc := usecase.MasterZakatUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler MasterZakatHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.MasterZakatUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler MasterZakatHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.MasterZakatRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.MasterZakatUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID, input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler MasterZakatHandler) Add(ctx echo.Context) error {
	input := new(requests.MasterZakatRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.MasterZakatUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler MasterZakatHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.MasterZakatUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
