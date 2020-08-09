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

	promotionRoute := route.RouteGroup.Group("/promotion-package")
	promotionRoute.Use(jwtMiddleware.JWTWithConfig)
	promotionRoute.GET("",handler.Browse)
	promotionRoute.GET("/:id",handler.Read)
	promotionRoute.PUT("/:id",handler.Edit)
	promotionRoute.POST("/:id",handler.Add)
	promotionRoute.DELETE("/:id",handler.Delete)
}