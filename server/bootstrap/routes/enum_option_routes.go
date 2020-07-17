package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type EnumOptionRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route EnumOptionRoutes) RegisterRoute(){
	handler := handlers.EnumHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	enumRoute := route.RouteGroup.Group("/options")
	enumRoute.Use(jwtMiddleware.JWTWithConfig)
	enumRoute.GET("/menu-permissions",handler.GetMenuPermissions)
}
