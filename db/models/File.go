package models

import "database/sql"

type File struct {
	ID        string         `db:"id"`
	Name      sql.NullString `db:"name"`
	Path      sql.NullString `db:"path"`
	UserID    sql.NullString `db:"user_id"`
	Type      sql.NullString `db:"type"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
