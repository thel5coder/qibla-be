package models

import "database/sql"

type MenuPermissionUser struct {
	ID               string         `db:"id"`
	UserID           string         `db:"user_id"`
	MenuPermissionID string         `db:"menu_permission_id"`
	Permission       string         `db:"permission"`
	MenuID           string         `db:"menu_id"`
	MenuName         string         `db:"menu_name"`
	CreatedAt        string         `db:"created_at"`
	UpdatedAt        string         `db:"updated_at"`
	DeletedAt        sql.NullString `db:"deleted_at"`
}
