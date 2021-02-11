package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type PromotionRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route PromotionRoutes) RegisterRoute() {
	handler := handlers.SettingPromotionHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	promotionRoute := route.RouteGroup.Group("/promotion")
	promotionRoute.Use(jwtMiddleware.JWTWithConfig)
	promotionRoute.GET("", handler.Browse)
	promotionRoute.GET("/all", handler.BrowseAll)
	promotionRoute.GET("/:id", handler.Read)
	promotionRoute.PUT("/:id", handler.Edit)
	promotionRoute.POST("", handler.Add)
	promotionRoute.DELETE("/:id", handler.Delete)
}
