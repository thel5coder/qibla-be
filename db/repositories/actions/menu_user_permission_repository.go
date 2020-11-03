package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"time"
)

type MenuPermissionUserRepository struct {
	DB *sql.DB
}

func NewMenuPermissionUserRepository(DB *sql.DB) contracts.IMenuPermissionUserRepository {
	return &MenuPermissionUserRepository{DB: DB}
}

func (repository MenuPermissionUserRepository) Browse(userID string) (data []models.MenuPermissionUser, err error) {
	statement := `select mpu.*, mp."permission",m."id",m."name" from menu_user_permissions mpu 
                  inner join "menu_permissions" mp on mp."id"=mpu."menu_permission_id" and mp."deleted_at" is null
                  inner join "menus" m on m."id"=mp."menu_id" and m."deleted_at" is null
                  where mpu.menu_user_id=$1 and mpu."deleted_at" is null`
	rows, err := repository.DB.Query(statement, userID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.MenuPermissionUser{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.UserID,
			&dataTemp.MenuPermissionID,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.Permission,
			&dataTemp.MenuID,
			&dataTemp.MenuName,
		)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}

	return data, err
}

func (MenuPermissionUserRepository) Add(userID, menuPermissionID, createdAt, updatedAt string, tx *sql.Tx) (err error) {
	statement := `insert into menu_user_permissions (menu_user_id,"menu_permission_id","created_at","updated_at") values($1,$2,$3,$4)`
	_, err = tx.Exec(statement, userID, menuPermissionID, datetime.StrParseToTime(createdAt, time.RFC3339), datetime.StrParseToTime(updatedAt, time.RFC3339))

	return err
}

func (MenuPermissionUserRepository) Delete(userID, menuPermissionID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update menu_user_permissions set "updated_at"=$1,"deleted_at"=$2 where menu_user_id=$3 and "menu_permission_id"=$4`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), userID, menuPermissionID)

	return err
}

func (MenuPermissionUserRepository) DeleteByUser(userID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update menu_user_permissions set "updated_at"=$1,"deleted_at"=$2 where menu_user_id=$3 `
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), userID)

	return err
}
