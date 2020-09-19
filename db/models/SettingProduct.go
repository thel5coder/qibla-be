package models

import "database/sql"

// SettingProduct ...
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
	Features            sql.NullString `db:"features"`
	Periods             sql.NullString `db:"periods"`
	CreatedAt           string         `db:"created_at"`
	UpdatedAt           string         `db:"updated_at"`
	DeletedAt           sql.NullString `db:"deleted_at"`
}

var (
	// SettingProductSelect ...
	SettingProductSelect = `select sp.*, mp."name", mp."subscription_type",
	array_to_string(array_agg(spf.id || '#' || spf.feature_name),'|') as features,
	array_to_string(array_agg(spp.id || '#' || spp.period),'|') as periods
	from "setting_products" sp 
	inner join "master_products" mp on mp."id" = sp."product_id" and mp."deleted_at" is null
	LEFT JOIN "setting_product_features" spf ON spf."setting_product_id" = sp."id"
	LEFT JOIN "setting_product_periods" spp ON spp."setting_product_id" = sp."id"`
	// SettingProductGroup ...
	SettingProductGroup = `GROUP BY sp."id", sp."product_id", sp."price", sp."price_unit", sp."maintenance_price",
	sp."discount", sp."discount_type", sp."discount_period_start", sp."discount_period_end",
	sp."description", sp."sessions", sp."created_at", sp."updated_at", sp."deleted_at",
	mp."name", mp."subscription_type"`
)
