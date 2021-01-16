package contracts

import "qibla-backend/db/models"

type ITourPackagePromotionRepository interface {
	Browse(filters map[string]interface{},order,sort string, limit, offset int) (data []models.TourPackagePromotion,count int,err error)

	ReadBy(column,value,operator string) (data models.TourPackagePromotion,err error)
}
