package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type CityRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route CityRoutes) RegisterRoute(){
	handler := handlers.CityHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	cityRoutes := route.RouteGroup.Group("/city")
	cityRoutes.Use(jwtMiddleware.JWTWithConfig)
	cityRoutes.GET("/all",handler.BrowseAllByProvince)
}
