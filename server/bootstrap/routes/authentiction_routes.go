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
	authenticationRoute.POST("/register", authenticationHandler.RegisterJamaahByEmail)
	authenticationRoute.POST("/forgot", authenticationHandler.ForgotPassword)
	authenticationRoute.POST("/register-by-gmail", authenticationHandler.RegisterByOauth)
	authenticationRoute.POST("/oauth", authenticationHandler.RegisterByOauth)
	authenticationRoute.GET("/code", authenticationHandler.ActivationUserByCode)

	setPinRoute := authenticationRoute.Group("/set-pin")
	setPinRoute.Use(jwtMiddleware.JWTWithConfig)
	setPinRoute.POST("", authenticationHandler.SetPin)
	setFingerPrintRoute := authenticationRoute.Group("/set-fingerprint")
	setFingerPrintRoute.Use(jwtMiddleware.JWTWithConfig)
	setFingerPrintRoute.POST("", authenticationHandler.SetFingerPrintStatus)
}
