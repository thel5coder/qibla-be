package models

import "database/sql"

type Promotion struct {
	ID                 string         `db:"id"`
	PackageName        string         `json:"package_name"`
	PromotionPackageID string         `db:"promotion_package_id"`
	PackagePromotion   string         `db:"package_promotion"`
	StartDate          string         `db:"start_date"`
	EndDate            string         `db:"end_date"`
	Platform           string         `db:"platform"`
	Position           string         `db:"position"`
	Price              string         `db:"price"`
	Description        string         `db:"description"`
	IsActive           bool           `db:"is_active"`
	CreatedAt          string         `db:"created_at"`
	UpdatedAt          string         `db:"updated_at"`
	DeletedAt          sql.NullString `db:"deleted_at"`
}
