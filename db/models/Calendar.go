package models

import "database/sql"

type Calendar struct {
	ID          string         `db:"id"`
	Title       string         `db:"title"`
	Start       string         `db:"start"`
	End         string         `db:"end"`
	Description string         `db:"description"`
	Remember    int            `json:"remember"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   string         `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}
