package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IMenuPermissionRepository interface {
	Browse(menuID string) (data []models.MenuPermission, err error)

	Add(menuID, permission, createdAt, updatedAt string, tx *sql.Tx) (err error)

	Edit(ID, permission, updatedAt string, tx *sql.Tx) (err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error)
}
