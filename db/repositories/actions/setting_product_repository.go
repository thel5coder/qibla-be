package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/usecase/viewmodel"
)

type SettingProductRepository struct {
	DB *sql.DB
}

func NewSettingProductRepository(DB *sql.DB) contracts.ISettingProductRepository {
	return &SettingProductRepository{DB: DB}
}

func (SettingProductRepository) Browse(search, order, sort string, limit, offset int) (data []models.SettingProduct, count int, err error) {
	panic("implement me")
}

func (SettingProductRepository) ReadBy(column, value string) (data models.SettingProduct, err error) {
	panic("implement me")
}

func (SettingProductRepository) Edit(input viewmodel.SettingProductVm, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (SettingProductRepository) Add(input viewmodel.SettingProductVm, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (SettingProductRepository) Delete(ID, updatedAt, deletedAt string,tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (SettingProductRepository) CountBy(ID, column, value string) (res int, err error) {
	panic("implement me")
}
