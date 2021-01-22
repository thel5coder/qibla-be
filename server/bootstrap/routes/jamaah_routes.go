package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type JamaahRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route JamaahRoutes) RegisterRoute() {
	handler := handlers.JamaahHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	jamaahRoutes := route.RouteGroup.Group("/jamaah")
	jamaahRoutes.Use(jwtMiddleware.JWTWithConfig)
	jamaahRoutes.POST("/health-status", handler.EditHealthStatus)
	jamaahRoutes.GET("/:id", handler.Read)
}
