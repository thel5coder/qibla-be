package models

import "database/sql"

type File struct {
	ID        string         `db:"id"`
	Name      string         `db:"name"`
	Path      string         `db:"path"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
