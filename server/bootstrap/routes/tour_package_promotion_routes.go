package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type TourPackagePromotionRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route TourPackagePromotionRoutes) RegisterRoute() {
	handler := handlers.TourPackagePromotionHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	tourPackagePromotionRoutes := route.RouteGroup.Group("/promotion-today")
	tourPackagePromotionRoutes.Use(jwtMiddleware.JWTWithConfig)
	tourPackagePromotionRoutes.GET("", handler.Browse)
	tourPackagePromotionRoutes.GET("/:id", handler.ReadByPk)
}
