package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type SettingProductRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route SettingProductRoutes) RegisterRoute() {
	handler := handlers.SettingProductHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	settingProductRoute := route.RouteGroup.Group("/setting-product")
	settingProductRoute.Use(jwtMiddleware.JWTWithConfig)
	settingProductRoute.GET("", handler.Browse)
	settingProductRoute.GET("/all", handler.BrowseAll)
	settingProductRoute.GET("/subscription-product", handler.BrowseSubscriptionProduct)
	settingProductRoute.GET("/webinar-website-product", handler.BrowseWebinarAndWebsiteProduct)
	settingProductRoute.GET("/product/:id", handler.ReadByProductID)
	settingProductRoute.GET("/:id", handler.Read)
	settingProductRoute.PUT("/:id", handler.Edit)
	settingProductRoute.POST("", handler.Add)
	settingProductRoute.DELETE("/:id", handler.Delete)
}
