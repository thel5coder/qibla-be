package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type EnumOptionRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route EnumOptionRoutes) RegisterRoute() {
	handler := handlers.EnumHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	enumRoute := route.RouteGroup.Group("/options")
	enumRoute.Use(jwtMiddleware.JWTWithConfig)
	enumRoute.GET("/menu-permissions", handler.GetMenuPermissions)
	enumRoute.GET("/web-comprof-category", handler.GetWebComprofCategori)
	enumRoute.GET("/promotion-package", handler.GetPromotionPackage)
	enumRoute.GET("/platform", handler.GetPlatform)
	enumRoute.GET("/position-promotion", handler.GetPositionPromotion)
	enumRoute.GET("/subscription-type", handler.GetSubscription)
	enumRoute.GET("/price-unit", handler.GetPriceUnit)
	enumRoute.GET("/discount-type", handler.GetDiscountType)
	enumRoute.GET("/complaint-status", handler.GetComplaintStatus)
	enumRoute.GET("/type-zakat", handler.GetTypeZakat)
	enumRoute.GET("/remember-calender", handler.GetRemember)
	enumRoute.GET("/sex", handler.GetSex)
	enumRoute.GET("/marital-status", handler.GetMaritalStatus)
}
