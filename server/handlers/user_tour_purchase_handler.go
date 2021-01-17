package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"qibla-backend/server/requests"
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
func (handler UserTourPurchaseHandler) CreatePurchase(ctx echo.Context) error {

	return handler.SendResponse(ctx, uuid.NewV4(), nil, nil)
}

//create passenger/participant
func (handler UserTourPurchaseHandler) CreatePassenger(ctx echo.Context) error {
	input := new(requests.TourPurchaseCreatePassengerRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	var res []map[string]interface{}
	for _, passenger := range input.Passengers {
		res = append(res, map[string]interface{}{
			"id":   uuid.NewV4(),
			"name": passenger.Name,
		})
	}

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
