package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ISettingPromotionRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.Promotion, count int, err error)

	BrowseAll(filters map[string]interface{}) (data []models.Promotion, err error)

	ReadBy(column, value string) (data models.Promotion, err error)

	Edit(input viewmodel.PromotionTodayVm, tx *sql.Tx) (res string, err error)

	Add(input viewmodel.PromotionTodayVm, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (res string, err error)

	CountBy(ID, promotionPackageID, column, value string) (res int, err error)
}
