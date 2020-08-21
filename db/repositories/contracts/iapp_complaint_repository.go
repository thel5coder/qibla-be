package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IAppComplaintRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.AppComplaint, count int, err error)

	ReadBy(column, value string) (data models.AppComplaint, err error)

	Edit(input viewmodel.AppComplaintVm) (res string, err error)

	Add(input viewmodel.AppComplaintVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)

	CountAll() (res int,err error)
}
