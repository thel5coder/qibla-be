package usecase

import (
	"database/sql"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type SettingProductFeatureUseCase struct {
	*UcContract
}

func (uc SettingProductFeatureUseCase) BrowseBySettingProductID(settingProductID string) (res []viewmodel.SubscriptionFeatureVm, err error) {
	repository := actions.NewSettingProductFeatureRepository(uc.DB)
	subscriptionFeatures, err := repository.BrowseBySettingProductID(settingProductID)
	if err != nil {
		return res, err
	}

	for _, subscriptionFeature := range subscriptionFeatures {
		res = append(res, viewmodel.SubscriptionFeatureVm{
			ID:          subscriptionFeature.ID,
			FeatureName: subscriptionFeature.FeatureName,
		})
	}

	return res, err
}

func (uc SettingProductFeatureUseCase) Add(settingProductID,featureName string, tx *sql.Tx) (err error) {
	repository := actions.NewSettingProductFeatureRepository(uc.DB)
	err = repository.Add(settingProductID, featureName, tx)

	return err
}

func (uc SettingProductFeatureUseCase) DeleteBySettingProductID(settingProductID string, tx *sql.Tx) (err error) {
	repository := actions.NewSettingProductFeatureRepository(uc.DB)
	err = repository.DeleteBySettingProductID(settingProductID, tx)

	return err
}

func (uc SettingProductFeatureUseCase) Store(settingProductID string, features []string, tx *sql.Tx) (err error) {
	err = uc.DeleteBySettingProductID(settingProductID,tx)
	if err != nil {
		return err
	}

	for _,feature := range features{
		err = uc.Add(settingProductID,feature,tx)
		if err != nil {
			return err
		}
	}

	return err
}
