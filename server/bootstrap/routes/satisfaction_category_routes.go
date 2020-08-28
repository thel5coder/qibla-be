package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type SatisfactionCategoryRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route SatisfactionCategoryRoutes) RegisterRoute(){
	handler := handlers.SatisfactionCategoryHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	satisfactionCategoryRoute := route.RouteGroup.Group("/satisfaction-category")
	satisfactionCategoryRoute.Use(jwtMiddleware.JWTWithConfig)
	satisfactionCategoryRoute.GET("/all/parent",handler.BrowseAllParent)
	satisfactionCategoryRoute.GET("/all/tree",handler.BrowseAllTree)
	satisfactionCategoryRoute.GET("/:id",handler.ReadByPk)
	satisfactionCategoryRoute.POST("",handler.Store)
	satisfactionCategoryRoute.DELETE("/:id",handler.DeleteByPk)
}
