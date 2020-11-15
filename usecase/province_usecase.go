package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type ProvinceUseCase struct {
	*UcContract
}

func (uc ProvinceUseCase) BrowseAll() (res []viewmodel.ProvinceVm, err error) {
	repository := actions.NewProvinceRepository(uc.DB)
	provinces, err := repository.BrowseAll()
	if err != nil {
		return res, err
	}

	for _, province := range provinces {
		res = append(res, uc.buildBody(province))
	}

	return res, nil
}

func (uc ProvinceUseCase) buildBody(model models.Province) viewmodel.ProvinceVm {
	return viewmodel.ProvinceVm{
		ID:   model.ID,
		Name: model.Name,
	}
}
