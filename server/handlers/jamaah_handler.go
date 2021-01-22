package handlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/pkg/jwt"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
)

type JamaahHandler struct {
	Handler
}

//jamaah status

func (handler JamaahHandler) EditHealthStatus(ctx echo.Context) error {
	input := new(requests.EditJamaahStatusRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	return handler.SendResponse(ctx, nil, nil, nil)
}

func (handler JamaahHandler) Read(ctx echo.Context) error {
	_ = ctx.Param("id")

	return handler.SendResponse(ctx, nil, nil, nil)
}

func (handler JamaahHandler) ReadProfile(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.CustomClaims)
	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("u.id", user.Id)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler JamaahHandler) EditProfile(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.CustomClaims)
	input := new(requests.EditProfileRequest)
	fmt.Println(user.Id)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.JamaahUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(input, user.Id)

	return handler.SendResponse(ctx, nil, nil, err)
}
