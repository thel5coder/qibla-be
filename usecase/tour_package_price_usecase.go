package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type TourPackagePriceUseCase struct {
	*UcContract
}

//read by
func (uc TourPackagePriceUseCase) ReadBy(column, value, operator string) (res viewmodel.TourPackagePriceVm, err error) {
	repository := actions.NewTourPackagePriceRepository(uc.DB)
	tourPackagePrice, err := repository.ReadBy(column, value, operator)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-tourPackagePrice-readBy")
		return res, err
	}
	res = uc.buildBody(tourPackagePrice)

	return res, nil
}

//build body
func (uc TourPackagePriceUseCase) buildBody(model models.TourPackagePrice) viewmodel.TourPackagePriceVm {
	return viewmodel.TourPackagePriceVm{
		ID:            model.ID,
		RoomType:      model.RoomType,
		RoomCapacity:  model.RoomCapacity,
		Price:         model.Price,
		PromoPrice:    model.PromoPrice,
		AirLineClass:  model.AirLineClass,
		TourPackageID: model.TourPackageID,
		RoomRateID:    model.RoomRateID,
		CreatedAt:     model.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     model.UpdatedAt.Format(time.RFC3339),
	}
}
