package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type VideoContentRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func(route VideoContentRoutes) RegisterRoute(){
	handler := handlers.VideoContentHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	roleRoute := route.RouteGroup.Group("/video-content")
	roleRoute.Use(jwtMiddleware.JWTWithConfig)
	roleRoute.GET("",handler.Browse)
	roleRoute.POST("",handler.Add)
}
