package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IGalleryImageRepository interface {
	Browse(galleryID string) (data []models.GalleryImages, err error)

	ReadBy(column, value string) (data models.GalleryImages, err error)

	Edit(ID, fileID, updatedAt string, tx *sql.Tx) (err error)

	Add(galleryID, fileID, createdAt, updatedAt string, tx *sql.Tx) (err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	DeleteByGalleryID(galleryID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountBy(ID, galleryID, column, value string) (res int, err error)
}
