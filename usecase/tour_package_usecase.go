package usecase

import (
	"fmt"
	"qibla-backend/helpers/datetime"
	"qibla-backend/helpers/interfacepkg"
	"qibla-backend/helpers/str"
	"qibla-backend/usecase/viewmodel"
	"strconv"
	"strings"
)

type TourPackageUseCase struct {
	*UcContract
}

func (uc TourPackageUseCase) BrowseAll() (res []viewmodel.TourPackageVm, err error) {
	odooUc := OdooUseCase{UcContract: uc.UcContract}
	var TravelPackageOdoos []viewmodel.TravelPackageOdooVm
	err = odooUc.Browse("travel.package", "", 0, 0, &TravelPackageOdoos)
	if err != nil {
		return res, err
	}

	for _, travelPackageOdoo := range TravelPackageOdoos {
		temp, err := uc.buildBody(travelPackageOdoo)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}
	fmt.Println(res)

	return res, nil
}

func (uc TourPackageUseCase) GetHotels(equipmentPackageIDs []interface{}) (res []viewmodel.TourPackageHotelVm, err error) {
	//get data from object equipment.package
	equipmentPackageIDStr := interfacepkg.InterfaceArrayToString(equipmentPackageIDs)
	equipmentPackageIDArr := strings.Split(equipmentPackageIDStr, ",")
	odooEquipmentPackages, err := uc.getEquipmentPackages(equipmentPackageIDArr)
	if err != nil {
		return res, err
	}

	//get data from object product.product
	productIDs := uc.getProductID(odooEquipmentPackages)
	odooProducts, err := uc.getProduct(productIDs)
	if err != nil {
		return res, err
	}

	//filter the object hotel ok is true append data to struct TourPackageHotelVm
	for _, odooProduct := range odooProducts {
		if odooProduct.HotelOK.Get() {
			res = append(res, viewmodel.TourPackageHotelVm{
				Name:           odooProduct.DisplayName,
				FacilityRating: int64(str.StringToInt(odooProduct.Rating)),
				Location:       "",
			})
		}
	}

	return res, nil
}

func (uc TourPackageUseCase) getEquipmentPackages(IDs []string) (res []viewmodel.OdooEquipmentPackageVm, err error) {
	odooUc := OdooUseCase{UcContract: uc.UcContract}
	var temp []viewmodel.OdooEquipmentPackageVm

	for _, ID := range IDs {
		err = odooUc.Read("equipment.package", int64(str.StringToInt(ID)), &temp)
		if err != nil {
			return res, err
		}
		res = append(res, temp[0])
	}

	return res, nil
}

func (uc TourPackageUseCase) getProductID(odoEquipmentPackages []viewmodel.OdooEquipmentPackageVm) (res []int64) {
	for _, odoEquipmentPackage := range odoEquipmentPackages {
		productStr := interfacepkg.InterfaceArrayToString(odoEquipmentPackage.ProductID)
		productStrArr := strings.Split(productStr, ",")
		res = append(res, int64(str.StringToInt(productStrArr[0])))
	}

	return res
}

func (uc TourPackageUseCase) getProduct(productIDs []int64) (res []viewmodel.OdooProductVm, err error) {
	odooUc := OdooUseCase{UcContract: uc.UcContract}
	var temp []viewmodel.OdooProductVm
	for _, productID := range productIDs {
		err = odooUc.Read("product.product", productID, &temp)
		if err != nil {
			return res, err
		}
		res = append(res, temp[0])
	}

	return res, nil
}

func (uc TourPackageUseCase) getTourPackagePriceRates(roomRateIDs []interface{}) (res []viewmodel.TourPackagePriceVm, err error) {
	var temps []viewmodel.TourPackagePriceVm

	//get data from object room.rates
	roomRateIDstr := interfacepkg.InterfaceArrayToString(roomRateIDs)
	roomRateIDarr := strings.Split(roomRateIDstr, ",")
	roomRates, err := uc.getRoomRates(roomRateIDarr)
	if err != nil {
		return res, err
	}

	for _, roomRate := range roomRates {
		//get data from object room.room
		roomID := uc.getRoomID(roomRate)
		rooms, err := uc.getRooms(roomID)
		if err != nil {
			return res, err
		}

		temps = append(temps,viewmodel.TourPackagePriceVm{
			RoomType:     rooms[0].DisplayName,
			RoomCapacity: rooms[0].NumberOfPerson,
			Price:        int64(roomRate.PriceUnit),
			PricePromo:   int64(roomRate.PricePromo),
			IsDefault:    false,
		})
	}


	return res, nil
}

func (uc TourPackageUseCase) getRoomRates(roomRateIDs []string) (res []viewmodel.OdooRoomRateVm, err error) {
	odooUc := OdooUseCase{UcContract: uc.UcContract}
	var temp []viewmodel.OdooRoomRateVm

	for _, roomRateID := range roomRateIDs {
		err = odooUc.Read("room.rate", int64(str.StringToInt(roomRateID)), &temp)
		if err != nil {
			return res, err
		}
		res = append(res, temp[0])
	}

	return res, err
}

func (uc TourPackageUseCase) getRoomID(odooRoomRate viewmodel.OdooRoomRateVm) (res int64) {
	odooRoomIDStr := interfacepkg.InterfaceArrayToString(odooRoomRate.RoomID)
	odooRoomIDArr := strings.Split(odooRoomIDStr, ",")
	res = int64(str.StringToInt(odooRoomIDArr[0]))

	return res
}

func (uc TourPackageUseCase) getRooms(roomID int64) (res []viewmodel.OdooRoomVm, err error) {
	odooUc := OdooUseCase{UcContract: uc.UcContract}

	err = odooUc.Read("room.room", roomID, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

//extract data from equipmentPackageId interface object
func (uc TourPackageUseCase) getPackageType(equipmentPackageID []interface{}) (res string) {
	equipmentPackageStr := interfacepkg.InterfaceArrayToString(equipmentPackageID)
	equipmentPackageArr := strings.Split(equipmentPackageStr, ",")

	return equipmentPackageArr[1]
}

func (uc TourPackageUseCase) countProgramDays(departurePackage, returnPackage string) (res string) {
	departureDate := datetime.StrParseToTime(departurePackage, "2006-01-02")
	returnDate := datetime.StrParseToTime(returnPackage, "2006-01-02")

	days := returnDate.Sub(departureDate).Hours() / 24
	res = strconv.Itoa(int(days)) + ` Hari`

	return res
}

func (uc TourPackageUseCase) buildBody(model viewmodel.TravelPackageOdooVm) (res viewmodel.TourPackageVm, err error) {
	tourPackageHotels, err := uc.GetHotels(model.EquipmentPackageIDs)
	if err != nil {
		return res, err
	}

	_, err = uc.getTourPackagePriceRates(model.RoomRateIDS)
	if err != nil {
		return res, err
	}

	res = viewmodel.TourPackageVm{
		OdooID:        model.ID,
		Name:          model.DisplayName,
		Package:       uc.getPackageType(model.PackageEquipmentID),
		ProgramDays:   uc.countProgramDays(model.ArrivalDate, model.ReturnDate),
		DepartureDate: model.ArrivalDate,
		ReturnDate:    model.ReturnDate,
		Hotels:        tourPackageHotels,
		Airlines:      nil,
		Prices:        nil,
	}

	return res, err
}
