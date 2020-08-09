package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type RoleRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route RoleRoutes) RegisterRoute(){
	handler := handlers.RoleHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	roleRoute := route.RouteGroup.Group("/role")
	roleRoute.Use(jwtMiddleware.JWTWithConfig)
	roleRoute.GET("",handler.Browse)
	roleRoute.GET("/:id",handler.Read)
	roleRoute.PUT("/:id",handler.Edit)
	roleRoute.POST("",handler.Add)
	roleRoute.DELETE("/:id",handler.Delete)
	roleRoute.GET("/res",handler.GetRes)
}
