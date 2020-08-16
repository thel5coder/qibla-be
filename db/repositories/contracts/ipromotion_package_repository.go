package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IPackagePromotionRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.PromotionPackage, count int, err error)

	BrowseAll() (data []models.PromotionPackage,err error)

	ReadBy(column, value string) (data models.PromotionPackage, err error)

	Edit(input viewmodel.PromotionPackageVm) (res string, err error)

	Add(input viewmodel.PromotionPackageVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
