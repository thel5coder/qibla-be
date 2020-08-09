package usecase

import (
	"qibla-backend/helpers/enums"
	"qibla-backend/usecase/viewmodel"
)

type EnumOptionsUseCase struct {
	*UcContract
}

func (uc EnumOptionsUseCase) GetMenuPermissionsEnums() (res []viewmodel.EnumVm){
	res = append(res ,viewmodel.EnumVm{
		Key:   enums.View,
		Value: enums.View,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.Add,
		Value: enums.Add,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.Edit,
		Value: enums.Edit,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.Delete,
		Value: enums.Delete,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.Approve,
		Value: enums.Approve,
	})

	return res
}

func (uc EnumOptionsUseCase) GetWebComprofCategoryEnums() (res []viewmodel.EnumVm){
	res = append(res ,viewmodel.EnumVm{
		Key:   enums.Gallery,
		Value: enums.Gallery,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.Faq,
		Value: enums.Faq,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.Testimoni,
		Value: enums.Testimoni,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.Article,
		Value: enums.Article,
	})

	return res
}

func (uc EnumOptionsUseCase) GetPromotionPackageEnum() (res []viewmodel.EnumVm){
	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PackagePromotionEnum1,
		Value: enums.PackagePromotionEnum1,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PackagePromotionEnum2,
		Value: enums.PackagePromotionEnum2,
	})

	return res
}

func (uc EnumOptionsUseCase) GetPlatformEnum() (res []viewmodel.EnumVm){
	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PlatformEnum1,
		Value: enums.PlatformEnum1,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PlatformEnum2,
		Value: enums.PlatformEnum2,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PlatformEnum3,
		Value: enums.PlatformEnum2,
	})

	return res
}

func (uc EnumOptionsUseCase) GetPositionPromotionEnum() (res []viewmodel.EnumVm){
	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PromotionPositionEnum1,
		Value: enums.PromotionPositionEnum1,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PromotionPositionEnum2,
		Value: enums.PromotionPositionEnum2,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PromotionPositionEnum3,
		Value: enums.PromotionPositionEnum3,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PromotionPositionEnum4,
		Value: enums.PromotionPositionEnum4,
	})

	res = append(res ,viewmodel.EnumVm{
		Key:   enums.PromotionPositionEnum5,
		Value: enums.PromotionPositionEnum5,
	})

	return res
}