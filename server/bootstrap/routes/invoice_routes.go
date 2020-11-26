package routes

import (
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"

	"github.com/labstack/echo"
)

// InvoiceRoutes ...
type InvoiceRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

// RegisterRoute ...
func (route InvoiceRoutes) RegisterRoute() {
	handler := handlers.InvoiceHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	InvoiceRoute := route.RouteGroup.Group("/invoices")
	InvoiceRoute.Use(jwtMiddleware.JWTWithConfig)
	InvoiceRoute.GET("", handler.BrowseInvoicesHandler)
}
