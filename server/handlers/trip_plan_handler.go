package handlers

import (
	"errors"
	"github.com/labstack/echo"
)

type TripPlanHandler struct {
	Handler
}

func(handler TripPlanHandler) Read(ctx echo.Context) error{
	return handler.SendResponse(ctx,nil,nil,errors.New("pq:sql no rows result set"))
}
