package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
)

type CrmStoryHandler struct {
	Handler
}

func (handler CrmStoryHandler) BrowseAll(ctx echo.Context) error {
	uc := usecase.CrmStoryUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler CrmStoryHandler) ReadByPk(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.CrmStoryUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("id", ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler CrmStoryHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.CrmStoryRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.CrmStoryUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID, input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler CrmStoryHandler) Add(ctx echo.Context) error {
	input := new(requests.CrmStoryRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.CrmStoryUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler CrmStoryHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.CrmStoryUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
