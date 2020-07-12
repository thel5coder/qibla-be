package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type TermConditionRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route TermConditionRoutes) RegisterRoute() {
	handler := handlers.TermConditionHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	termConditionRoute := route.RouteGroup.Group("/term-condition")
	termConditionRoute.Use(jwtMiddleware.JWTWithConfig)
	termConditionRoute.GET("", handler.Browse)
	termConditionRoute.GET("/:id", handler.Read)
	termConditionRoute.PUT("/:id", handler.Edit)
	termConditionRoute.POST("", handler.Add)
	termConditionRoute.DELETE("/:id", handler.Delete)
}
