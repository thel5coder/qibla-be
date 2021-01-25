package usecase

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/server/requests"
	"time"
)

// UserTourPurchaseUseCase ...
type UserTourPurchaseUseCase struct {
	*UcContract
}

// Add ...
func (uc UserTourPurchaseUseCase) Add(input *requests.CreatePurchaseRequest) (res string, err error) {
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
	res, err = repository.Add(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-userTourPurchase-add")
		return res, err
	}

	userTourPurchaseRoomUc := UserTourPurchaseRoomUseCase{UcContract:uc.UcContract}
	err = userTourPurchaseRoomUc.Store(res,input.RoomRates)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-userTourPurchaseRoom-store")
		return res, err
	}

	return res, nil
}

// Update ...
func (uc UserTourPurchaseUseCase) Update(input *requests.CreatePassengerRequest) (res map[string]interface{}, err error) {
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
		return
	}

	res = map[string]interface{}{
		"package_purchase_id": input.PackagePurchaseID,
		"passengers":          []string{""},
	}

	return res, nil
}
