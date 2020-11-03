package models

type MenuPermissionUser struct {
	ID               string `db:"id"`
	MenuID           string `db:"menu_id"`
	MenuPermissionID string `db:"menu_permission_id"`
}
