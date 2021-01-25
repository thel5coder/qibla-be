package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/server/requests"
	"time"
)

type UserTourPurchaseRoomUseCase struct {
	*UcContract
}

func (uc UserTourPurchaseRoomUseCase) Store(userTourPurchaseID string,inputs []requests.RoomRateRequest) (err error){
	tourPackagePriceUc := TourPackagePriceUseCase{UcContract:uc.UcContract}

	for _, input := range inputs {
		roomRate,err := tourPackagePriceUc.ReadBy("id",input.ID,"=")
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-tourPackagePrice-readByID")
			return err
		}
		roomRateRequest := requests.RoomRatePurchaseRequest{
			ID:       input.ID,
			Quantity: input.Quantity,
			Price:    roomRate.Price,
		}
		err = uc.Add(userTourPurchaseID,roomRateRequest)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"uc-userTourPurchaseRoom-add")
			return err
		}
	}

	return nil
}

func (uc UserTourPurchaseRoomUseCase) Add(userTourPurchaseID string, input requests.RoomRatePurchaseRequest) (err error) {
	repository := actions.NewUserTourPurchaseRoomRepository(uc.DB)
	now := time.Now().UTC()

	model := models.UserTourPurchaseRoom{
		UserTourPurchaseID: userTourPurchaseID,
		TourPackagePriceID: input.ID,
		Price:              input.Price,
		Quantity:           input.Quantity,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
	err = repository.Add(model,uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-userTourPurchaseRoom-add")
		return err
	}

	return nil
}
