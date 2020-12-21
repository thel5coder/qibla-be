package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IMenuRepository interface {
	Browse(parentID,search, order, sort string, limit, offset int) (data []models.Menu, count int, err error)

	BrowseAllBy(column,value,operator string,isActive bool) (data []models.Menu,err error)

	ReadBy(column, value,operator string) (data models.Menu, err error)

	Edit(input viewmodel.MenuVm, tx *sql.Tx) (err error)

	Add(input viewmodel.MenuVm, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	DeleteChild(parentID string,updatedAt,deletedAt string,tx *sql.Tx) (err error)

	CountBy(ID, column, value string) (res int, err error)

	CountByPk(ID string) (res int, err error)
}
