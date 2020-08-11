package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/helpers/str"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type MenuRepository struct{
	DB *sql.DB
}

func NewMenuRepository(DB *sql.DB) contracts.IMenuRepository{
	return &MenuRepository{DB: DB}
}

func (repository MenuRepository) Browse(parentID,search, order, sort string, limit, offset int) (data []models.Menu, count int, err error) {
	statement := ``
	var rows *sql.Rows
	if parentID != ""{
		statement = `select * from "menus" 
                     where "deleted_at" is null and "parent_id"=$1`
		rows,err = repository.DB.Query(statement,parentID)
	}else{
		statement = `select * from "menus" 
                     where ("menu_id" like $1 or lower("name") like $1 or lower("url") like $1) and "deleted_at" is null and "parent_id" is null
                     order by `+order+` `+sort+` limit $2 offset $3`
		rows,err = repository.DB.Query(statement,"%"+strings.ToLower(search)+"%",limit,offset)
	}

	if err != nil {
		return data,count,err
	}

	for rows.Next() {
		dataTemp := models.Menu{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.MenuID,
			&dataTemp.Name,
			&dataTemp.Url,
			&dataTemp.ParentID,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			)
		if err != nil {
			return data,count,err
		}

		data = append(data,dataTemp)
	}

	if parentID != ""{
		statement = `select count("id") from "menus" 
                     where ("menu_id" like $1 or lower("name") like $1 or lower("url") like $1) and "deleted_at" is null and "parent_id"=$2`
		err = repository.DB.QueryRow(statement,"%"+strings.ToLower(search)+"%",parentID).Scan(&count)
	}else{
		statement = `select count("id") from "menus" 
                     where ("menu_id" like $1 or lower("name") like $1 or lower("url") like $1) and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,"%"+strings.ToLower(search)+"%").Scan(&count)
	}

	if err != nil {
		return data,count,err
	}

	return data,count,err
}

func (repository MenuRepository) ReadBy(column, value string) (data models.Menu, err error) {
	statement := `select * from "menus" where `+column+`=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,value).Scan(
		&data.ID,
		&data.MenuID,
		&data.Name,
		&data.Url,
		&data.ParentID,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		)

	return data,err
}

func (repository MenuRepository) Edit(input viewmodel.MenuVm, tx *sql.Tx) (err error) {
	statement := `update "menus" set "name"=$1, "url"=$2, "is_active"=$3, "updated_at"=$4 where "id"=$5 and "deleted_at" is null returning "id"`
	_,err = tx.Exec(
		statement,
		input.Name,
		input.Url,
		input.IsActive,
		datetime.StrParseToTime(input.UpdatedAt,time.RFC3339),
		input.ID,
		)

	return err
}

func (repository MenuRepository) Add(input viewmodel.MenuVm, tx *sql.Tx) (res string, err error) {
	statement := `insert into "menus" ("menu_id","name","url","parent_id","is_active","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7) returning "id"`
	err = tx.QueryRow(
		statement,
		input.MenuID,
		input.Name,
		input.Url,
		str.EmptyString(input.ParentID),
		input.IsActive,
		datetime.StrParseToTime(input.CreatedAt,time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt,time.RFC3339),
		).Scan(&res)

	return res,err
}

func (repository MenuRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "menus" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3`
	_, err = tx.Exec(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID)

	return err
}

func (repository MenuRepository) DeleteChild(parentID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "menus" set "updated_at"=$1, "deleted_at"=$2 where "parent_id"=$3`
	_, err = tx.Exec(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),parentID)

	return err
}

func (repository MenuRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID != ""{
		statement := `select count("id") from "menus" where `+column+`=$1 and "deleted_at" is null and "id"<>$2`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}else{
		if value == "" {
			statement := `select count("id") from "menus" where `+column+` is null and "deleted_at" is null`
			err = repository.DB.QueryRow(statement).Scan(&res)
		}else{
			statement := `select count("id") from "menus" where `+column+`=$1 and "deleted_at" is null`
			err = repository.DB.QueryRow(statement,value).Scan(&res)
		}
	}

	return res,err
}

func (repository MenuRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "menus" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,ID).Scan(&res)

	return res,err
}

