package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IGalleryRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.Gallery, count int, err error)

	ReadBy(column, value string) (data models.Gallery, err error)

	Edit(ID,galleryName,updatedAt string,tx *sql.Tx) (err error)

	Add(webContentCategoryID,galleryName,createdAt,updatedAt string,tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string,tx *sql.Tx) (err error)

	CountBy(ID, column, value string) (res int, err error)
}
