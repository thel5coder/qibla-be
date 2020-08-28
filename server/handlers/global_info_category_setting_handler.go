package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

type GlobalInfoCategorySettingHandler struct {
	Handler
}

func (handler GlobalInfoCategorySettingHandler) Browse(ctx echo.Context) error{
	search := ctx.QueryParam("search")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.GlobalInfoCategorySettingUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse("",search, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler GlobalInfoCategorySettingHandler) BrowseByGlobalInfoCategory(ctx echo.Context) error{
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	globalInfoCategory := ctx.Param("globalInfoCategory")

	uc := usecase.GlobalInfoCategorySettingUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(globalInfoCategory,"", order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler GlobalInfoCategorySettingHandler) Read(ctx echo.Context) error{
	ID := ctx.Param("id")

	uc := usecase.GlobalInfoCategorySettingUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler GlobalInfoCategorySettingHandler) Edit(ctx echo.Context) error{
	ID := ctx.Param("id")
	input := new(requests.GlobalInfoCategorySettingRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.GlobalInfoCategorySettingUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID, input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler GlobalInfoCategorySettingHandler) Add(ctx echo.Context) error{
	input := new(requests.GlobalInfoCategorySettingRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.GlobalInfoCategorySettingUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler GlobalInfoCategorySettingHandler) Delete(ctx echo.Context) error{
	ID := ctx.Param("id")

	uc := usecase.GlobalInfoCategorySettingUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
