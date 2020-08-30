package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type CrmBoardRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route CrmBoardRoutes) RegisterRoute() {
	handler := handlers.CrmBoardHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	crmStoryRoute := route.RouteGroup.Group("/crm-board")
	crmStoryRoute.Use(jwtMiddleware.JWTWithConfig)
	crmStoryRoute.GET("",handler.Browse)
	//crmStoryRoute.GET("/:id",handler.ReadByPk)
	crmStoryRoute.PUT("/:id",handler.Edit)
	crmStoryRoute.PUT("/edit-story/:id",handler.EditBoardStory)
	crmStoryRoute.POST("",handler.Add)
	//crmStoryRoute.DELETE("/:id",handler.Delete)
}
