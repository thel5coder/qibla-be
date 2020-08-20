package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type CalendarRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route CalendarRoutes) RegisterRoute(){
	handler := handlers.CalendarHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	roleRoute := route.RouteGroup.Group("/calendar")
	roleRoute.Use(jwtMiddleware.JWTWithConfig)
	roleRoute.GET("",handler.BrowseByYearMonth)
	roleRoute.GET("/:id",handler.Read)
	roleRoute.PUT("/:id",handler.Edit)
	roleRoute.POST("",handler.Add)
	roleRoute.DELETE("/:id",handler.Delete)
}