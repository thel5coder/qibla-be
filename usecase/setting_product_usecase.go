package usecase

import (
	"errors"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/enums"
	"qibla-backend/pkg/messages"
	"qibla-backend/pkg/str"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

// SettingProductUseCase ...
type SettingProductUseCase struct {
	*UcContract
}

// Browse ...
func (uc SettingProductUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.SettingProductVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewSettingProductRepository(uc.DB)

	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)
	settingProducts, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, settingProduct := range settingProducts {
		res = append(res, uc.buildBody(&settingProduct))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

// BrowseAll ...
func (uc SettingProductUseCase) BrowseAll() (res []viewmodel.SettingProductVm, err error) {
	repository := actions.NewSettingProductRepository(uc.DB)

	settingProducts, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, settingProduct := range settingProducts {
		res = append(res, uc.buildBody(&settingProduct))
	}

	return res, err
}

// BrowseBy ...
func (uc SettingProductUseCase) BrowseBy(column, value, operator string) (res []viewmodel.SettingProductVm, err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	settingProducts, err := repository.BrowseBy(column, value, operator)

	for _, settingProduct := range settingProducts {
		res = append(res, uc.buildBody(&settingProduct))
	}

	return res, err
}

// ReadBy ...
func (uc SettingProductUseCase) ReadBy(column, value string) (res viewmodel.SettingProductVm, err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	settingProduct, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res = uc.buildBody(&settingProduct)

	return res, err
}

// ReadByPk ...
func (uc SettingProductUseCase) ReadByPk(ID string) (res viewmodel.SettingProductVm, err error) {
	res, err = uc.ReadBy("sp.id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

// Edit ...
func (uc SettingProductUseCase) Edit(ID string, input *requests.SettingProductRequest) (err error) {
	count, err := uc.countBy(ID, "product_id", input.ProductID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	masterProductUc := MasterProductUseCase{UcContract: uc.UcContract}
	masterProduct, err := masterProductUc.ReadByPk(input.ProductID)
	if err != nil {
		return err
	}

	var priceUnit string
	if masterProduct.SubscriptionType == enums.KeySubscriptionEnum1 {
		priceUnit = enums.KeyPriceUnitEnum3
	} else if masterProduct.SubscriptionType == enums.KeySubscriptionEnum3 {
		priceUnit = enums.KeyPriceUnitEnum4
	} else {
		priceUnit = input.PriceUnit
	}

	now := time.Now().UTC().Format(time.RFC3339)
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
	repository := actions.NewSettingProductRepository(uc.DB)
	err = repository.Edit(body, uc.TX)
	if err != nil {
		return err
	}

	settingProductPeriodUc := SettingProductPeriodUseCase{UcContract: uc.UcContract}
	err = settingProductPeriodUc.Store(ID, input.SettingPeriods, uc.TX)
	if err != nil {
		return err
	}

	settingProductFeatureUc := SettingProductFeatureUseCase{UcContract: uc.UcContract}
	err = settingProductFeatureUc.Store(ID, input.SettingFeatures, uc.TX)
	if err != nil {
		return err
	}

	return nil
}

// Add ...
func (uc SettingProductUseCase) Add(input *requests.SettingProductRequest) (err error) {
	count, err := uc.countBy("", "product_id", input.ProductID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(messages.DataAlreadyExist)
	}

	masterProductUc := MasterProductUseCase{UcContract: uc.UcContract}
	masterProduct, err := masterProductUc.ReadByPk(input.ProductID)
	if err != nil {
		return err
	}

	var priceUnit string
	if masterProduct.SubscriptionType == enums.KeySubscriptionEnum1 {
		priceUnit = enums.KeyPriceUnitEnum3
	} else if masterProduct.SubscriptionType == enums.KeySubscriptionEnum3 {
		priceUnit = enums.KeyPriceUnitEnum4
	} else {
		priceUnit = input.PriceUnit
	}

	now := time.Now().UTC().Format(time.RFC3339)
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
	repository := actions.NewSettingProductRepository(uc.DB)
	body.ID, err = repository.Add(body, uc.TX)
	if err != nil {
		return err
	}

	settingProductPeriodUc := SettingProductPeriodUseCase{UcContract: uc.UcContract}
	err = settingProductPeriodUc.Store(body.ID, input.SettingPeriods, uc.TX)
	if err != nil {
		return err
	}

	settingProductFeatureUc := SettingProductFeatureUseCase{UcContract: uc.UcContract}
	err = settingProductFeatureUc.Store(body.ID, input.SettingFeatures, uc.TX)
	if err != nil {
		return err
	}

	return err
}

// Delete ...
func (uc SettingProductUseCase) Delete(ID string) (err error) {
	count, _ := uc.countBy("", "id", ID)
	if count == 0 {
		return errors.New(messages.DataNotFound)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	repository := actions.NewSettingProductRepository(uc.DB)
	err = repository.Delete(ID, now, now, uc.TX)
	if err != nil {
		return err
	}

	settingProductFeatureUc := SettingProductFeatureUseCase{UcContract: uc.UcContract}
	err = settingProductFeatureUc.DeleteBySettingProductID(ID, uc.TX)
	if err != nil {
		return err
	}

	settingProductPeriodUc := SettingProductPeriodUseCase{UcContract: uc.UcContract}
	err = settingProductPeriodUc.DeleteBySettingProductID(ID, uc.TX)

	return err
}

func (uc SettingProductUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewSettingProductRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc SettingProductUseCase) buildSettingProductFeature(data string) (res []viewmodel.SettingProductFeatureVm) {
	dataArr := str.Unique(strings.Split(data, "|"))
	for _, d := range dataArr {
		dSplit := strings.Split(d, "#")
		if len(dSplit) != 2 {
			continue
		}

		res = append(res, viewmodel.SettingProductFeatureVm{
			ID:          dSplit[0],
			FeatureName: dSplit[1],
		})
	}
	return res
}

func (uc SettingProductUseCase) buildSettingProductPeriod(data string) (res []viewmodel.SettingProductPeriodVm) {
	dataArr := str.Unique(strings.Split(data, "|"))
	for _, d := range dataArr {
		dSplit := strings.Split(d, "#")
		if len(dSplit) != 2 {
			continue
		}

		res = append(res, viewmodel.SettingProductPeriodVm{
			ID:     dSplit[0],
			Period: str.StringToInt(dSplit[1]),
		})
	}
	return res
}

func (uc SettingProductUseCase) buildBody(data *models.SettingProduct) (res viewmodel.SettingProductVm) {
	settingProductFeatures := uc.buildSettingProductFeature(data.Features.String)
	settingProductPeriods := uc.buildSettingProductPeriod(data.Periods.String)

	return viewmodel.SettingProductVm{
		ID:                    data.ID,
		ProductID:             data.ProductID,
		ProductName:           data.ProductName,
		ProductType:           data.ProductType,
		Price:                 data.Price,
		PriceUnit:             data.PriceUnit.String,
		MaintenancePrice:      data.MaintenancePrice.Int32,
		Discount:              data.Discount.Int32,
		DiscountType:          data.DiscountType.String,
		DiscountPeriodStart:   data.DiscountPeriodStart.String,
		DiscountPeriodEnd:     data.DiscountPeriodEnd.String,
		DiscountPeriod:        data.DiscountPeriodStart.String + ` - ` + data.DiscountPeriodEnd.String,
		Description:           data.Description,
		Sessions:              data.Sessions.String,
		CreatedAt:             data.CreatedAt,
		UpdatedAt:             data.UpdatedAt,
		DeletedAt:             data.DeletedAt.String,
		SettingProductFeature: settingProductFeatures,
		SettingProductPeriods: settingProductPeriods,
	}
}
