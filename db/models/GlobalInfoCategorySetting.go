package models

import "database/sql"

type GlobalInfoCategorySetting struct {
	ID                     string         `db:"id"`
	GlobalInfoCategoryID   string         `db:"global_info_category_id"`
	GlobalInfoCategoryName string         `db:"global_info_category_name"`
	Description            string         `db:"description"`
	IsActive               bool           `db:"is_active"`
	CreatedAt              string         `db:"created_at"`
	UpdatedAt              string         `db:"updated_at"`
	DeletedAt              sql.NullString `db:"deleted_at"`
}
