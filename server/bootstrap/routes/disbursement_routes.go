package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

// DisbursementRoutes ...
type DisbursementRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

// RegisterRoute ...
func (route DisbursementRoutes) RegisterRoute() {
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}
	disbursementRoute := route.RouteGroup.Group("/disbursement")
	disbursementRoute.Use(jwtMiddleware.JWTWithConfig)

	disbursementHandler := handlers.DisbursementHandler{Handler: route.Handler}
	disbursementRoute.GET("", disbursementHandler.Browse)
	disbursementRoute.GET("/all", disbursementHandler.BrowseAll)
	disbursementRoute.GET("/pdf/:id", disbursementHandler.PdfExport)
	disbursementRoute.POST("/request", disbursementHandler.Request)
	disbursementRoute.GET("/:id", disbursementHandler.Read)
}
