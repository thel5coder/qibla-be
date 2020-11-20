package models

import (
	"database/sql"
	"time"
)

type TourPackageHotel struct {
	ID                    string       `db:"id"`
	Name                  string       `db:"name"`
	Rating                int32        `db:"rating"`
	Location              string       `db:"location"`
	OdooProductTemplateID int32        `db:"odoo_product_template_id"`
	TourPackageID         string       `db:"tour_package_id"`
	CreatedAt             time.Time    `db:"created_at"`
	UpdatedAt             time.Time    `db:"updated_at"`
	DeletedAt             sql.NullTime `db:"deleted_at"`
}
