package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type CityUseCase struct {
	*UcContract
}

func (uc CityUseCase) BrowseAllByProvince(provinceID string) (res []viewmodel.CityVm, err error) {
	repository := actions.NewCityRepository(uc.DB)
	cities, err := repository.BrowseAllByProvince(provinceID)
	if err != nil {
		return res, err
	}

	for _, city := range cities {
		res = append(res, uc.buildBody(city))
	}

	return res, nil
}

func (uc CityUseCase) buildBody(model models.City) viewmodel.CityVm {
	return viewmodel.CityVm{
		ID:         model.ID,
		Name:       model.Name,
		ProvinceID: model.ProvinceID,
	}
}
