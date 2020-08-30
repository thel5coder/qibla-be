package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
)

type CrmBoardHandler struct {
	Handler
}

func (handler CrmBoardHandler) Browse(ctx echo.Context) error {
	crmStoryID := ctx.QueryParam("crmStoryId")

	uc := usecase.CrmBoardUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseByCrmStoryID(crmStoryID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler CrmBoardHandler) Edit(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.CrmBoardRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.CrmBoardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID, input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler CrmBoardHandler) EditBoardStory(ctx echo.Context) error {
	ID := ctx.Param("id")
	input := new(requests.CrmBoardEditStoryRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.CrmBoardUseCase{UcContract: handler.UseCaseContract}
	err := uc.EditBoardStory(ID, input.CrmStoryID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler CrmBoardHandler) Add(ctx echo.Context) error {
	input := new(requests.CrmBoardRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.CrmBoardUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}
