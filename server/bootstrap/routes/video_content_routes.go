package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type VideoContentRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route VideoContentRoutes) RegisterRoute() {
	handler := handlers.VideoContentHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	videoContentRoute := route.RouteGroup.Group("/video-content")
	videoContentRoute.Use(jwtMiddleware.JWTWithConfig)
	videoContentRoute.GET("", handler.Browse)
	videoContentRoute.GET("/:id", handler.ReaByPk)
	videoContentRoute.PUT("/:id", handler.Edit)
	videoContentRoute.POST("", handler.Add)
	videoContentRoute.DELETE("/:id", handler.Delete)
}
