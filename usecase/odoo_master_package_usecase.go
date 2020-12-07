package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/odoohelper"
	"qibla-backend/pkg/str"
	"qibla-backend/usecase/viewmodel"
	"strings"
)

type OdooMasterPackageUseCase struct {
	*UcContract
}

func (uc OdooMasterPackageUseCase) BrowseAll(partnerID string) (res []viewmodel.OdooMasterPackageVm, err error) {
	odooDb := odoohelper.Connection{
		Host:     "staging.qibla.co.id",
		DbName:   "himupd",
		User:     "czmiusbdajga",
		Password: "SwaddlingChoosingMatador",
		Port:     "5433",
		SslMode: "disable",
	}
	conn, err := odooDb.DbConnect()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "odoo-connection-initialization")
		return res, err
	}

	repository := actions.NewOdooMasterPackageRepository(conn)
	odooMasterPackages, err := repository.BrowseAll()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-browseAll-odoo-master-package")
		return res, err
		return res, err
	}
	defer conn.Close()

	for _, odooMasterPackage := range odooMasterPackages {
		res = append(res, uc.buildBody(odooMasterPackage))
	}

	return res, nil
}

func (uc OdooMasterPackageUseCase) ReadBy(column,value,operator string) (res viewmodel.OdooMasterPackageVm,err error){
	odooDb := odoohelper.Connection{
		Host:     "staging.qibla.co.id",
		DbName:   "himupd",
		User:     "czmiusbdajga",
		Password: "SwaddlingChoosingMatador",
		Port:     "5433",
		SslMode: "disable",
	}
	conn, err := odooDb.DbConnect()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "odoo-connection-initialization")
		return res, err
	}

	repository := actions.NewOdooMasterPackageRepository(conn)
	odooMasterPackage,err := repository.ReadBy(column,value,operator)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-readBy-odoo-master-package")
		return res, err
	}
	res = uc.buildBody(odooMasterPackage)

	return res,nil
}

func (uc OdooMasterPackageUseCase) hotelBuildBody(hotels string) (res []viewmodel.OdooMasterPackagePackageHotelVm) {
	odooMasterPackageHotels := strings.Split(hotels, ",")
	for _, odooMasterPackageHotel := range odooMasterPackageHotels {
		hotelArr := strings.Split(odooMasterPackageHotel, ":")
		res = append(res, viewmodel.OdooMasterPackagePackageHotelVm{
			OdooHotelID:       int32(str.StringToInt(hotelArr[0])),
			ProductTemplateID: int32(str.StringToInt(hotelArr[1])),
			Name:              hotelArr[2],
			FacilityRating:    int32(str.StringToInt(hotelArr[3])),
			Location:          "",
		})
	}

	return res
}

func (uc OdooMasterPackageUseCase) transportationBuildBody(transportations string) (res []viewmodel.OdooMasterPackageTransportationVm) {
	odooMasterPackageTransportations := strings.Split(transportations, ",")
	for _, odooMasterPackageTransportation := range odooMasterPackageTransportations {
		transportationArr := strings.Split(odooMasterPackageTransportation, ":")
		res = append(res, viewmodel.OdooMasterPackageTransportationVm{
			OdooTransportationID:  int32(str.StringToInt(transportationArr[0])),
			OdooProductTemplateID: int32(str.StringToInt(transportationArr[1])),
			Name:                  transportationArr[2],
		})
	}

	return res
}

func (uc OdooMasterPackageUseCase) mealBuildBody(meals string) (res []viewmodel.OdooMasterPackageMealsVm) {
	odooMasterPackageMeals := strings.Split(meals, ",")
	for _, odooMasterPackageMeal := range odooMasterPackageMeals {
		mealArr := strings.Split(odooMasterPackageMeal, ":")
		res = append(res, viewmodel.OdooMasterPackageMealsVm{
			OdooMealID:        int32(str.StringToInt(mealArr[0])),
			ProductTemplateID: int32(str.StringToInt(mealArr[1])),
			Name:              mealArr[2],
		})
	}

	return res
}

func (uc OdooMasterPackageUseCase) airlineBuildBody(airlines string) (res []viewmodel.OdooMasterPackageAirlineVm) {
	odooMasterPackageAirlines := strings.Split(airlines, ",")
	for _, odooMasterPackageAirline := range odooMasterPackageAirlines {
		airlineArr := strings.Split(odooMasterPackageAirline, ":")
		res = append(res, viewmodel.OdooMasterPackageAirlineVm{
			OdooTourAirlineID: int32(str.StringToInt(airlineArr[0])),
			Name:              airlineArr[1],
		})
	}

	return res
}

func (uc OdooMasterPackageUseCase) roomRateBuildBody(roomRates string) (res []viewmodel.OdooMasterPackageRoomRate) {
	odooMasterPackageRoomRates := strings.Split(roomRates, ",")
	for _, odooMasterPackageRoomRate := range odooMasterPackageRoomRates {
		roomRateArr := strings.Split(odooMasterPackageRoomRate, ":")
		res = append(res, viewmodel.OdooMasterPackageRoomRate{
			RoomRateID:   int32(str.StringToInt(roomRateArr[0])),
			RoomType:     roomRateArr[1],
			Price:        float32(str.StringToInt(roomRateArr[2])),
			PricePromo:   float32(str.StringToInt(roomRateArr[3])),
			RoomCapacity: int32(str.StringToInt(roomRateArr[4])),
			AirlineClass: roomRateArr[5],
		})
	}

	return res
}

func (uc OdooMasterPackageUseCase) buildBody(model models.OdooMasterPackage) viewmodel.OdooMasterPackageVm {
	odooMasterPackageHotelVm := uc.hotelBuildBody(model.Hotels)
	odooMasterPackageTransportationVm := uc.transportationBuildBody(model.Transportations)
	odooMasterPackageMealVm := uc.mealBuildBody(model.Meals)
	odooMasterPackageAirlineVm := uc.airlineBuildBody(model.Airlines)
	odooMasterPackageRoomRateVm := uc.roomRateBuildBody(model.RoomRates)

	departureDate := datetime.StrParseToTime(model.DepartureDate, "2006-01-02")
	returnDate := datetime.StrParseToTime(model.ReturnDate, "2006-01-02")
	days := returnDate.Sub(departureDate).Hours() / 24

	return viewmodel.OdooMasterPackageVm{
		OdooID:          int32(str.StringToInt(model.ID)),
		Name:            model.Name,
		PackageType:     model.EquipmentPackageName,
		PackageTypeID:   model.EquipmentPackageID,
		ProgramDays:     int32(days),
		DepartureDate:   model.DepartureDate,
		ReturnDate:      model.ReturnDate,
		Quota:           model.Quota,
		Notes:           model.Notes.String,
		WebDescription:  model.WebsiteDescription.String,
		Hotels:          odooMasterPackageHotelVm,
		Airlines:        odooMasterPackageAirlineVm,
		Meals:           odooMasterPackageMealVm,
		Transportations: odooMasterPackageTransportationVm,
		RoomRates:       odooMasterPackageRoomRateVm,
	}
}
