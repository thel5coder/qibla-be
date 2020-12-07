package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/pkg/enums"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

// SettingProductHandler ...
type SettingProductHandler struct {
	Handler
}

// Browse ..
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

// BrowseSubscriptionProduct ..
func (handler SettingProductHandler) BrowseSubscriptionProduct(ctx echo.Context) error {
	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseBy("mp.subscription_type", enums.KeySubscriptionEnum1, "=")

	return handler.SendResponse(ctx, res, nil, err)
}

// BrowseWebinarAndWebsiteProduct ...
func (handler SettingProductHandler) BrowseWebinarAndWebsiteProduct(ctx echo.Context) error {
	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseBy("mp.subscription_type", enums.KeySubscriptionEnum1, "<>")

	return handler.SendResponse(ctx, res, nil, err)
}

// BrowseAll ...
func (handler SettingProductHandler) BrowseAll(ctx echo.Context) error {
	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll()

	return handler.SendResponse(ctx, res, nil, err)
}

// Read ...
func (handler SettingProductHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

// ReadByProductID ...
func (handler SettingProductHandler) ReadByProductID(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("product_id", ID)

	return handler.SendResponse(ctx, res, nil, err)
}

// Edit ...
func (handler SettingProductHandler) Edit(ctx echo.Context) error {
	input := new(requests.SettingProductRequest)
	ID := ctx.Param("id")

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	var err error
	handler.UseCaseContract.TX, err = handler.UseCaseContract.DB.Begin()
	if err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	err = uc.Edit(ID, input)
	if err != nil {
		handler.UseCaseContract.TX.Rollback()
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	handler.UseCaseContract.TX.Commit()

	return handler.SendResponse(ctx, nil, nil, err)
}

// Add ...
func (handler SettingProductHandler) Add(ctx echo.Context) error {
	input := new(requests.SettingProductRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	var err error
	handler.UseCaseContract.TX, err = handler.UseCaseContract.DB.Begin()
	if err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	err = uc.Add(input)
	if err != nil {
		handler.UseCaseContract.TX.Rollback()
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	handler.UseCaseContract.TX.Commit()

	return handler.SendResponse(ctx, nil, nil, err)
}

// Delete ...
func (handler SettingProductHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")

	var err error
	handler.UseCaseContract.TX, err = handler.UseCaseContract.DB.Begin()
	if err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	uc := usecase.SettingProductUseCase{UcContract: handler.UseCaseContract}
	err = uc.Delete(ID)
	if err != nil {
		handler.UseCaseContract.TX.Rollback()
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	handler.UseCaseContract.TX.Commit()

	return handler.SendResponse(ctx, nil, nil, err)
}
