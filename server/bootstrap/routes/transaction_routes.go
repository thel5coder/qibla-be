package routes

import (
	"github.com/labstack/echo"
	"qibla-backend/server/handlers"
	"qibla-backend/server/middleware"
)

type TransactionRoute struct {
	RouteGroup *echo.Group
	Handler handlers.Handler
}

func (route TransactionRoute) RegisterRoute(){
	handler := handlers.TransactionHandler{Handler:route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	transactionRoute := route.RouteGroup.Group("/transaction")
	transactionRoute.Use(jwtMiddleware.JWTWithConfig)
	transactionRoute.GET("/invoice",handler.GetInvoiceCount)
	transactionRoute.GET("/:trxId",handler.ReadByTrxID)
}