package models

import "database/sql"

type MasterPromotion struct {
	ID        string         `db:"id"`
	Slug      string         `db:"slug"`
	Name      string         `db:"name"`
	IsActive  bool           `db:"is_active"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
