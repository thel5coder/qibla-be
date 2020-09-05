package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type VideoKajianRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route VideoKajianRoutes) RegisterRoute(){
	handler := handlers.VideoKajianHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	videoKajianRoute:= route.RouteGroup.Group("/video-kajian")
	videoKajianRoute.Use(jwtMiddleware.JWTWithConfig)
	videoKajianRoute.GET("",handler.Browse)
	videoKajianRoute.GET("/:id",handler.Read)
}