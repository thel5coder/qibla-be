package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type ZakatRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route ZakatRoutes) RegisterRoute() {
	handler := handlers.ZakatHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	zakatRoute := route.RouteGroup.Group("/zakat")
	zakatRoute.Use(jwtMiddleware.JWTWithConfig)
	zakatRoute.GET("/place", handler.BrowseAll)

	userZakatHandler := handlers.UserZakatHandler{Handler: route.Handler}
	zakatRoute.GET("", userZakatHandler.Browse)
	zakatRoute.GET("/all", userZakatHandler.BrowseAll)
	zakatRoute.GET("/:id", userZakatHandler.Read)
	zakatRoute.PUT("/payment/:id", userZakatHandler.EditPayment)
	zakatRoute.POST("", userZakatHandler.Add)
	zakatRoute.DELETE("/:id", userZakatHandler.Delete)
}
