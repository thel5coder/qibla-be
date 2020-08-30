package models

import "database/sql"

type CrmStory struct {
	ID        string         `db:"id"`
	Slug      string         `db:"slug"`
	Name      string         `db:"name"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
