package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type ZakatRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route ZakatRoutes) RegisterRoute(){
	handler := handlers.ZakatHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	zakatRoute := route.RouteGroup.Group("/zakat")
	zakatRoute.Use(jwtMiddleware.JWTWithConfig)
	zakatRoute.GET("/place",handler.BrowseAll)
}
