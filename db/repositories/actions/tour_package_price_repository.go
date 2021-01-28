package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type TourPackagePriceRepository struct {
	DB *sql.DB
}

const (
	tourPackagePriceSelectStatement = `select id,room_type,room_capacity,price,promo_price,airline_class,tour_package_id,room_rate_id,created_at,updated_at`
)

func NewTourPackagePriceRepository(DB *sql.DB) contracts.ITourPackagePriceRepository {
	return &TourPackagePriceRepository{DB: DB}
}

func (repository TourPackagePriceRepository) scanRow(row *sql.Row) (res models.TourPackagePrice, err error) {
	if err = row.Scan(&res.ID, &res.RoomType, &res.RoomCapacity, &res.Price, &res.PromoPrice, &res.AirLineClass, &res.TourPackageID, &res.TourPackageID, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return res, err
	}

	return res, nil
}

func (repository TourPackagePriceRepository) ReadBy(column, value, operator string) (data models.TourPackagePrice, err error) {
	statement := tourPackagePriceSelectStatement + ` from tour_package_prices where ` + column + `` + operator + `$1 and deleted_at is null`
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
