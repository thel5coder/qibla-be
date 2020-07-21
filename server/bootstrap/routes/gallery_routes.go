package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type GalleryRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route GalleryRoutes) RegisterRoute(){
	handler := handlers.GalleryHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	roleRoute := route.RouteGroup.Group("/gallery")
	roleRoute.Use(jwtMiddleware.JWTWithConfig)
	roleRoute.GET("",handler.Browse)
	roleRoute.GET("/:id",handler.Read)
	roleRoute.PUT("/:id",handler.Edit)
	roleRoute.POST("",handler.Add)
	roleRoute.DELETE("/:id",handler.Delete)
}
