package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IPromotionRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.Promotion, count int, err error)

	ReadBy(column, value string) (data models.Promotion, err error)

	Edit(input viewmodel.PromotionVm) (res string, err error)

	Add(input viewmodel.PromotionVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
