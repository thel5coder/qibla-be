package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
)

type CalendarHandler struct {
	Handler
}

func (handler CalendarHandler) BrowseByYearMonth(ctx echo.Context) error {
	yearMonth := ctx.QueryParam("yearMonth")

	uc := usecase.CalendarUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByYearMonth(yearMonth)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler CalendarHandler) Read(ctx echo.Context) error{
	ID := ctx.Param("id")

	uc := usecase.CalendarUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler CalendarHandler) Edit(ctx echo.Context) error{
	ID := ctx.Param("id")
	input := new(requests.CalendarRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.CalendarUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID,input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler CalendarHandler) Add(ctx echo.Context) error{
	input := new(requests.CalendarRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.CalendarUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler CalendarHandler) Delete(ctx echo.Context) error{
	ID := ctx.Param("id")

	uc := usecase.CalendarUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}
