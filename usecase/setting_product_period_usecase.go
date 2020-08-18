package usecase

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type SettingProductPeriodUseCase struct {
	*UcContract
}

func (uc SettingProductPeriodUseCase) BrowseBySettingProductID(settingProductID string) (res []viewmodel.SettingProductPeriodVm, err error) {
	repository := actions.NewSettingProductPeriodRepository(uc.DB)
	subscriptionPeriods, err := repository.BrowseBySettingProductID(settingProductID)
	if err != nil {
		return res, err
	}

	for _, subscriptionPeriod := range subscriptionPeriods {
		res = append(res, viewmodel.SettingProductPeriodVm{
			ID:     subscriptionPeriod.ID,
			Period: subscriptionPeriod.Period,
		})
	}

	return res, err
}

func (uc SettingProductPeriodUseCase) Add(settingProductID string, period int, tx *sql.Tx) (err error) {
	repository := actions.NewSettingProductPeriodRepository(uc.DB)
	err = repository.Add(settingProductID, period, tx)

	return err
}

func (uc SettingProductPeriodUseCase) DeleteBySettingProductID(settingProductID string, tx *sql.Tx) (err error) {
	repository := actions.NewSettingProductPeriodRepository(uc.DB)
	err = repository.DeleteBySettingProductID(settingProductID, tx)

	return err
}

func (uc SettingProductPeriodUseCase) Store(settingProductID string, periods []int, tx *sql.Tx) (err error) {
	rows, _ := uc.BrowseBySettingProductID(settingProductID)

	if len(rows) > 0 {
		err = uc.DeleteBySettingProductID(settingProductID, tx)
		if err != nil {
			return err
		}
	}

	for _, period := range periods {
		fmt.Println(settingProductID)
		err = uc.Add(settingProductID, period, tx)
		if err != nil {
			return err
		}
	}

	return nil
}
