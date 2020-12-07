package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IMasterZakatRepository interface {
	Browse(search map[string]interface{}, order, sort string, limit, offset int) (data []models.MasterZakat, count int, err error)

	BrowseAll() (data []models.MasterZakat, err error)

	ReadBy(column, value string) (data models.MasterZakat, err error)

	Edit(input viewmodel.MasterZakatVm) (res string, err error)

	Add(input viewmodel.MasterZakatVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
