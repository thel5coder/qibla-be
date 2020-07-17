package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type UserRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route UserRoutes) RegisterRoute() {
	userHandler := handlers.UserHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	userRoute := route.RouteGroup.Group("/user")
	userRoute.Use(jwtMiddleware.JWTWithConfig)
	userRoute.GET("", userHandler.Browse)
	userRoute.GET("/:id", userHandler.Read)
	userRoute.PUT("/:id", userHandler.Edit)
	userRoute.POST("", userHandler.Add)
	userRoute.DELETE("/:id", userHandler.Delete)
}