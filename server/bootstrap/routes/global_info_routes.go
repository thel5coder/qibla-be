package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type GlobalInfoRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route GlobalInfoRoutes) RegisterRoute() {
	globalInfoCategoryHandler := handlers.GlobalInfoCategoryHandler{Handler: route.Handler}
	globalInfoCategorySettingHandler := handlers.GlobalInfoCategorySettingHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}
	globalInfoRoute := route.RouteGroup.Group("/global-info")

	globalInfoCategoryRoute := globalInfoRoute.Group("/category")
	globalInfoCategoryRoute.Use(jwtMiddleware.JWTWithConfig)
	globalInfoCategoryRoute.GET("", globalInfoCategoryHandler.Browse)
	globalInfoCategoryRoute.GET("/:id", globalInfoCategoryHandler.Read)
	globalInfoCategoryRoute.PUT("/:id", globalInfoCategoryHandler.Edit)
	globalInfoCategoryRoute.POST("", globalInfoCategoryHandler.Add)
	globalInfoCategoryRoute.DELETE("/:id", globalInfoCategoryHandler.Delete)

	globalInfoCategorySettingRoute := globalInfoRoute.Group("/category-setting")
	globalInfoCategorySettingRoute.Use(jwtMiddleware.JWTWithConfig)
	globalInfoCategorySettingRoute.GET("", globalInfoCategorySettingHandler.Browse)
	globalInfoCategorySettingRoute.GET("/category/:globalInfoCategory", globalInfoCategorySettingHandler.BrowseByGlobalInfoCategory)
	globalInfoCategorySettingRoute.GET("/:id", globalInfoCategorySettingHandler.Read)
	globalInfoCategorySettingRoute.PUT("/:id", globalInfoCategorySettingHandler.Edit)
	globalInfoCategorySettingRoute.POST("", globalInfoCategorySettingHandler.Add)
	globalInfoCategorySettingRoute.DELETE("/:id", globalInfoCategorySettingHandler.Delete)
}
