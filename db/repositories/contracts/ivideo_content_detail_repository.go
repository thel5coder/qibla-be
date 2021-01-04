package contracts

import (
	"qibla-backend/db/models"
)

type IVideoContentDetailRepository interface {
	BrowseByVideoContentID(videoContentID, search, order, sort string, limit, offset int) (data []models.VideoContentDetails,count int,err error)

	Read(ID string) (data models.VideoContentDetails,err error)

	Add(model models.VideoContentDetails) (res string,err error)

	DeleteBy(column,value,operator string, model models.VideoContentDetails) (res string,err error)

	CountBy(column,value,operator string) (res int,err error)
}
