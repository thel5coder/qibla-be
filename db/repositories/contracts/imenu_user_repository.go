package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IMenuUserRepository interface {
	BrowseBy(column,value,operator string) (data []models.MenuUser,err error)

	Add(UserID,menuID string,tx *sql.Tx) (err error)

	Delete(UserID string,tx *sql.Tx) (err error)
}
