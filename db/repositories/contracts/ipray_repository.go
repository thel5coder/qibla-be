package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IPrayRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.Pray, count int, err error)

	ReadBy(column, value string) (data models.Pray, err error)

	Edit(input viewmodel.PrayVm) (res string, err error)

	Add(input viewmodel.PrayVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
