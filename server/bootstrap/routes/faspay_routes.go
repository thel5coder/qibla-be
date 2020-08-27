package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type FaspayRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func(route FaspayRoutes) RegisterRoute(){
	handler := handlers.FasPayHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}


	faspayRoute := route.RouteGroup.Group("/faspay")
	faspayRoute.Use(jwtMiddleware.JWTWithConfig)
	faspayRoute.GET("/get-payment-method",handler.GetLisPaymentMethods)
}
