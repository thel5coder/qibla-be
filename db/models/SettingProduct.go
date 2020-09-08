package models

import "database/sql"

type SettingProduct struct {
	ID                  string         `db:"id"`
	ProductID           string         `db:"product_id"`
	ProductName         string         `db:"product_name"`
	ProductType         string         `db:"product_type"`
	Price               int            `db:"price"`
	PriceUnit           sql.NullString `db:"price_unit"`
	MaintenancePrice    sql.NullInt32  `db:"maintenance_price"`
	Discount            sql.NullInt32  `db:"discount"`
	DiscountType        sql.NullString `db:"discount_type"`
	DiscountPeriodStart sql.NullString `db:"discount_period_start"`
	DiscountPeriodEnd   sql.NullString `db:"discount_period_end"`
	Description         string         `db:"description"`
	Sessions            sql.NullString `db:"sessions"`
	CreatedAt           string         `db:"created_at"`
	UpdatedAt           string         `db:"updated_at"`
	DeletedAt           sql.NullString `db:"deleted_at"`
}
