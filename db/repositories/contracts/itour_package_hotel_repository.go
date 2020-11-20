package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"time"
)

type ITourPackageHotelRepository interface {
	Edit(input models.TourPackageHotel, tx *sql.Tx) (err error)

	Add(input models.TourPackageHotel, tx *sql.Tx) (err error)

	DeleteBy(column, value, operator string, updatedAt, deletedAt time.Time, tx *sql.Tx) (err error)

	CountBy(tourPackageID, column, value, operator string) (res int, err error)
}
