package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/helpers/jwt"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

type UserHandler struct {
	Handler
}

func (handler UserHandler) Browse(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(search, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler UserHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler UserHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.UserRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID, input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler UserHandler) Add(ctx echo.Context) error {
	input := new(requests.UserRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler UserHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler UserHandler) GetCurrentUser(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.CustomClaims)

	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(user.Id)

	return handler.SendResponse(ctx, res, nil, err)
}
