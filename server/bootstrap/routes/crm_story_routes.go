package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type CrmStoryRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route CrmStoryRoutes) RegisterRoute(){
	handler := handlers.CrmStoryHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	crmStoryRoute := route.RouteGroup.Group("/crm-story")
	crmStoryRoute.Use(jwtMiddleware.JWTWithConfig)
	crmStoryRoute.GET("",handler.BrowseAll)
	crmStoryRoute.GET("/:id",handler.ReadByPk)
	crmStoryRoute.PUT("/:id",handler.Edit)
	crmStoryRoute.POST("",handler.Add)
	crmStoryRoute.DELETE("/:id",handler.Delete)
}
