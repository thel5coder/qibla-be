package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

// FlipRoutes ...
type FlipRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

// RegisterRoute ...
func (route FlipRoutes) RegisterRoute() {
	handler := handlers.FlipHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	flipRoute := route.RouteGroup.Group("/flip")
	flipRoute.Use(jwtMiddleware.JWTWithConfig)
	flipRoute.GET("/bank", handler.GetBank)
	flipRoute.GET("/bank/:code", handler.GetBankByCode)

	flipRouteNoAuth := route.RouteGroup.Group("/flip")
	flipRouteNoAuth.POST("/disbursement/callback", handler.DisbursementCallback)
}
