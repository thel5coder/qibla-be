package models

import (
	"database/sql"
	"time"
)

type UserTourPurchaseRoom struct {
	ID                 string       `db:"id"`
	UserTourPurchaseID string       `db:"user_tour_purchase_id"`
	TourPackagePriceID string       `db:"tour_package_price_id"`
	Price              int64        `db:"price"`
	Quantity           int          `db:"quantity"`
	CreatedAt          time.Time    `db:"created_at"`
	UpdatedAt          time.Time    `db:"updated_at"`
	DeletedAt          sql.NullTime `db:"deleted_at"`
}
