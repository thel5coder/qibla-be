package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type TourPackageRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route TourPackageRoutes) RegisterRoute(){
	_ = handlers.TourPackageHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	tourPackageRoute := route.RouteGroup.Group("/tour-package")
	tourPackageRoute.Use(jwtMiddleware.JWTWithConfig)
}
