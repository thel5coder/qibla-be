package models

import (
	"database/sql"
	"time"
)

type TourPackagePromotion struct {
	ID              string         `db:"id"`
	TourPackageID   string         `db:"tour_package_id"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
	DeletedAt       sql.NullTime   `db:"deleted_at"`
	TravelAgentID   string         `db:"travel_agent_id"`
	TravelAgentName string         `db:"travel_agent_name"`
	Branch          string         `db:"branch"`
	Phone           string         `db:"phone"`
	AreaCode        string         `db:"area_code"`
	TourPackage     TourPackage    `db:"tour_package"`
	Hotels          sql.NullString `db:"hotels"`
	RoomRates       sql.NullString `db:"room_rates"`
}
