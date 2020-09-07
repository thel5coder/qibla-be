package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IMasterProductRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.MasterProduct, count int, err error)

	BrowseAll() (data []models.MasterProduct,err error)

	BrowseExtraProducts() (data []models.MasterProduct,err error)

	ReadBy(column, value string) (data models.MasterProduct, err error)

	Edit(input viewmodel.MasterProductVm) (res string, err error)

	Add(input viewmodel.MasterProductVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}