package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
	"qibla-backend/server/requests"
	"qibla-backend/usecase"
)

type UserTourPurchaseHandler struct {
	Handler
}

//browse
func (handler UserTourPurchaseHandler) Browse(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, nil)
}

//read
func (handler UserTourPurchaseHandler) Read(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, nil)
}

//create tour purchase
func (handler UserTourPurchaseHandler) CreatePurchase(ctx echo.Context) (err error) {
	input := new(requests.CreatePurchaseRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	handler.UseCaseContract.TX, err = handler.Db.Begin()
	if err != nil {
		handler.UseCaseContract.TX.Rollback()

		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	uc := usecase.UserTourPurchaseUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Add(input)
	if err != nil {
		handler.UseCaseContract.TX.Rollback()

		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	handler.UseCaseContract.TX.Commit()

	return handler.SendResponse(ctx, res, nil, nil)
}

//create passenger/participant
func (handler UserTourPurchaseHandler) CreatePassenger(ctx echo.Context) (err error) {
	input := new(requests.CreatePassengerRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	handler.UseCaseContract.TX, err = handler.Db.Begin()
	if err != nil {
		handler.UseCaseContract.TX.Rollback()

		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	uc := usecase.UserTourPurchaseUseCase{UcContract:handler.UseCaseContract}
	res,err := uc.CreatePassenger(input)
	if err != nil {
		handler.UseCaseContract.TX.Rollback()

		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	handler.UseCaseContract.TX.Commit()

	return handler.SendResponse(ctx, res, nil, nil)
}

//create document
func (handler UserTourPurchaseHandler) CreateDocument(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, nil)
}

//create payment
func (handler UserTourPurchaseHandler) CreatePayment(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, nil)
}

//cancel purchase
func (handler UserTourPurchaseHandler) RequestCancelPurchase(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, nil)
}
