package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type FileRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route FileRoutes) RegisterRoute(){
	handler := handlers.FileHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	fileRoute:=route.RouteGroup.Group("/file")
	fileRoute.Use(jwtMiddleware.JWTWithConfig)
	fileRoute.GET("/:id",handler.Read)
	fileRoute.POST("",handler.Add)
}
