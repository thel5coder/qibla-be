package handlers

import (
	"errors"
	"github.com/labstack/echo"
)

type TravelInformationHandler struct {
	Handler
}

//trip itinerary
func (handler TravelInformationHandler) BrowseTripItinerary(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, errors.New("pq:sql no rows result set"))
}

//bus list
func (handler TravelInformationHandler) BrowseBus(ctx echo.Context) error {
	var res []map[string]interface{}
	{
	}

	for i := 1; i < 3; i++ {
		res = append(res, map[string]interface{}{
			"id":          i,
			"name":        "Bus ",
			"seat_number": i,
		})
	}

	return handler.SendResponse(ctx, res, nil, nil)
}

//browse airlines
func (handler TravelInformationHandler) BrowseAirlines(ctx echo.Context) error {
	var res = []map[string]interface{}{}

	res = append(res, map[string]interface{}{
		"id":             1,
		"departure_from": "Soekarno Hatta",
		"arrival_to":     "Makkah International Airport",
		"departure_date": "2020-02-11",
		"ticket_number":  0,
		"e_ticket":       nil,
	})

	res = append(res, map[string]interface{}{
		"id":             2,
		"departure_from": "Makkah International Airport",
		"arrival_to":     "Soekarno Hatta",
		"departure_date": "2020-02-20",
		"ticket_number":  0,
		"e_ticket":       nil,
	})

	return handler.SendResponse(ctx, res, nil, nil)
}

//meals
func(handler TravelInformationHandler) BrowseMeals(ctx echo.Context) error{
	return handler.SendResponse(ctx, nil, nil, nil)
}

//group list
func(handler TravelInformationHandler) BrowseGroup(ctx echo.Context) error {
	return handler.SendResponse(ctx, nil, nil, nil)
}

//passenger file
func(handler TravelInformationHandler) ReadPassengerFile(ctx echo.Context) error{
	return handler.SendResponse(ctx, nil, nil, nil)
}



