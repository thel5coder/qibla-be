package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type MasterProductRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route MasterProductRoutes) RegisterRoute() {
	handler := handlers.MasterProductHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	masterProductRoute := route.RouteGroup.Group("/master-product")
	masterProductRoute.Use(jwtMiddleware.JWTWithConfig)
	masterProductRoute.GET("", handler.Browse)
	masterProductRoute.GET("/all", handler.BrowseAll)
	masterProductRoute.GET("/extra-product",handler.BrowseExtraProducts)
	masterProductRoute.GET("/:id", handler.Read)
	masterProductRoute.PUT("/:id", handler.Edit)
	masterProductRoute.POST("", handler.Add)
	masterProductRoute.DELETE("/:id", handler.Delete)
}
