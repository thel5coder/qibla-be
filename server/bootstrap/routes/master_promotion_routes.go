package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type MasterPromotionRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route MasterPromotionRoutes) RegisterRoute(){
	handler := handlers.MasterPromotionHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	masterPromotionRoutes := route.RouteGroup.Group("/master-promotion")
	masterPromotionRoutes.Use(jwtMiddleware.JWTWithConfig)
	masterPromotionRoutes.GET("",handler.Browse)
	masterPromotionRoutes.GET("/all",handler.BrowseAll)
	masterPromotionRoutes.GET("/:id",handler.Read)
	masterPromotionRoutes.PUT("/:id",handler.Edit)
	masterPromotionRoutes.POST("",handler.Add)
	masterPromotionRoutes.DELETE("/:id",handler.Delete)
}
