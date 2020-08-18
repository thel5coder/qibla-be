package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type ContactRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route ContactRoutes) RegisterRoute(){
	handler := handlers.ContactHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	contactRoute := route.RouteGroup.Group("/contact")
	contactRoute.Use(jwtMiddleware.JWTWithConfig)
	contactRoute.GET("", handler.Browse)
	contactRoute.GET("/all", handler.BrowseAll)
	contactRoute.GET("/:id", handler.Read)
	contactRoute.PUT("/:id", handler.Edit)
	contactRoute.POST("", handler.Add)
	contactRoute.DELETE("/:id", handler.Delete)
}
