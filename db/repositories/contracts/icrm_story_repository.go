package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ICrmStoryRepository interface {
	BrowseAll() (data []models.CrmStory,err error)

	ReadBy(column, value string) (data models.CrmStory, err error)

	Edit(input viewmodel.CrmStoryVm) (res string, err error)

	Add(input viewmodel.CrmStoryVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
