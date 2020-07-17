package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IGlobalInfoCategoryRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.GlobalInfoCategory, count int, err error)

	ReadBy(column, value string) (data models.GlobalInfoCategory, err error)

	Edit(input viewmodel.GlobalInfoCategoryVm) (res string, err error)

	Add(input viewmodel.GlobalInfoCategoryVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
