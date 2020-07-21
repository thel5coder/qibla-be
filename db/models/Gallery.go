package models

import "database/sql"

type Gallery struct {
	ID                   string         `db:"id"`
	WebContentCategoryID string         `db:"web_content_category_id"`
	GalleryName          string         `db:"gallery_name"`
	CreatedAt            string         `db:"created_at"`
	UpdatedAt            string         `db:"updated_at"`
	DeletedAt            sql.NullString `db:"deleted_at"`
}
