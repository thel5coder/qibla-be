package handlers

import (
	"errors"
	"github.com/labstack/echo"
)

type MeetPointHandler struct {
	Handler
}

//browse
func (handler MeetPointHandler) Browse(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, nil)
}

//browse passenger location
func (handler MeetPointHandler) BrowsePassengerLocation(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, errors.New("pq: sql no rows result set"))
}

//create
func (handler MeetPointHandler) Create(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, nil)
}
