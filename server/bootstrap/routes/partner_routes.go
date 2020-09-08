package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type PartnerRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route PartnerRoutes) RegisterRoute() {
	handler := handlers.PartnerHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	partnerRoute := route.RouteGroup.Group("/partner")
	partnerRoute.Use(jwtMiddleware.JWTWithConfig)
	partnerRoute.GET("", handler.Browse)
	partnerRoute.GET("/profile", handler.BrowseProfilePartner)
	partnerRoute.GET("/:id", handler.ReadByPk)
	partnerRoute.PUT("/:id", handler.Edit)
	partnerRoute.PUT("/status/:id", handler.EditAccountStatus)
	partnerRoute.PUT("/webinar-status/:id", handler.EditWebinarStatus)
	partnerRoute.PUT("/verify/:id", handler.EditVerify)
	partnerRoute.POST("", handler.Add)
	partnerRoute.DELETE("/:id", handler.DeleteByPk)
}
