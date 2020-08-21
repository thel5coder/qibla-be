package models

import "database/sql"

type VideoContent struct {
	ID        string         `db:"id"`
	Channel   string         `db:"channel"`
	Links     string         `db:"links"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
