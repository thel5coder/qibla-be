package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IMenuPermissionUserRepository interface {
	Browse(userID string) (data []models.MenuPermissionUser, err error)

	Add(userID, menuPermissionID, createdAt, updatedAt string, tx *sql.Tx) (err error)

	Delete(userID, menuPermissionID, updatedAt, deletedAt string, tx *sql.Tx) (err error)
}
