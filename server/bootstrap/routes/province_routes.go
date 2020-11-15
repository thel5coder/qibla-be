package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type ProvinceRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route ProvinceRoutes) RegisterRoute() {
	handler := handlers.ProvinceHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	provinceRoutes := route.RouteGroup.Group("/province")
	provinceRoutes.Use(jwtMiddleware.JWTWithConfig)
	provinceRoutes.GET("/all", handler.BrowseAll)
}
