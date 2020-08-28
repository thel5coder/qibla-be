package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type PrayRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route PrayRoutes) RegisterRoute(){
	handler := handlers.PrayHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	prayRoute := route.RouteGroup.Group("/pray")
	prayRoute.Use(jwtMiddleware.JWTWithConfig)
	prayRoute.GET("",handler.Browse)
	prayRoute.GET("/:id",handler.Read)
	prayRoute.PUT("/:id",handler.Edit)
	prayRoute.POST("",handler.Add)
	prayRoute.DELETE("/:id",handler.Delete)

}
