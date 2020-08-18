package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type SettingProductPeriodRepository struct{
	DB *sql.DB
}

func NewSettingProductPeriodRepository(DB *sql.DB) contracts.ISettingProductPeriodRepository {
	return &SettingProductPeriodRepository{DB: DB}
}

func (repository SettingProductPeriodRepository) BrowseBySettingProductID(settingProductID string) (data []models.SubscriptionPeriod, err error) {
	statement := `select * from "setting_product_periods" where "setting_product_id"=$1`
	rows,err := repository.DB.Query(statement,settingProductID)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := models.SubscriptionPeriod{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.SettingProductID,
			&dataTemp.Period,
			)
		if err != nil {
			return data,err
		}

		data = append(data,dataTemp)
	}

	return data,err
}

func (SettingProductPeriodRepository) Add(settingProductID string, period int, tx *sql.Tx) (err error) {
	statement :=`insert into "setting_product_periods" ("setting_product_id","period") values($1,$2)`
	_,err = tx.Exec(statement,settingProductID,period)

	return err
}

func (SettingProductPeriodRepository) DeleteBySettingProductID(settingProductID string, tx *sql.Tx) (err error) {
	statement := `delete from "setting_product_periods" where "setting_product_id"=$1`
	_,err = tx.Exec(statement,settingProductID)

	return err
}

