package models

import "database/sql"

type Menu struct {
	ID          string         `db:"id"`
	MenuID      string         `db:"menu_id"`
	Name        string         `db:"name"`
	Url         string         `db:"url"`
	ParentID    sql.NullString `db:"parent_id"`
	IsActive    bool           `db:"is_active"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   string         `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
	Permissions sql.NullString `db:"permissions"`
}
