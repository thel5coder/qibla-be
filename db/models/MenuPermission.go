package models

import "database/sql"

type MenuPermission struct {
	ID         string         `db:"id"`
	MenuID     string         `db:"menu_id"`
	Permission string         `db:"permission"`
	CreatedAt  string         `db:"created_at"`
	UpdatedAt  string         `db:"updated_at"`
	DeletedAt  sql.NullString `db:"deleted_at"`
}
