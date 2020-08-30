package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

type VideoContentHandler struct {
	Handler
}

func (handler VideoContentHandler) Browse(ctx echo.Context) error {
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.VideoContentUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler VideoContentHandler) ReaByPk(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.VideoContentUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("id", ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler VideoContentHandler) Edit(ctx echo.Context) error{
	ID := ctx.Param("id")
	input := new(requests.VideoContentRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.VideoContentUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID,input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler VideoContentHandler) Add(ctx echo.Context) error {
	input := new(requests.VideoContentRequest)
	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.VideoContentUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler VideoContentHandler) Delete(ctx echo.Context) error{
	ID := ctx.Param("id")

	uc := usecase.VideoContentUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
