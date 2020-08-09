package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type PromotionPackageRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route PromotionPackageRoutes) RegisterRoute(){
	handler := handlers.PromotionPackageHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	promotionPackageRoute := route.RouteGroup.Group("/promotion-package")
	promotionPackageRoute.Use(jwtMiddleware.JWTWithConfig)
	promotionPackageRoute.GET("",handler.Browse)
	promotionPackageRoute.GET("/:id",handler.Read)
	promotionPackageRoute.PUT("/:id",handler.Edit)
	promotionPackageRoute.POST("",handler.Add)
	promotionPackageRoute.DELETE("/:id",handler.Delete)
}
