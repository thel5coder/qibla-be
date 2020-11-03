package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IMenuUserPermissionRepository interface {
	Browse(menuID string) (data []models.MenuUserPermission, err error)

	Add(menuID, menuPermissionID string, tx *sql.Tx) (err error)

	Delete(menuID string, tx *sql.Tx) (err error)
}
