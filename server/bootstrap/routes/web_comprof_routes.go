package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type WebComprofRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route WebComprofRoutes) RegisterRoute(){
	webComprofCategoryHandler := handlers.WebComprofCategoryHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}
	webComprofRoute := route.RouteGroup.Group("/web-comprof")

	webComprofCategoryRoute := webComprofRoute.Group("/category")
	webComprofCategoryRoute.Use(jwtMiddleware.JWTWithConfig)
	webComprofCategoryRoute.GET("", webComprofCategoryHandler.Browse)
	webComprofCategoryRoute.GET("/:id", webComprofCategoryHandler.Read)
	webComprofCategoryRoute.PUT("/:id", webComprofCategoryHandler.Edit)
	webComprofCategoryRoute.POST("", webComprofCategoryHandler.Add)
	webComprofCategoryRoute.DELETE("/:id", webComprofCategoryHandler.Delete)
}
