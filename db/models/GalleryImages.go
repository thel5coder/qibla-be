package models

import "database/sql"

type GalleryImages struct {
	ID        string         `db:"id"`
	GalleryID string         `db:"gallery_id"`
	FileID    string         `db:"file_id"`
	Path      string         `db:"path"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
