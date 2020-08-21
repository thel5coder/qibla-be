package models

import "database/sql"

type Pray struct {
	ID        string         `db:"id"`
	Name      string         `db:"name"`
	FileID    string         `db:"file_id"`
	IsActive  bool           `db:"is_active"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
