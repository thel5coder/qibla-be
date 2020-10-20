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

	//odoo route
	odooRoutes := routes.OdooRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	odooRoutes.RegisterRoute()

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

	//adminUserRoutes
	adminUserRoutes := routes.UserRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	adminUserRoutes.RegisterRoute()

	//term condition route
	termConditionRoute := routes.TermConditionRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	termConditionRoute.RegisterRoute()

	//menu route
	menuRoute := routes.MenuRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	menuRoute.RegisterRoute()

	//global info route
	globalInfoCategoryMasterRoute := routes.GlobalInfoRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	globalInfoCategoryMasterRoute.RegisterRoute()

	//web comprof route
	webComprofRoute := routes.WebComprofRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	webComprofRoute.RegisterRoute()

	//file route
	fileRoute := routes.FileRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	fileRoute.RegisterRoute()

	//gallery route
	galleryRoute := routes.GalleryRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	galleryRoute.RegisterRoute()

	//testimonial route
	testimonialRoute := routes.TestimonialRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	testimonialRoute.RegisterRoute()

	//faq route
	faqRoute := routes.FaqRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	faqRoute.RegisterRoute()

	//promotion package route
	promotionPackageRoutes := routes.PromotionPackageRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	promotionPackageRoutes.RegisterRoute()

	//promotion route
	promotionRoutes := routes.PromotionRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	promotionRoutes.RegisterRoute()

	//tour package
	tourPackageRoutes := routes.TourPackageRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	tourPackageRoutes.RegisterRoute()

	//master product
	masterProductRoutes := routes.MasterProductRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	masterProductRoutes.RegisterRoute()

	//setting product
	settingProductRoutes := routes.SettingProductRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	settingProductRoutes.RegisterRoute()

	//contact route
	contactRoutes := routes.ContactRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	contactRoutes.RegisterRoute()

	//calendar routes
	calendarRoutes := routes.CalendarRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	calendarRoutes.RegisterRoute()

	//complaint routes
	appComplaintRoutes := routes.AppComplaintRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	appComplaintRoutes.RegisterRoute()

	//pray routes
	prayRoutes := routes.PrayRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	prayRoutes.RegisterRoute()

	//video content routes
	videoContentRoutes := routes.VideoContentRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	videoContentRoutes.RegisterRoute()

	//master zakat routes
	masterZakatRoutes := routes.MasterZakatRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	masterZakatRoutes.RegisterRoute()

	//satisfaction category
	satisfactionCategoryRoutes := routes.SatisfactionCategoryRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	satisfactionCategoryRoutes.RegisterRoute()

	//crm story
	crmStoryRoutes := routes.CrmStoryRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	crmStoryRoutes.RegisterRoute()

	//crm board
	crmBoardRoutes := routes.CrmBoardRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	crmBoardRoutes.RegisterRoute()

	//partner routes
	partnerRoutes := routes.PartnerRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	partnerRoutes.RegisterRoute()

	faspayRoutes := routes.FaspayRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	faspayRoutes.RegisterRoute()

	transactionRoutes := routes.TransactionRoute{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	transactionRoutes.RegisterRoute()


	//mobile
	//video kajian routes
	videoKajianRoutes := routes.VideoKajianRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	videoKajianRoutes.RegisterRoute()

	//zakat
	zakatRoutes := routes.ZakatRoutes{
		RouteGroup: apiV1,
		Handler:    handlerType,
	}
	zakatRoutes.RegisterRoute()

}
