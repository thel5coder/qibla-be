package models

import "database/sql"

type User struct {
	ID        string         `db:"id"`
	UserName  string         `db:"user_name"`
	Email     string         `db:"email"`
	Password  string         `db:"password"`
	IsActive  bool           `db:"is_active"`
	RoleID    string         `db:"role_id"`
	RoleModel Role           `db:"role_model"`
	CreatedAt string         `db:"activation_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
