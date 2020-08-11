package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type OdooRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route OdooRoutes) RegisterRoute(){
	handler := handlers.OdooHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	odooRoute := route.RouteGroup.Group("/odoo")
	odooRoute.Use(jwtMiddleware.JWTWithConfig)
	odooRoute.GET("/get-field/:objectName",handler.GetField)
	odooRoute.GET("/:objectName",handler.Browse)
}
