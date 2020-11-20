package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"time"
)

type ITourPackageRepository interface {
	BrowseAllBy(column, value, operator string) (data []models.TourPackage, err error)

	ReadBy(column, value, operator string) (data models.TourPackage, err error)

	Edit(input models.TourPackage, tx *sql.Tx) (err error)

	Add(input models.TourPackage, tx *sql.Tx) (res string, err error)

	DeleteBy(column, value, operator string, updatedAt, deletedAt time.Time, tx *sql.Tx) (err error)
}
