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
