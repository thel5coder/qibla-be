package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type AuthenticationRoutes struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route AuthenticationRoutes) RegisterRoute() {
	authenticationHandler := handlers.AuthenticationHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract: route.Handler.UseCaseContract}

	authenticationRoute := route.RouteGroup.Group("/auth")
	authenticationRoute.POST("/login", authenticationHandler.Login)
	authenticationRoute.POST("/register", authenticationHandler.RegisterJaamaahByEmail)
	authenticationRoute.POST("/forgot", authenticationHandler.ForgotPassword)

	setPinRoute := authenticationRoute.Group("/set-pin")
	setPinRoute.Use(jwtMiddleware.JWTWithConfig)
	setPinRoute.POST("", authenticationHandler.SetPin)
}
