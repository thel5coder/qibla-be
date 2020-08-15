package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type ISubscriptionFeatureRepository interface {
	BrowseBySettingProductID(settingProductID string) (data []models.SubscriptionFeature,err error)

	Add(settingProductID,featureName string,tx *sql.Tx) (err error)

	DeleteBySettingProductID(settingProductID string,tx *sql.Tx) (err error)
}
