package contracts

import "qibla-backend/db/models"

type ICityRepository interface {
	BrowseAllByProvince(provinceID string) (data []models.City,err error)
}
