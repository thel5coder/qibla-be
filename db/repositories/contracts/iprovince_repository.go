package contracts

import "qibla-backend/db/models"

type IProvinceRepository interface {
	BrowseAll() (data []models.Province,err error)
}
