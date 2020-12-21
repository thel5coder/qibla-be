package models

import "database/sql"

type User struct {
	ID                 string         `db:"id"`
	ProfilePictureID   sql.NullString `db:"profile_picture_id"`
	UserName           string         `db:"user_name"`
	Email              sql.NullString `db:"email"`
	Name               sql.NullString `db:"name"`
	Password           string         `db:"password"`
	MobilePhone        sql.NullString `db:"mobile_phone"`
	PIN                sql.NullString `db:"pin"`
	IsActive           bool           `db:"is_active"`
	IsAdminPanel       sql.NullBool   `db:"is_admin_panel"`
	IsFingerprintSet   sql.NullBool   `db:"is_fingerprint_set"`
	OdooUserID         sql.NullInt32  `db:"odoo_user_id"`
	FcmDeviceToken     sql.NullString `db:"fcm_device_token"`
	CreatedAt          string         `db:"activation_at"`
	UpdatedAt          string         `db:"updated_at"`
	DeletedAt          sql.NullString `db:"deleted_at"`
	RoleModel          Role           `db:"role_model"`
	FileModel          File           `db:"file_model"`
	MenuUser           sql.NullString `db:"menu_user"`
	MenuPermissionUser sql.NullString `db:"menu_permission_user"`
}
