package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type MenuRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route MenuRoutes) RegisterRoute() {
	handler := handlers.MenuHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	menuRoute := route.RouteGroup.Group("/menu")
	menuRoute.Use(jwtMiddleware.JWTWithConfig)
	menuRoute.GET("", handler.Browse)
	menuRoute.GET("/all", handler.BrowseAllTree)
	menuRoute.GET("/:id", handler.Read)
	menuRoute.POST("/edit", handler.Edit)
	menuRoute.POST("/add", handler.Add)
	menuRoute.DELETE("/:id", handler.Delete)
	menuRoute.GET("/get-menu-id", handler.GetMenuID)
}
