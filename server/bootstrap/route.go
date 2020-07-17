package bootstrap

import (
	"net/http"
	"qibla-backend/server/bootstrap/routes"
	api "qibla-backend/server/handlers"

	"github.com/labstack/echo"
)

func (boot *Bootstrap) RegisterRouters() {
	handlerType := api.Handler{
		E:               boot.E,
		Db:              boot.Db,
		UseCaseContract: &boot.UseCaseContract,
		Jwe:             boot.Jwe,
		Validate:        boot.Validator,
		Translator:      boot.Translator,
		JwtConfig:       boot.JwtConfig,
	}

	boot.E.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "Work")
	})

	 apiV1 := boot.E.Group("/api/v1")

	 //authentication route
	 authenticationRoute := routes.AuthenticationRoutes{
		 RouteGroup: apiV1,
		 Handler:    handlerType,
	 }
	 authenticationRoute.RegisterRoute()

	 //enum options route
	 enumOptionsRoute := routes.EnumOptionRoutes{
		 RouteGroup: apiV1,
		 Handler:    handlerType,
	 }
	 enumOptionsRoute.RegisterRoute()

	 //roleroute
	 roleRoute := routes.RoleRoutes{
		 RouteGroup: apiV1,
		 Handler:    handlerType,
	 }
	 roleRoute.RegisterRoute()

	 //userRoute
	 userRoute := routes.UserRoutes{
		 RouteGroup: apiV1,
		 Handler:    handlerType,
	 }
	 userRoute.RegisterRoute()

	 //term condition route
	 termConditionRoute := routes.TermConditionRoutes{
		 RouteGroup: apiV1,
		 Handler:    handlerType,
	 }
	 termConditionRoute.RegisterRoute()

	 menuRoute := routes.MenuRoutes{
		 RouteGroup: apiV1,
		 Handler:    handlerType,
	 }
	 menuRoute.RegisterRoute()

	 globalInfoCategoryMasterRoute := routes.GlobalInfoRoutes{
		 RouteGroup: apiV1,
		 Handler:    handlerType,
	 }
	 globalInfoCategoryMasterRoute.RegisterRoute()
}
