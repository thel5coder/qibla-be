package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/helpers/enums"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

type SettingProductHandler struct {
	Handler
}

func (handler SettingProductHandler) Browse(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(search, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler SettingProductHandler) BrowseSubscriptionProduct(ctx echo.Context) error{
	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res,err := uc.BrowseBy("mp.subscription_type",enums.KeySubscriptionEnum1,"=")

	return handler.SendResponse(ctx,res,nil,err)
}

func (handler SettingProductHandler) BrowseWebinarAndWebsiteProduct(ctx echo.Context) error{
	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res,err := uc.BrowseBy("mp.subscription_type",enums.KeySubscriptionEnum1,"<>")

	return handler.SendResponse(ctx,res,nil,err)
}

func (handler SettingProductHandler) BrowseAll(ctx echo.Context) error {
	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll()

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler SettingProductHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler SettingProductHandler) ReadByProductID(ctx echo.Context) error{
	ID := ctx.Param("id")

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("product_id",ID)

	return handler.SendResponse(ctx, res, nil, err)
}

func (handler SettingProductHandler) Edit(ctx echo.Context) error {
	input := new(requests.SettingProductRequest)
	ID := ctx.Param("id")

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	err := uc.Edit(ID, input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler SettingProductHandler) Add(ctx echo.Context) error {
	input := new(requests.SettingProductRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	err := uc.Add(input)

	return handler.SendResponse(ctx, nil, nil, err)
}

func (handler SettingProductHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	err := uc.Delete(ID)

	return handler.SendResponse(ctx, nil, nil, err)
}
