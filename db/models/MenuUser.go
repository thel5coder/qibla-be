package models

type MenuUser struct {
	ID              string `db:"id"`
	UserID          string `db:"user_id"`
	MenuID          string `db:"menu_id"`
	MenuPermissions string `db:"menu_permissions"`
}
