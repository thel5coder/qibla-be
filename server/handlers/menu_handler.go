package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

type MenuHandler struct {
	Handler
}

func (handler MenuHandler) Browse(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse("", search, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler MenuHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler MenuHandler) Edit(ctx echo.Context) error {
	inputs := new([]requests.EditMenuRequest)

	for _, input := range *inputs {
		if err := ctx.Bind(input); err != nil {
			return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
		}
		if err := handler.Validate.Struct(input); err != nil {
			return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
		}
	}

	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(inputs)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler MenuHandler) Add(ctx echo.Context) error {
	inputs := new(requests.MenuRequest)

	if err := ctx.Bind(inputs); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	for _,input := range inputs.Menus{
		if err := handler.Validate.Struct(input); err != nil {
			return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
		}
	}
	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(inputs)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler MenuHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler MenuHandler) GetMenuID(ctx echo.Context) error {
	parentID := ctx.QueryParam("parentId")

	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.GetMenuID(parentID)

	return handler.SendResponse(ctx, res, nil, err)
}
