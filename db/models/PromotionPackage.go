package models

import "database/sql"

type PromotionPackage struct {
	ID          string         `db:"id"`
	Slug        string         `db:"slug"`
	PackageName string         `db:"package_name"`
	IsActive    bool           `db:"is_active"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   string         `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}
