package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type OdooMasterPackageRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route OdooMasterPackageRoutes) RegisterRoute() {
	handler := handlers.OdooMasterPackageHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	odooMasterPackageRoutes := route.RouteGroup.Group("/odoo-master-package")
	odooMasterPackageRoutes.Use(jwtMiddleware.JWTWithConfig)
	odooMasterPackageRoutes.GET("/all", handler.BrowseAll)
	odooMasterPackageRoutes.GET("/:id", handler.ReadByPk)
}
