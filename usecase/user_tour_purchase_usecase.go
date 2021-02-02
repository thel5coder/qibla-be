package usecase

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

// UserTourPurchaseUseCase ...
type UserTourPurchaseUseCase struct {
	*UcContract
}

// Add ...
func (uc UserTourPurchaseUseCase) Add(input *requests.CreatePurchaseRequest) (res viewmodel.UserTourPurchaseRespVm, err error) {
	repository := actions.NewUserTourPurchaseRepository(uc.DB)
	now := time.Now().UTC()

	tourPackagePriceUc := TourPackagePriceUseCase{UcContract: uc.UcContract}
	tourPackagePrice, err := tourPackagePriceUc.ReadBy("id", input.RoomRates[0].ID, "=")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-tourPackagePrice-readByID")
		return res, err
	}

	model := models.UserTourPurchase{
		TourPackageID: tourPackagePrice.TourPackageID,
		UserID:        uc.UserID,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	res.UserTourPurchaseID, err = repository.Add(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-userTourPurchase-add")
		return res, err
	}

	userTourPurchaseRoomUc := UserTourPurchaseRoomUseCase{UcContract: uc.UcContract}
	err = userTourPurchaseRoomUc.Store(res.UserTourPurchaseID, input.RoomRates)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-userTourPurchaseRoom-store")
		return res, err
	}

	var roomRatesRespVm []viewmodel.RoomRatesRespVM
	for _,roomRate := range input.RoomRates {
		roomRateRes,err := tourPackagePriceUc.ReadBy("id",roomRate.ID,"=")
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-tourPackagePrice-readByID")
			return res,err
		}
		roomRatesRespVm = append(roomRatesRespVm,viewmodel.RoomRatesRespVM{
			ID:       roomRate.ID,
			Name:     roomRateRes.RoomType,
			Quantity: roomRate.Quantity,
		})
	}
	res.RoomRates = roomRatesRespVm

	return res, nil
}

// CreatePassenger ...
func (uc UserTourPurchaseUseCase) CreatePassenger(input *requests.CreatePassengerRequest) (res []viewmodel.UserTourParticipantResVm, err error) {
	repository := actions.NewUserTourPurchaseRepository(uc.DB)
	now := time.Now().UTC()

	registrantRequest := input.Registrant
	model := models.UserTourPurchase{
		CustomerIdentityType: sql.NullString{String: registrantRequest.TypeOfIdentity},
		IdentityNumber:       sql.NullString{String: registrantRequest.IdentityNumber},
		FullName:             sql.NullString{String: registrantRequest.Name},
		Sex:                  sql.NullString{String: registrantRequest.Sex},
		BirthDate:            sql.NullTime{Time: datetime.StrParseToTime(registrantRequest.BirthDate, "2006-01-02")},
		BirthPlace:           sql.NullString{String: registrantRequest.BirthPlace},
		PhoneNumber:          sql.NullString{String: registrantRequest.Phone},
		CityID:               sql.NullString{String: registrantRequest.CityID},
		MaritalStatus:        sql.NullString{String: registrantRequest.MaritalStatus},
		CustomerAddress:      sql.NullString{String: registrantRequest.Address},
		UpdatedAt:            now,
	}
	err = repository.Edit(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-userTourPurchase-edit")
		return res,err
	}

	for i := 0; i < len(input.Passengers); i++ {
		if input.Passengers[i].IsRegistrant {
			input.Passengers[i] = requests.PassengerRequest{
				IsRegistrant:   input.Passengers[i].IsRegistrant,
				Email:          input.Passengers[i].Email,
				TypeOfIdentity: input.Registrant.TypeOfIdentity,
				IdentityNumber: input.Registrant.IdentityNumber,
				Name:           input.Registrant.Name,
				Sex:            input.Registrant.Sex,
				BirthDate:      input.Registrant.BirthDate,
				BirthPlace:     input.Registrant.BirthPlace,
				Phone:          input.Registrant.Phone,
				Address:        input.Registrant.Address,
				CityID:         input.Registrant.CityID,
				MaritalStatus:  input.Registrant.MaritalStatus,
			}
		}
	}

	userTourPurchaseParticipantUc := UserTourPurchaseParticipantUseCase{UcContract:uc.UcContract}
	res,err = userTourPurchaseParticipantUc.Store(input.Passengers,input.PackagePurchaseID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-userTourPurchaseParticipant-store")
		return res,err
	}

	return res, nil
}

//create document
func(uc UserTourPurchaseUseCase) CreateDocument() {

}
