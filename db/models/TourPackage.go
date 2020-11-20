package models

import (
	"database/sql"
	"time"
)

type TourPackage struct {
	ID                     string       `db:"id"`
	OdooPackageID          int32        `db:"odoo_package_id"`
	OdooPackageProgramID   int32        `db:"odoo_package_program_id"`
	PackageProgram         string       `db:"package_program"`
	OdooPackageProgramName string       `db:"odoo_package_program_name"`
	Name                   string       `db:"name"`
	DepartureDate          time.Time    `db:"departure_date"`
	ReturnDate             time.Time    `db:"return_date"`
	ProgramDay             int32        `db:"program_day"`
	Description            string       `db:"description"`
	Notes                  string       `db:"notes"`
	PartnerID              string       `db:"partner_id"`
	PartnerName            string       `db:"partner_name"`
	Hotels                 string       `db:"hotels"`
	Meals                  string       `db:"meals"`
	Airlines               string       `db:"airlines"`
	Busses                 string       `db:"busses"`
	Prices                 string       `db:"prices"`
	CreatedAt              time.Time    `db:"created_at"`
	UpdatedAt              time.Time    `db:"updated_at"`
	DeletedAt              sql.NullTime `db:"deleted_at"`
}
