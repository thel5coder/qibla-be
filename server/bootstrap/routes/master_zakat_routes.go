package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type MasterZakatRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route MasterZakatRoutes) RegisterRoute() {
	handler := handlers.MasterZakatHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	masterZakat := route.RouteGroup.Group("/master-zakat")
	masterZakat.Use(jwtMiddleware.JWTWithConfig)
	masterZakat.GET("", handler.Browse)
	masterZakat.GET("/all", handler.BrowseAll)
	masterZakat.GET("/:id", handler.Read)
	masterZakat.PUT("/:id", handler.Edit)
	masterZakat.POST("", handler.Add)
	masterZakat.DELETE("/:id", handler.Delete)
}
