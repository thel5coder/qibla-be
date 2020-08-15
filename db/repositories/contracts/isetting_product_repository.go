package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ISettingProductRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.SettingProduct, count int, err error)

	ReadBy(column, value string) (data models.SettingProduct, err error)

	Edit(input viewmodel.SettingProductVm,tx *sql.Tx) (res string, err error)

	Add(input viewmodel.SettingProductVm,tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string,tx *sql.Tx) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
