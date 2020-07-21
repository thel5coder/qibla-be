package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IFileRepository interface {
	ReadBy(column,value string) (data models.File,err error)

	Add(input viewmodel.FileVm) (res string,err error)

	Delete(ID,updatedAt,deletedAt string) (res string,err error)

	CountBy(column,value string) (res int,err error)
}
