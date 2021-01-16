package models

import (
	"database/sql"
	"time"
)

type TourPackage struct {
	ID                       string         `json:"id"`
	OdooPackageID            string         `json:"odoo_package_id"`
	Name                     string         `json:"name"`
	DepartureDate            time.Time      `json:"departure_date"`
	ReturnDate               time.Time      `json:"return_date"`
	Description              sql.NullString `json:"description"`
	CreatedAt                time.Time      `json:"created_at"`
	UpdatedAt                time.Time      `json:"updated_at"`
	DeletedAt                sql.NullTime   `json:"deleted_at"`
	PartnerID                string         `json:"partner_id"`
	PackageType              string         `json:"package_type"`
	ProgramDay               int            `json:"program_day"`
	Notes                    sql.NullString `json:"notes"`
	Image                    sql.NullString `json:"image"`
	DepartureAirport         string         `json:"departure_airport"`
	DestinationAirport       string         `json:"destination_airport"`
	ReturnDepartureAirport   string         `json:"return_departure_airport"`
	ReturnDestinationAirport string         `json:"return_destination_airport"`
	Quota                    int            `json:"quota"`
}
