package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type GlobalInfoRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route GlobalInfoRoutes) RegisterRoute(){
	globalInfoCategoryHandler := handlers.GlobalInfoCategoryHandler{Handler: route.Handler}
	globalInfoCategorySettinghandler := handlers.GlobalInfoCategorySettingHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}
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
	globalInfoCategorySettingRoute.GET("", globalInfoCategorySettinghandler.Browse)
	globalInfoCategorySettingRoute.GET("/:id", globalInfoCategorySettinghandler.Read)
	globalInfoCategorySettingRoute.PUT("/:id", globalInfoCategorySettinghandler.Edit)
	globalInfoCategorySettingRoute.POST("", globalInfoCategorySettinghandler.Add)
	globalInfoCategorySettingRoute.DELETE("/:id", globalInfoCategorySettinghandler.Delete)
}
