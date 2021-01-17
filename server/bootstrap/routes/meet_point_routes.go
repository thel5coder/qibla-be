package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type MeetPointRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route MeetPointRoutes) RegisterRoute() {
	handler := handlers.MeetPointHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	meetPointRoutes := route.RouteGroup.Group("/meet-point")
	meetPointRoutes.Use(jwtMiddleware.JWTWithConfig)
	meetPointRoutes.GET("", handler.Browse)
	meetPointRoutes.GET("/passenger-location", handler.BrowsePassengerLocation)
	meetPointRoutes.POST("", handler.Create)
}
