package models

import "database/sql"

type WebComprofCategory struct {
	ID           string         `db:"id"`
	Slug         string         `db:"slug"`
	Name         string         `db:"name"`
	CategoryType string         `db:"category_type"`
	CreatedAt    string         `db:"created_at"`
	UpdatedAt    string         `db:"updated_at"`
	DeletedAt    sql.NullString `db:"deleted_at"`
}
