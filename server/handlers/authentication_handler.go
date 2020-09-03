package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/helpers/jwt"
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
	res, err := uc.Login(input.UserName, input.Password)
	if err != nil {
		return handler.SendResponseUnauthorized(ctx, err)
	}

	return handler.SendResponse(ctx, res, nil, nil)
}

func (handler AuthenticationHandler) RegisterByGmail(ctx echo.Context) error {
	input := new(requests.RegisterByGmailRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.AuthenticationUseCase{UcContract: handler.UseCaseContract}
	err := uc.RegisterByGmail(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler AuthenticationHandler) RegisterJaamaahByEmail(ctx echo.Context) error {
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

	jamaahUc := usecase.JamaahUseCase{UcContract: uc.UcContract}
	err := jamaahUc.EditPin(user.Id, input.Pin)

	return handler.SendResponse(ctx, nil, nil, err)
}
