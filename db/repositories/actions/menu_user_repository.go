package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type MenuUserRepository struct {
	DB *sql.DB
}

func NewMenuUserRepository(DB *sql.DB) contracts.IMenuUserRepository {
	return &MenuUserRepository{DB: DB}
}

const (
	selectMenuUser = `select "id","user_id","menu_id"`
)

func (repository MenuUserRepository) scanRows(rows *sql.Rows) (res models.MenuUser, err error) {
	err = rows.Scan(&res.ID, &res.UserID, &res.MenuID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository MenuUserRepository) BrowseBy(column, value, operator string) (data []models.MenuUser, err error) {
	statement := selectMenuUser + ` from "menu_users" where ` + column + `` + operator + `$1`
	rows, err := repository.DB.Query(statement, value)
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

func (MenuUserRepository) Add(UserID, menuID string, tx *sql.Tx) (err error) {
	statement := `insert into "menu_users" ("user_id","menu_id") values($1,$2)`
	_, err = tx.Exec(statement, UserID, menuID)
	if err != nil {
		return err
	}

	return nil
}

func (MenuUserRepository) Delete(UserID string, tx *sql.Tx) (err error) {
	statement := `delete from "menu_users" where "user_id"=$1`
	_, err = tx.Exec(statement, UserID)
	if err != nil {
		return err
	}

	return nil
}
