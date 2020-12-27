package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
)

type SatisfactionCategoryHandler struct {
	Handler
}

func (handler SatisfactionCategoryHandler) Browse(ctx echo.Context) error {
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	filters := make(map[string]interface{})

	if ctx.QueryParam("name") != ""{
		filters["name"]=ctx.QueryParam("name")
	}

	if ctx.QueryParam("description") !=""{
		filters["description"] = ctx.QueryParam("description")
	}

	if ctx.QueryParam("updated_at") != ""{
		filters["updated_at"]=ctx.QueryParam("updated_at")
	}

	uc := usecase.SatisfactionCategoryUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAllBy(filters,order,sort)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler SatisfactionCategoryHandler) ReadByPk(ctx echo.Context) error {
	ID := ctx.Param("id")
	uc := usecase.SatisfactionCategoryUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("id", ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler SatisfactionCategoryHandler) Store(ctx echo.Context) error {
	input := new(requests.SatisfactionCategoryRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.SatisfactionCategoryUseCase{UcContract: handler.UseCaseContract}
	err := uc.Store(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler SatisfactionCategoryHandler) DeleteByPk(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.SatisfactionCategoryUseCase{UcContract: handler.UseCaseContract}
	err := uc.DeleteByPk(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
