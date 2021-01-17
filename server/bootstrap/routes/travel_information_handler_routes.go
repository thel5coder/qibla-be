package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type TravelInformationRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route TravelInformationRoutes) RegisterRoute() {
	handler := handlers.TravelInformationHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	travelInformationRoutes := route.RouteGroup.Group("/travel-info")
	travelInformationRoutes.Use(jwtMiddleware.JWTWithConfig)
	travelInformationRoutes.GET("/trip-itinerary", handler.BrowseTripItinerary)
	travelInformationRoutes.GET("/transportation", handler.BrowseBus)
	travelInformationRoutes.GET("/airlines", handler.BrowseAirlines)
	travelInformationRoutes.GET("/meals", handler.BrowseMeals)
	travelInformationRoutes.GET("/group-list", handler.BrowseGroup)
	travelInformationRoutes.GET("/passenger-file/:fileType", handler.ReadPassengerFile)
}
