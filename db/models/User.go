package models

import "database/sql"

type User struct {
	ID             string         `db:"id"`
	UserName       string         `db:"user_name"`
	Email          string         `db:"email"`
	Name           sql.NullString `db:"name"`
	ProfilePicture sql.NullString `db:"profile_picture"`
	Password       string         `db:"password"`
	MobilePhone    sql.NullString `db:"mobile_phone"`
	PIN            sql.NullString `db:"pin"`
	IsActive       bool           `db:"is_active"`
	IsAdminPanel   bool           `db:"is_admin_panel"`
	RoleID         sql.NullString `db:"role_id"`
	RoleModel      Role           `db:"role_model"`
	OdooUserID     sql.NullString `db:"odoo_user_id"`
	CreatedAt      string         `db:"activation_at"`
	UpdatedAt      string         `db:"updated_at"`
	DeletedAt      sql.NullString `db:"deleted_at"`
}
