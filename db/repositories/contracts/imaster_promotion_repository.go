package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IMasterPromotionRepository interface {
	Browse(search map[string]interface{}, order, sort string, limit, offset int) (data []models.MasterPromotion, count int, err error)

	BrowseAll() (data []models.MasterPromotion,err error)

	ReadBy(column, value string) (data models.MasterPromotion, err error)

	Edit(input viewmodel.MasterPromotionVm) (res string, err error)

	Add(input viewmodel.MasterPromotionVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
