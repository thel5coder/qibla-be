package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IVideoContentRepository interface {
	Browse(order, sort string, limit, offset int) (data []models.VideoContent, count int, err error)

	Add(input viewmodel.VideoContentVm) (res string, err error)
}
