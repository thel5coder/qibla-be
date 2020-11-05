package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type AdminRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route AdminRoutes) RegisterRoute() {
	handler := handlers.AdminHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	userRoute := route.RouteGroup.Group("/admin")
	userRoute.Use(jwtMiddleware.JWTWithConfig)
	userRoute.GET("", handler.Browse)
	userRoute.GET("/:id", handler.Read)
	userRoute.PUT("/:id", handler.Edit)
	userRoute.POST("", handler.Add)
	userRoute.DELETE("/:id", handler.Delete)
}
