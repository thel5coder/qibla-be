package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/pkg/jwt"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
)

type AuthenticationHandler struct {
	Handler
}

func (handler AuthenticationHandler) Login(ctx echo.Context) error {
	input := new(requests.LoginRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Login(input.UserName, input.Password, input.FcmDeviceToken)
	if err != nil {
		return handler.SendResponseUnauthorized(ctx, err)
	}

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler AuthenticationHandler) RegisterByOauth(ctx echo.Context) error {
	input := new(requests.RegisterByOauthRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	res,err := uc.RegisterByOauth(input)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler AuthenticationHandler) RegisterJamaahByEmail(ctx echo.Context) error {
	input := new(requests.RegisterByMailRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	err := uc.RegisterByEmail(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler AuthenticationHandler) ForgotPassword(ctx echo.Context) error {
	input := new(requests.ForgotPasswordRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	err := uc.ForgotPassword(input.Email)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler AuthenticationHandler) SetPin(ctx echo.Context) error {
	input := new(requests.SetPinRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	user := ctx.Get("user").(*jwt.CustomClaims)
	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}

	userUc := usecase.UserUseCase{UcContract: uc.UcContract}
	err := userUc.EditPin(user.Id, input.Pin)

	return handler.SendResponse(ctx, nil, nil, err)
}

func(handler AuthenticationHandler) SetFingerPrintStatus(ctx echo.Context) error{
	uc := usecase.UserUseCase{UcContract: handler.UseCaseContract}
	user := ctx.Get("user").(*jwt.CustomClaims)

	err := uc.EditFingerPrintStatus(user.Id,true)

	return handler.SendResponse(ctx, nil, nil, err)
}

func(handler AuthenticationHandler) ActivationUserByCode(ctx echo.Context) error{
	code := ctx.QueryParam("code")

	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	err := uc.ActivationUserByCode(code)

	return handler.SendResponse(ctx,nil,nil,err)
}
