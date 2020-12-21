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
	parentId := ctx.QueryParam("parentId")

	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(parentId, search, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler MenuHandler) BrowseAllTree(ctx echo.Context) error {
	isActive := ctx.QueryParam("is_active")
	var isActiveBool bool
	if isActive == "true" {
		isActiveBool = true
	}else{
		isActiveBool = false
	}
	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAllBy("m.parent_id", "", "=",isActiveBool)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler MenuHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("m.id", ID, "=")

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler MenuHandler) Edit(ctx echo.Context) error {
	inputs := new(requests.EditMenuRequest)

	if err := ctx.Bind(inputs); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	for _, input := range inputs.Menus {
		if err := handler.Validate.Struct(input); err != nil {
			return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
		}
	}

	uc := usecase.MenuUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(inputs)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler MenuHandler) Add(ctx echo.Context) error {
	inputs := new(requests.AddMenuRequest)

	if err := ctx.Bind(inputs); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	for _, input := range inputs.Menus {
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
