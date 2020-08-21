package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type AppComplaintRoutes struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route AppComplaintRoutes) RegisterRoute(){
	handler := handlers.AppComplaintHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	complaintRoutes := route.RouteGroup.Group("/complaint")
	complaintRoutes.Use(jwtMiddleware.JWTWithConfig)
	complaintRoutes.GET("", handler.Browse)
	complaintRoutes.GET("/ticket-number",handler.GetTicketNumber)
	complaintRoutes.GET("/:id", handler.Read)
	complaintRoutes.PUT("/:id", handler.Edit)
	complaintRoutes.POST("", handler.Add)
	complaintRoutes.DELETE("/:id", handler.Delete)
}
