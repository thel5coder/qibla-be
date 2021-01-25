package models

import (
	"database/sql"
	"time"
)

type TourPackagePrice struct {
	ID            string       `db:"id"`
	RoomType      string       `db:"room_type"`
	RoomCapacity  int          `db:"room_capacity"`
	Price         int64        `db:"price"`
	PromoPrice    int64        `db:"promo_price"`
	AirLineClass  string       `db:"air_line_class"`
	TourPackageID string       `db:"tour_package_id"`
	RoomRateID    string       `db:"room_rate_id"`
	CreatedAt     time.Time    `db:"created_at"`
	UpdatedAt     time.Time    `db:"updated_at"`
	DeletedAt     sql.NullTime `db:"deleted_at"`
}
