package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IGlobalInfoCategorySettingRepository interface {
	Browse(globalInfoCategorySlug,search, order, sort string, limit, offset int) (data []models.GlobalInfoCategorySetting, count int, err error)

	ReadBy(column, value string) (data models.GlobalInfoCategorySetting, err error)

	Edit(input viewmodel.GlobalInfoCategorySettingVm) (res string, err error)

	Add(input viewmodel.GlobalInfoCategorySettingVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
