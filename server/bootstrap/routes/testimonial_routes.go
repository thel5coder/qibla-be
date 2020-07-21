package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type TestimonialRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route TestimonialRoutes) RegisterRoute(){
	handler := handlers.TestimonialHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	testimonialRoute := route.RouteGroup.Group("/testimony")
	testimonialRoute.Use(jwtMiddleware.JWTWithConfig)
	testimonialRoute.GET("",handler.Browse)
	testimonialRoute.GET("/:id",handler.Read)
	testimonialRoute.PUT("/:id",handler.Edit)
	testimonialRoute.POST("",handler.Add)
	testimonialRoute.DELETE("/:id",handler.Delete)
}
