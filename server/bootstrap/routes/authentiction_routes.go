package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
)

type AuthenticationRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route AuthenticationRoutes) RegisterRoute(){
	authenticationHandler := handlers.AuthenticationHandler{Handler:route.Handler}

	authenticationRoute:=route.RouteGroup.Group("/auth")
	authenticationRoute.POST("/login",authenticationHandler.Login)
}
