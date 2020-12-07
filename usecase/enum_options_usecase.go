package usecase

import (
	"qibla-backend/pkg/enums"
	"qibla-backend/pkg/enums/maritalstatusenum"
	"qibla-backend/pkg/enums/sexenum"
	"qibla-backend/usecase/viewmodel"
	"strconv"
)

type EnumOptionsUseCase struct {
	*UcContract
}

func (uc EnumOptionsUseCase) GetMenuPermissionsEnums() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.View,
		Value: enums.View,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.Add,
		Value: enums.Add,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.Edit,
		Value: enums.Edit,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.Delete,
		Value: enums.Delete,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.Approve,
		Value: enums.Approve,
	})

	return res
}

func (uc EnumOptionsUseCase) GetWebComprofCategoryEnums() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.Gallery,
		Value: enums.Gallery,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.Faq,
		Value: enums.Faq,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.Testimoni,
		Value: enums.Testimoni,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.Article,
		Value: enums.Article,
	})

	return res
}

func (uc EnumOptionsUseCase) GetPromotionPackageEnum() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPackagePromotionEnum1,
		Value: enums.ValuePackagePromotionEnum1,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPackagePromotionEnum2,
		Value: enums.ValuePackagePromotionEnum2,
	})

	return res
}

func (uc EnumOptionsUseCase) GetPlatformEnum() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPlatformEnum1,
		Value: enums.ValuePlatformEnum1,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPlatformEnum2,
		Value: enums.ValuePlatformEnum2,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPlatformEnum3,
		Value: enums.ValuePlatformEnum3,
	})

	return res
}

func (uc EnumOptionsUseCase) GetPositionPromotionEnum() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPromotionPositionEnum1,
		Value: enums.ValuePromotionPositionEnum1,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPromotionPositionEnum2,
		Value: enums.ValuePromotionPositionEnum2,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPromotionPositionEnum3,
		Value: enums.ValuePromotionPositionEnum3,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPromotionPositionEnum4,
		Value: enums.ValuePromotionPositionEnum4,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPromotionPositionEnum5,
		Value: enums.ValuePromotionPositionEnum5,
	})

	return res
}

func (uc EnumOptionsUseCase) GetSubscriptionEnum() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeySubscriptionEnum1,
		Value: enums.ValueSubscriptionEnum1,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeySubscriptionEnum2,
		Value: enums.ValueSubscriptionEnum2,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeySubscriptionEnum3,
		Value: enums.ValueSubscriptionEnum3,
	})

	return res
}

func (uc EnumOptionsUseCase) GetPriceUnitEnum() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPriceUnitEnum1,
		Value: enums.ValuePriceUnitEnum1,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyPriceUnitEnum2,
		Value: enums.ValuePriceUnitEnum2,
	})

	return res
}

func (uc EnumOptionsUseCase) GetDiscountTypeEnum() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyDiscountTypeEnum1,
		Value: enums.ValueDiscountTypeEnum1,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyDiscountTypeEnum2,
		Value: enums.ValueDiscountTypeEnum2,
	})

	return res
}

func (uc EnumOptionsUseCase) GetStatusComplaint() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyComplaintStatusEnum1,
		Value: enums.ValueComplaintStatusEnum1,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyComplaintStatusEnum2,
		Value: enums.ValueComplaintStatusEnum2,
	})

	return res
}

func (uc EnumOptionsUseCase) GetTypeZakat() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyTypeZakatEnum1,
		Value: enums.ValueTypeZakatEnum1,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   enums.KeyTypeZakatEnum2,
		Value: enums.ValueTypeZakatEnum2,
	})

	return res
}

func (uc EnumOptionsUseCase) GetRememberOptions() (res []viewmodel.EnumVm) {
	res = append(res, viewmodel.EnumVm{
		Key:   strconv.Itoa(enums.KeyRememberCalender1),
		Value: enums.ValueRememberCalender1,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   strconv.Itoa(enums.KeyRememberCalender2),
		Value: enums.ValueRememberCalender2,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   strconv.Itoa(enums.KeyRememberCalender3),
		Value: enums.ValueRememberCalender3,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   strconv.Itoa(enums.KeyRememberCalender4),
		Value: enums.ValueRememberCalender4,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   strconv.Itoa(enums.KeyRememberCalender5),
		Value: enums.ValueRememberCalender5,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   strconv.Itoa(enums.KeyRememberCalender6),
		Value: enums.ValueRememberCalender6,
	})

	res = append(res, viewmodel.EnumVm{
		Key:   strconv.Itoa(enums.KeyRememberCalender7),
		Value: enums.ValueRememberCalender7,
	})

	return res
}

func (uc EnumOptionsUseCase) GetSexEnum() []map[string]interface{} {
	var res []map[string]interface{}
	sexEnums := sexenum.GetEnums()

	for _, sexEnum := range sexEnums {
		res = append(res, map[string]interface{}{
			"key":  sexEnum["key"],
			"text": sexEnum["text"],
		})
	}

	return res
}

// GetMaritalStatusEnum ...
func (uc EnumOptionsUseCase) GetMaritalStatusEnum() []map[string]interface{} {
	return maritalstatusenum.GetEnums()
}
