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

func (handler VideoContentHandler) Browse(ctx echo.Context) error{
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.VideoContentUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler VideoContentHandler) Add(ctx echo.Context) error{
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
