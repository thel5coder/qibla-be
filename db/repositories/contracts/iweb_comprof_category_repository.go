package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IWebComprofCategoryRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.WebComprofCategory, count int, err error)

	ReadBy(column, value string) (data models.WebComprofCategory, err error)

	Edit(input viewmodel.WebComprofCategoryVm) (res string, err error)

	Add(input viewmodel.WebComprofCategoryVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
