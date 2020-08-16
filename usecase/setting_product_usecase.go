package usecase

import (
	"errors"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/enums"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type SettingProductUseCase struct {
	*UcContract
}

func (uc SettingProductUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.SettingProductVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	settingProductFeatureUc := SettingProductFeatureUseCase{UcContract: uc.UcContract}
	settingProductPeriodUc := SettingProductPeriodUseCase{UcContract: uc.UcContract}

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)
	settingProducts, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, settingProduct := range settingProducts {
		var settingProductFeatures []viewmodel.SettingProductFeatureVm
		var settingProductPeriods []viewmodel.SettingProductPeriodVm
		settingProductFeatures, _ = settingProductFeatureUc.BrowseBySettingProductID(settingProduct.ID)
		settingProductPeriods, _ = settingProductPeriodUc.BrowseBySettingProductID(settingProduct.ID)

		res = append(res, viewmodel.SettingProductVm{
			ID:                    settingProduct.ID,
			ProductID:             settingProduct.ProductID,
			ProductName:           settingProduct.ProductName,
			Price:                 settingProduct.Price,
			PriceUnit:             settingProduct.PriceUnit.String,
			MaintenancePrice:      settingProduct.MaintenancePrice.Int32,
			Discount:              settingProduct.Discount.Int32,
			DiscountType:          settingProduct.DiscountType.String,
			DiscountPeriodStart:   settingProduct.DiscountPeriodStart.String,
			DiscountPeriodEnd:     settingProduct.DiscountPeriodEnd.String,
			DiscountPeriod:        settingProduct.DiscountPeriodStart.String + ` - ` + settingProduct.DiscountPeriodEnd.String,
			Description:           settingProduct.Description,
			Sessions:              settingProduct.Sessions.String,
			CreatedAt:             settingProduct.CreatedAt,
			UpdatedAt:             settingProduct.UpdatedAt,
			DeletedAt:             settingProduct.DeletedAt.String,
			SettingProductFeature: settingProductFeatures,
			SettingProductPeriods: settingProductPeriods,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc SettingProductUseCase) BrowseAll() (res []viewmodel.SettingProductVm, err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	settingProductFeatureUc := SettingProductFeatureUseCase{UcContract: uc.UcContract}
	settingProductPeriodUc := SettingProductPeriodUseCase{UcContract: uc.UcContract}

	settingProducts, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, settingProduct := range settingProducts {
		var settingProductFeatures []viewmodel.SettingProductFeatureVm
		var settingProductPeriods []viewmodel.SettingProductPeriodVm
		settingProductFeatures, _ = settingProductFeatureUc.BrowseBySettingProductID(settingProduct.ID)
		for _, settingProductFeature := range settingProductFeatures {
			settingProductFeatures = append(settingProductFeatures, viewmodel.SettingProductFeatureVm{
				ID:          settingProductFeature.ID,
				FeatureName: settingProductFeature.FeatureName,
			})
		}

		settingProductPeriods, _ = settingProductPeriodUc.BrowseBySettingProductID(settingProduct.ID)
		for _, settingProductPeriod := range settingProductPeriods {
			settingProductPeriods = append(settingProductPeriods, viewmodel.SettingProductPeriodVm{
				ID:     settingProductPeriod.ID,
				Period: settingProductPeriod.Period,
			})
		}

		res = append(res, viewmodel.SettingProductVm{
			ID:                    settingProduct.ID,
			ProductID:             settingProduct.ProductID,
			ProductName:           settingProduct.ProductName,
			Price:                 settingProduct.Price,
			PriceUnit:             settingProduct.PriceUnit.String,
			MaintenancePrice:      settingProduct.MaintenancePrice.Int32,
			Discount:              settingProduct.Discount.Int32,
			DiscountType:          settingProduct.DiscountType.String,
			DiscountPeriodStart:   settingProduct.DiscountPeriodStart.String,
			DiscountPeriodEnd:     settingProduct.DiscountPeriodEnd.String,
			DiscountPeriod:        settingProduct.DiscountPeriodStart.String + ` - ` + settingProduct.DiscountPeriodEnd.String,
			Description:           settingProduct.Description,
			Sessions:              settingProduct.Sessions.String,
			CreatedAt:             settingProduct.CreatedAt,
			UpdatedAt:             settingProduct.UpdatedAt,
			DeletedAt:             settingProduct.DeletedAt.String,
			SettingProductFeature: settingProductFeatures,
			SettingProductPeriods: settingProductPeriods,
		})
	}

	return res, err
}

func (uc SettingProductUseCase) ReadBy(column, value string) (res viewmodel.SettingProductVm, err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	settingProductFeatureUc := SettingProductFeatureUseCase{UcContract: uc.UcContract}
	settingProductPeriodUc := SettingProductPeriodUseCase{UcContract: uc.UcContract}

	settingProduct, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	var settingProductFeatures []viewmodel.SettingProductFeatureVm
	var settingProductPeriods []viewmodel.SettingProductPeriodVm
	settingProductFeatures, _ = settingProductFeatureUc.BrowseBySettingProductID(settingProduct.ID)
	for _, settingProductFeature := range settingProductFeatures {
		settingProductFeatures = append(settingProductFeatures, viewmodel.SettingProductFeatureVm{
			ID:          settingProductFeature.ID,
			FeatureName: settingProductFeature.FeatureName,
		})
	}

	settingProductPeriods, _ = settingProductPeriodUc.BrowseBySettingProductID(settingProduct.ID)
	for _, settingProductPeriod := range settingProductPeriods {
		settingProductPeriods = append(settingProductPeriods, viewmodel.SettingProductPeriodVm{
			ID:     settingProductPeriod.ID,
			Period: settingProductPeriod.Period,
		})
	}

	res = viewmodel.SettingProductVm{
		ID:                    settingProduct.ID,
		ProductID:             settingProduct.ProductID,
		ProductName:           settingProduct.ProductName,
		Price:                 settingProduct.Price,
		PriceUnit:             settingProduct.PriceUnit.String,
		MaintenancePrice:      settingProduct.MaintenancePrice.Int32,
		Discount:              settingProduct.Discount.Int32,
		DiscountType:          settingProduct.DiscountType.String,
		DiscountPeriodStart:   settingProduct.DiscountPeriodStart.String,
		DiscountPeriodEnd:     settingProduct.DiscountPeriodEnd.String,
		DiscountPeriod:        settingProduct.DiscountPeriodStart.String + ` - ` + settingProduct.DiscountPeriodEnd.String,
		Description:           settingProduct.Description,
		Sessions:              settingProduct.Sessions.String,
		CreatedAt:             settingProduct.CreatedAt,
		UpdatedAt:             settingProduct.UpdatedAt,
		DeletedAt:             settingProduct.DeletedAt.String,
		SettingProductFeature: settingProductFeatures,
		SettingProductPeriods: settingProductPeriods,
	}

	return res, err
}

func (uc SettingProductUseCase) ReadByPk(ID string) (res viewmodel.SettingProductVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc SettingProductUseCase) Edit(ID string, input *requests.SettingProductRequest) (err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	settingProductFeatureUc := SettingProductFeatureUseCase{UcContract: uc.UcContract}
	settingProductPeriodUc := SettingProductPeriodUseCase{UcContract: uc.UcContract}
	masterProductUc := MasterProductUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)
	var priceUnit string

	count, err := uc.countBy(ID, "product_id", input.ProductID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	masterProduct, err := masterProductUc.ReadByPk(input.ProductID)
	if err != nil {
		return err
	}

	if masterProduct.SubscriptionType == enums.KeySubscriptionEnum1 {
		priceUnit = enums.KeyPriceUnitEnum3
	} else if masterProduct.SubscriptionType == enums.KeySubscriptionEnum3 {
		priceUnit = enums.KeyPriceUnitEnum4
	} else {
		priceUnit = input.PriceUnit
	}

	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	body := viewmodel.SettingProductVm{
		ID:                  ID,
		ProductID:           input.ProductID,
		Price:               input.Price,
		PriceUnit:           priceUnit,
		MaintenancePrice:    input.MaintenancePrice,
		Discount:            input.Discount,
		DiscountType:        input.DiscountType,
		DiscountPeriodStart: input.DiscountPeriodStart,
		DiscountPeriodEnd:   input.DiscountPeriodEnd,
		Description:         input.Description,
		Sessions:            input.Sessions,
		UpdatedAt:           now,
	}
	err = repository.Edit(body, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	err = settingProductPeriodUc.Store(ID, input.SettingPeriods, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	err = settingProductFeatureUc.Store(ID, input.SettingFeatures, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	uc.TX.Commit()

	return err
}

func (uc SettingProductUseCase) Add(input *requests.SettingProductRequest) (err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	settingProductFeatureUc := SettingProductFeatureUseCase{UcContract: uc.UcContract}
	settingProductPeriodUc := SettingProductPeriodUseCase{UcContract: uc.UcContract}
	masterProductUc := MasterProductUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)
	var priceUnit string

	count, err := uc.countBy("", "product_id", input.ProductID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	masterProduct, err := masterProductUc.ReadByPk(input.ProductID)
	if err != nil {
		return err
	}

	if masterProduct.SubscriptionType == enums.KeySubscriptionEnum1 {
		priceUnit = enums.KeyPriceUnitEnum3
	} else if masterProduct.SubscriptionType == enums.KeySubscriptionEnum3 {
		priceUnit = enums.KeyPriceUnitEnum4
	} else {
		priceUnit = input.PriceUnit
	}

	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	body := viewmodel.SettingProductVm{
		ProductID:           input.ProductID,
		Price:               input.Price,
		PriceUnit:           priceUnit,
		MaintenancePrice:    input.MaintenancePrice,
		Discount:            input.Discount,
		DiscountType:        input.DiscountType,
		DiscountPeriodStart: input.DiscountPeriodStart,
		DiscountPeriodEnd:   input.DiscountPeriodEnd,
		Description:         input.Description,
		Sessions:            input.Sessions,
		CreatedAt:           now,
		UpdatedAt:           now,
	}
	body.ID, err = repository.Add(body, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	err = settingProductPeriodUc.Store(body.ID, input.SettingPeriods, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	err = settingProductFeatureUc.Store(body.ID, input.SettingFeatures, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	uc.TX.Commit()

	return err
}

func (uc SettingProductUseCase) Delete(ID string) (err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	settingProductFeatureUc := SettingProductFeatureUseCase{UcContract: uc.UcContract}
	settingProductPeriodUc := SettingProductPeriodUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.countBy("", "id", ID)
	if err != nil {
		return err
	}

	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	if count > 0 {
		err = repository.Delete(ID, now, now, uc.TX)
		if err != nil {
			uc.TX.Rollback()

			return err
		}

		err = settingProductFeatureUc.DeleteBySettingProductID(ID, uc.TX)
		if err != nil {
			uc.TX.Rollback()

			return err
		}

		err = settingProductPeriodUc.DeleteBySettingProductID(ID, uc.TX)
		if err != nil {
			uc.TX.Rollback()

			return err
		}
	}
	uc.TX.Commit()

	return err
}

func (uc SettingProductUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}
