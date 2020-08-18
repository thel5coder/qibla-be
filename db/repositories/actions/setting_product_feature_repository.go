package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type SettingProductFeatureRepository struct{
	DB *sql.DB
}

func NewSettingProductFeatureRepository(DB *sql.DB) contracts.ISettingProductFeatureRepository {
	return &SettingProductFeatureRepository{DB: DB}
}

func (repository SettingProductFeatureRepository) BrowseBySettingProductID(settingProductID string) (data []models.SubscriptionFeature, err error) {
	statement := `select * from setting_product_features where "setting_product_id"=$1`
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

func (SettingProductFeatureRepository) Add(settingProductID, featureName string, tx *sql.Tx) (err error) {
	statement :=`insert into setting_product_features ("setting_product_id","feature_name") values($1,$2)`
	_,err = tx.Exec(statement,settingProductID,featureName)

	return err
}

func (SettingProductFeatureRepository) DeleteBySettingProductID(settingProductID string, tx *sql.Tx) (err error) {
	statement := `delete from "setting_product_features" where "setting_product_id"=$1`
	_,err = tx.Exec(statement,settingProductID)

	return err
}
