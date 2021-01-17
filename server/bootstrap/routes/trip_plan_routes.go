package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type TripPlanRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route TripPlanRoutes) RegisterRoute() {
	handler := handlers.TripPlanHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	tripPlanRoutes := route.RouteGroup.Group("/trip-plan")
	tripPlanRoutes.Use(jwtMiddleware.JWTWithConfig)
	tripPlanRoutes.GET("/:id", handler.Read)
}
