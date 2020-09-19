package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type ISettingProductPeriodRepository interface {
	BrowseBySettingProductID(settingProductID string) (data []models.SubscriptionPeriod,err error)

	Add(settingProductID string, period int,tx *sql.Tx) (err error)

	DeleteBySettingProductID(settingProductID string,tx *sql.Tx) (err error)
}
