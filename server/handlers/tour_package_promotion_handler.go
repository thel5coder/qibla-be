package handlers

import (
	"github.com/labstack/echo"
	"qibla-backend/usecase"
	"strconv"
)

type TourPackagePromotionHandler struct {
	Handler
}

//browse
func (handler TourPackagePromotionHandler) Browse(ctx echo.Context) error {
	filters := make(map[string]interface{})

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	uc := usecase.TourPackagePromotionUseCase{UcContract: handler.UseCaseContract}
	res, pagination, err := uc.Browse(filters, "tpp.id", "asc", page, limit)

	return handler.SendResponse(ctx, res, pagination, err)
}

func (handler TourPackagePromotionHandler) ReadByPk(ctx echo.Context) error {
	ID := ctx.Param("id")

	uc := usecase.TourPackagePromotionUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.ReadBy("tpp.tour_package_id", ID, "=")

	return handler.SendResponse(ctx, res, nil, err)
}
