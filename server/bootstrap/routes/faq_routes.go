package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type FaqRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route FaqRoutes) RegisterRoute(){
	handler := handlers.FaqHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	faqRoute := route.RouteGroup.Group("/faq")
	faqRoute.Use(jwtMiddleware.JWTWithConfig)
	faqRoute.GET("",handler.Browse)
	faqRoute.GET("/:id",handler.Read)
	faqRoute.PUT("/:id",handler.Edit)
	faqRoute.POST("",handler.Add)
	faqRoute.DELETE("",handler.Delete)
}
