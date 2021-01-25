package contracts

import "qibla-backend/db/models"

type ITourPackagePriceRepository interface {
	ReadBy(column, value, operator string) (data models.TourPackagePrice, err error)
}
