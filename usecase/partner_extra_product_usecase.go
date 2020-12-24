package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
)

type PartnerExtraProductUseCase struct {
	*UcContract
}

func (uc PartnerExtraProductUseCase) BrowseByPartnerID(partnerID string) (res []viewmodel.PartnerExtraProductVm, err error) {
	repository := actions.NewPartnerExtraProductRepository(uc.DB)
	partnerProductSubscriptions, err := repository.BrowseByPartnerID(partnerID)
	if err != nil {
		return res, err
	}

	for _, partnerProductSubscription := range partnerProductSubscriptions {
		res = append(res, viewmodel.PartnerExtraProductVm{
			ID:        partnerProductSubscription.Product.ID,
			Name:      partnerProductSubscription.Product.Name,
			Price:     partnerProductSubscription.Product.Price,
			PriceUnit: partnerProductSubscription.Product.PriceUnit,
			Session:   partnerProductSubscription.Product.Session.String,
		})
	}

	return res, err
}

func (uc PartnerExtraProductUseCase) ReadBy(column, value string) (res viewmodel.PartnerExtraProductVm, err error) {
	repository := actions.NewPartnerExtraProductRepository(uc.DB)
	partnerProductSubscription, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = viewmodel.PartnerExtraProductVm{
		ID:        partnerProductSubscription.Product.ID,
		Name:      partnerProductSubscription.Product.Name,
		Price:     partnerProductSubscription.Product.Price,
		PriceUnit: partnerProductSubscription.Product.PriceUnit,
		Session:   partnerProductSubscription.Product.Session.String,
	}

	return res, err
}

func (uc PartnerExtraProductUseCase) Add(partnerID, productID string) (err error) {
	repository := actions.NewPartnerExtraProductRepository(uc.DB)
	err = repository.Add(partnerID, productID, uc.TX)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"query-partnerExtraProduct-add")
		return err
	}

	return nil
}

func (uc PartnerExtraProductUseCase) DeleteBy(column, value string) (err error) {
	repository := actions.NewPartnerExtraProductRepository(uc.DB)
	err = repository.DeleteBy(column, value, uc.TX)

	return err
}

func (uc PartnerExtraProductUseCase) Store(partnerID string, extraProducts []requests.ExtraProductRequest) (err error) {
	//count,err := uc.CountBy("partner_id",partnerID)
	//if err != nil {
	//	return err
	//}
	//
	//if count > 0 {
	//	err = uc.DeleteBy("partner_id",partnerID)
	//	if err != nil{
	//		return err
	//	}
	//}

	for _, product := range extraProducts {
		err = uc.Add(partnerID,product.ID)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel,err.Error(), functioncaller.PrintFuncName(),"uc-partnerExtraProduct-add")
			return err
		}
	}

	return nil
}

func (uc PartnerExtraProductUseCase) CountBy(column,value string) (res int,err error){
	repository := actions.NewPartnerExtraProductRepository(uc.DB)
	res,err = repository.CountBy(column,value)

	return res,err
}
