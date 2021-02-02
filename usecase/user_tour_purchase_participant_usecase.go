package usecase

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/enums"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/hashing"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/str"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type UserTourPurchaseParticipantUseCase struct {
	*UcContract
}

//store....
func (uc UserTourPurchaseParticipantUseCase) Store(inputs []requests.PassengerRequest, userTourPurchaseID string) (res []viewmodel.UserTourParticipantResVm, err error) {
	for _, input := range inputs {
		userParticipantID, err := uc.add(input, userTourPurchaseID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-userTourPurchaseParticipant-add")
			return res, err
		}
		res = append(res, viewmodel.UserTourParticipantResVm{
			ID:   userParticipantID,
			Name: input.Name,
		})
	}

	return res, nil
}

//add ...
func (uc UserTourPurchaseParticipantUseCase) add(input requests.PassengerRequest, userTourPurchaseID string) (res string, err error) {
	repository := actions.NewUserTourPurchaseParticipantRepository(uc.DB)
	now := time.Now().UTC()

	//set user
	roleUc := RoleUseCase{UcContract: uc.UcContract}
	role, err := roleUc.ReadBy("slug", "jamaah")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-role-readBySlug")
		return res, err
	}
	userUc := UserUseCase{UcContract: uc.UcContract}
	password, _ := hashing.HashAndSalt(str.RandomString(6))
	userID, err := userUc.Add(input.Name, input.Email, input.Email, input.Phone, role.ID, password, "", false, false)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "uc-user-add")
		return res, err
	}

	model := models.UserTourPurchaseParticipant{
		ID:                 "",
		UserTourPurchaseID: sql.NullString{String: userTourPurchaseID},
		UserID:             sql.NullString{String: userID},
		IdentityType:       sql.NullString{String: input.TypeOfIdentity},
		IdentityNumber:     sql.NullString{String: input.IdentityNumber},
		FullName:           sql.NullString{String: input.Name},
		Sex:                sql.NullString{String: input.Sex},
		BirthDate:          datetime.StrParseToTime(input.BirthDate, "2006-01-02"),
		BirthPlace:         sql.NullString{String: input.BirthPlace},
		PhoneNumber:        sql.NullString{String: input.Phone},
		CityID:             sql.NullString{String: input.CityID},
		Address:            sql.NullString{String: input.Address},
		Status:             sql.NullString{String: enums.StatusUserTourParticipant[0]},
		CreatedAt:          now,
		UpdatedAt:          now,
	}
	res, err = repository.Add(model, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query-userTourPurchaseParticipant-add")
		return res, err
	}

	return res, nil
}
