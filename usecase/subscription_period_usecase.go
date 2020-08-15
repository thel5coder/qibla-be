package usecase

import (
	"database/sql"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type SubscriptionPeriodUseCase struct {
	*UcContract
}

func (uc SubscriptionPeriodUseCase) BrowseBySettingProductID(settingProductID string) (res []viewmodel.SubscriptionPeriodVm, err error) {
	repository := actions.NewSubscriptionPeriodRepository(uc.DB)
	subscriptionPeriods, err := repository.BrowseBySettingProductID(settingProductID)
	if err != nil {
		return res, err
	}

	for _, subscriptionPeriod := range subscriptionPeriods {
		res = append(res, viewmodel.SubscriptionPeriodVm{
			ID:     subscriptionPeriod.ID,
			Period: subscriptionPeriod.Period,
		})
	}

	return res, err
}

func (uc SubscriptionPeriodUseCase) Add(settingProductID string, period int, tx *sql.Tx) (err error) {
	repository := actions.NewSubscriptionPeriodRepository(uc.DB)
	err = repository.Add(settingProductID, period, tx)

	return err
}

func (uc SubscriptionPeriodUseCase) DeleteBySettingProductID(settingProductID string, tx *sql.Tx) (err error) {
	repository := actions.NewSubscriptionPeriodRepository(uc.DB)
	err = repository.DeleteBySettingProductID(settingProductID, tx)

	return err
}

func (uc SubscriptionPeriodUseCase) Store(settingProductID string, periods []int, tx *sql.Tx) (err error) {
	err = uc.DeleteBySettingProductID(settingProductID,tx)
	if err != nil {
		return err
	}

	for _,period := range periods{
		err = uc.Add(settingProductID,period,tx)
		if err != nil {
			return err
		}
	}

	return err
}
