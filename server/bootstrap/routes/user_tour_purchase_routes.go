package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type UserTourPurchaseRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route UserTourPurchaseRoutes) RegisterRoute() {
	handler := handlers.UserTourPurchaseHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	userTourPurchaseRoutes := route.RouteGroup.Group("/package-purchase")
	userTourPurchaseRoutes.Use(jwtMiddleware.JWTWithConfig)
	userTourPurchaseRoutes.GET("", handler.Browse)
	userTourPurchaseRoutes.GET("/:id", handler.Read)
	userTourPurchaseRoutes.POST("", handler.CreatePurchase)
	userTourPurchaseRoutes.POST("/registrant", handler.CreatePassenger)
	userTourPurchaseRoutes.POST("/passenger-document", handler.CreateDocument)
	userTourPurchaseRoutes.POST("/payment", handler.CreatePayment)
	userTourPurchaseRoutes.POST("/request-cancel",handler.RequestCancelPurchase)
}
