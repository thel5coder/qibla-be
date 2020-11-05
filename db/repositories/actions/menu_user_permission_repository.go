package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type MenuUserPermissionRepository struct {
	DB *sql.DB
}

func NewMenuUserPermissionRepository(DB *sql.DB) contracts.IMenuUserPermissionRepository {
	return &MenuUserPermissionRepository{DB: DB}
}

const (
	menuUserPermissionSelect = `select "id","menu_id","menu_permission_id"`
)

func (repository MenuUserPermissionRepository) scanRows(rows *sql.Rows) (res models.MenuUserPermission, err error) {
	err = rows.Scan(&res.ID, &res.MenuID, &res.MenuPermissionID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository MenuUserPermissionRepository) Browse(menuID string) (data []models.MenuUserPermission, err error) {
	statement := menuUserPermissionSelect + ` from "menu_user_permissions" where "menu_id"=$1`
	rows, err := repository.DB.Query(statement, menuID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

func (MenuUserPermissionRepository) Add(menuID, menuPermissionID string, tx *sql.Tx) (err error) {
	statement := `insert into "menu_user_permissions" ("menu_id","menu_permission_id") values($1,$2)`
	_, err = tx.Exec(statement, menuID, menuPermissionID)
	if err != nil {
		return err
	}

	return nil
}

func (MenuUserPermissionRepository) Delete(menuID string, tx *sql.Tx) (err error) {
	statement := `delete from "menu_user_permissions" where "menu_id"=$1`
	_, err = tx.Exec(statement, menuID)
	if err != nil {
		return err
	}

	return nil
}
