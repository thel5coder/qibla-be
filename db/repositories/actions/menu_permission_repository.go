package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"time"
)

type MenuPermissionRepository struct{
	DB *sql.DB
}

func NewMenuPermissionRepository(DB *sql.DB) contracts.IMenuPermissionRepository{
	return &MenuPermissionRepository{DB: DB}
}

func (repository MenuPermissionRepository) Browse(menuID string) (data []models.MenuPermission, err error) {
	statement := `select * from "menu_permissions" where "menu_id"=$1 and "deleted_at" is null`
	rows,err := repository.DB.Query(statement,menuID)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := models.MenuPermission{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.MenuID,
			&dataTemp.Permission,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			)
		if err != nil {
			return data,err
		}
		data = append(data,dataTemp)
	}

	return data,err
}

func (MenuPermissionRepository) Add(menuID, permission, createdAt, updatedAt string, tx *sql.Tx) (err error) {
	fmt.Print(menuID)
	statement := `insert into "menu_permissions" ("menu_id","permission","created_at","updated_at") values($1,$2,$3,$4)`
	_,err = tx.Exec(statement,menuID,permission,datetime.StrParseToTime(createdAt,time.RFC3339),datetime.StrParseToTime(updatedAt,time.RFC3339))

	return err
}

func (MenuPermissionRepository) Edit(ID, permission, updatedAt string, tx *sql.Tx) (err error) {
	statement := `update "menu_permissions" set "permission"=$1,"updated_at"=$2 where "id"=$3`
	_,err = tx.Exec(statement,permission,datetime.StrParseToTime(updatedAt,time.RFC3339),ID)

	return err
}

func (MenuPermissionRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "menu_permissions" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3`
	_,err = tx.Exec(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID)

	return err
}

