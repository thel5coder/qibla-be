package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/str"
	"qibla-backend/usecase/viewmodel"
	"strings"
)

type TourPackagePromotionUseCase struct {
	*UcContract
}

//browse
func (uc TourPackagePromotionUseCase) Browse(filters map[string]interface{}, order, sort string, page, limit int) (res []viewmodel.TourPackagePromotionVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewTourPackagePromotionRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	tourPackagePromotions, count, err := repository.Browse(filters, order, sort, limit, offset)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-tourPackagePromotion-browse")
		return res, pagination, err
	}

	for _, tourPackagePromotion := range tourPackagePromotions {
		res = append(res, uc.buildBody(tourPackagePromotion))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, nil
}

//read
func(uc TourPackagePromotionUseCase) ReadBy(column,value,operator string) (res viewmodel.TourPackagePromotionVm,err error){
	repository := actions.NewTourPackagePromotionRepository(uc.DB)

	tourPackagePromotion,err := repository.ReadBy(column,value,operator)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-tourPackagePromotion-read")
		return res, err
	}
	res = uc.buildBody(tourPackagePromotion)

	return res,nil
}

//build body hotel
func (uc TourPackagePromotionUseCase) buildBodyHotel(hotels string) (res []viewmodel.TourPackagePromotionHotelVm) {
	tourPackagePromotionHotels := str.Unique(strings.Split(hotels, ","))

	for _, hotel := range tourPackagePromotionHotels {
		hotelArr := strings.Split(hotel, ":")
		res = append(res, viewmodel.TourPackagePromotionHotelVm{
			City: hotelArr[3],
			Name: hotelArr[1],
		})
	}

	return res
}

//build body room rates
func (uc TourPackagePromotionUseCase) buildBodyRoomRates(roomRate string) (res []viewmodel.TourPackagePromotionRoomRateVm) {
	roomRates := str.Unique(strings.Split(roomRate, ","))

	for _, roomRate := range roomRates {
		roomRateArr := strings.Split(roomRate, ":")
		res = append(res, viewmodel.TourPackagePromotionRoomRateVm{
			ID:           roomRateArr[0],
			Room:         roomRateArr[1],
			AirlineClass: roomRateArr[5],
			Price:        int64(str.StringToInt(roomRateArr[3])),
			PromoPrice:   int64(str.StringToInt(roomRateArr[4])),
			RoomCapacity: str.StringToInt(roomRateArr[2]),
		})
	}

	return res
}

//build body
func (uc TourPackagePromotionUseCase) buildBody(model models.TourPackagePromotion) viewmodel.TourPackagePromotionVm {
	//build room rates
	var roomRatesVm []viewmodel.TourPackagePromotionRoomRateVm
	if model.RoomRates.String != "" {
		roomRatesVm = uc.buildBodyRoomRates(model.RoomRates.String)
	}

	//build hotels
	var hotels []viewmodel.TourPackagePromotionHotelVm
	if model.Hotels.String != "" {
		hotels = uc.buildBodyHotel(model.Hotels.String)
	}

	return viewmodel.TourPackagePromotionVm{
		ID: model.TourPackageID,
		TravelAgent: viewmodel.TourPackagePromotionTravelAgentVm{
			ID:     model.TravelAgentID,
			Name:   model.TravelAgentName,
			Branch: model.Branch,
			Phone:  model.Phone,
		},
		Name:               model.TourPackage.Name,
		ProgramDay:         model.TourPackage.ProgramDay,
		DepartureDate:      model.TourPackage.DepartureDate.Format("2006-01-02"),
		ReturnDate:         model.TourPackage.ReturnDate.Format("2006-01-02"),
		DepartureAirport:   model.TourPackage.DepartureAirport,
		DestinationAirport: model.TourPackage.DestinationAirport,
		Description:        model.TourPackage.Description.String,
		Image:              "",
		PackageType:        model.TourPackage.PackageType,
		Hotels:             hotels,
		RoomRates:          roomRatesVm,
	}
}
