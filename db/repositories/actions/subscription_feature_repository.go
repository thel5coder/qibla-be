package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type SubscriptionFeatureRepository struct{
	DB *sql.DB
}

func NewSubscriptionFeatureRepository(DB *sql.DB) contracts.ISubscriptionFeatureRepository{
	return &SubscriptionFeatureRepository{DB: DB}
}

func (repository SubscriptionFeatureRepository) BrowseBySettingProductID(settingProductID string) (data []models.SubscriptionFeature, err error) {
	statement := `select * from "subscription_features" where "setting_product_id"=$1`
	rows,err := repository.DB.Query(statement,settingProductID)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := models.SubscriptionFeature{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.SettingProductID,
			&dataTemp.FeatureName,
		)
		if err != nil {
			return data,err
		}

		data = append(data,dataTemp)
	}

	return data,err
}

func (SubscriptionFeatureRepository) Add(settingProductID, featureName string, tx *sql.Tx) (err error) {
	statement :=`insert into "subscription_features" ("setting_product_id","feature_name") values($1,$2)`
	_,err = tx.Exec(statement,settingProductID,featureName)

	return err
}

func (SubscriptionFeatureRepository) DeleteBySettingProductID(settingProductID string, tx *sql.Tx) (err error) {
	statement := `delete from "subscription_periods" where "setting_product_id"=$1`
	_,err = tx.Exec(statement,settingProductID)

	return err
}
