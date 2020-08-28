package models

import "database/sql"

type SatisfactionCategory struct {
	ID          string         `db:"id"`
	ParentID    sql.NullString `db:"parent_id"`
	Slug        string         `db:"slug"`
	Name        string         `db:"name"`
	IsActive    bool           `db:"is_active"`
	Description sql.NullString `json:"description"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   string         `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}
