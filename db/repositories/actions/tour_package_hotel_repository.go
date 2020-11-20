package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"time"
)

type TourPackageHotelRepository struct {
	DB *sql.DB
}

func (TourPackageHotelRepository) Edit(input models.TourPackageHotel, tx *sql.Tx) (err error) {
	panic("implement me")
}

func (TourPackageHotelRepository) Add(input models.TourPackageHotel, tx *sql.Tx) (err error) {
	panic("implement me")
}

func (TourPackageHotelRepository) DeleteBy(column, value, operator string, updatedAt, deletedAt time.Time, tx *sql.Tx) (err error) {
	panic("implement me")
}

func (repository TourPackageHotelRepository) CountBy(tourPackageID, column, value, operator string) (res int, err error) {
	statement := `select count(id) from tour_package_hotels where tour_package_id=$1 and ` + column + `` + operator + `$2`
	err = repository.DB.QueryRow(statement, tourPackageID, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
