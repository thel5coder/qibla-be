package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IVideoContentRepository interface {
	Browse(order, sort string, limit, offset int) (data []models.VideoContent, count int, err error)

	BrowseAll() (data []models.VideoContent,err error)

	ReadBy(column, value string) (data models.VideoContent, err error)

	Edit(input viewmodel.VideoContentVm) (res string, err error)

	Add(input viewmodel.VideoContentVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
