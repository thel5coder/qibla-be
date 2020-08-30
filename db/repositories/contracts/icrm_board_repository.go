package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ICrmBoardRepository interface {
	BrowseByCrmStoryID(crmStoryID string) (data []models.CrmBoard, err error)

	ReadBy(column, value string) (data models.CrmBoard, err error)

	Edit(input viewmodel.CrmBoardVm) (res string, err error)

	EditBoardStory(ID, crmStoryID, updatedAt string) (res string, err error)

	Add(input viewmodel.CrmBoardVm) (res string, err error)

	DeleteBy(column, value, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, crmStoryID, column, value string) (res int, err error)
}
