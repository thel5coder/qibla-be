package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
	"strconv"
)

// UserZakatHandler ...
type UserZakatHandler struct {
	Handler
}

// Browse ..
func (handler UserZakatHandler) Browse(ctx echo.Context) error {
	search := ctx.QueryParam("search")
	order := ctx.QueryParam("order")
	sort := ctx.QueryParam("sort")
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	page, _ := strconv.Atoi(ctx.QueryParam("page"))

	uc := usecase.UserZakatUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(search, order, sort, page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

// BrowseAll ...
func (handler UserZakatHandler) BrowseAll(ctx echo.Context) error {
	uc := usecase.UserZakatUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.BrowseAll()

	return handler.SendResponse(ctx, res, nil, err)
}

// Read ...
func (handler UserZakatHandler) Read(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.UserZakatUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadByPk(ID)

	return handler.SendResponse(ctx, res, nil, err)
}

// EditPayment ...
func (handler UserZakatHandler) EditPayment(ctx echo.Context) error {
	input := new(requests.UserZakatPaymentRequest)
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

	inputBody := requests.UserZakatRequest{
		PaymentMethodCode: input.PaymentMethodCode,
		BankName:          input.BankName,
	}
	uc := usecase.UserZakatUseCase{UcContract: handler.UseCaseContract}
	err = uc.EditPaymentMethod(ID, &inputBody)
	if err != nil {
		handler.UseCaseContract.TX.Rollback()
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	handler.UseCaseContract.TX.Commit()

	return handler.SendResponse(ctx, nil, nil, err)
}

// Add ...
func (handler UserZakatHandler) Add(ctx echo.Context) error {
	input := new(requests.UserZakatRequest)

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

	uc := usecase.UserZakatUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Add(input)
	if err != nil {
		handler.UseCaseContract.TX.Rollback()
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	handler.UseCaseContract.TX.Commit()

	return handler.SendResponse(ctx, res, nil, err)
}

// Delete ...
func (handler UserZakatHandler) Delete(ctx echo.Context) error {
	ID := ctx.Param("id")

	var err error
	handler.UseCaseContract.TX, err = handler.UseCaseContract.DB.Begin()
	if err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}

	uc := usecase.UserZakatUseCase{UcContract: handler.UseCaseContract}
	err = uc.Delete(ID)
	if err != nil {
		handler.UseCaseContract.TX.Rollback()
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	handler.UseCaseContract.TX.Commit()

	return handler.SendResponse(ctx, nil, nil, err)
}
