package usecase

import (
	"database/sql"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type SubscriptionFeatureUseCase struct {
	*UcContract
}

func (uc SubscriptionFeatureUseCase) BrowseBySettingProductID(settingProductID string) (res []viewmodel.SubscriptionFeatureVm, err error) {
	repository := actions.NewSubscriptionFeatureRepository(uc.DB)
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

func (uc SubscriptionFeatureUseCase) Add(settingProductID,featureName string, tx *sql.Tx) (err error) {
	repository := actions.NewSubscriptionFeatureRepository(uc.DB)
	err = repository.Add(settingProductID, featureName, tx)

	return err
}

func (uc SubscriptionFeatureUseCase) DeleteBySettingProductID(settingProductID string, tx *sql.Tx) (err error) {
	repository := actions.NewSubscriptionFeatureRepository(uc.DB)
	err = repository.DeleteBySettingProductID(settingProductID, tx)

	return err
}

func (uc SubscriptionFeatureUseCase) Store(settingProductID string, features []string, tx *sql.Tx) (err error) {
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
