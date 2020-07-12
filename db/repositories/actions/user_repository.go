package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) contracts.IUserRepository {
	return UserRepository{DB: DB}
}

func (repository UserRepository) Browse(search, order, sort string, limit, offset int) (data []models.User, count int, err error) {
	statement := `select "users".*,r.* from "users"
                 inner join "roles" r on r."id"="users"."role_id"
                 where (lower("users"."username") like $1 or lower("users"."email") like $1 or lower(r."name") like $1) and "users"."deleted_at" is null order by "users".` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data,count,err
	}

	for rows.Next() {
		dataTemp := models.User{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.UserName,
			&dataTemp.Email,
			&dataTemp.Password,
			&dataTemp.IsActive,
			&dataTemp.RoleID,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.RoleModel.ID,
			&dataTemp.RoleModel.Name,
			&dataTemp.RoleModel.CreatedAt,
			&dataTemp.RoleModel.UpdatedAt,
			&dataTemp.RoleModel.DeletedAt,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count("users"."id") from "users"
                 inner join "roles" r on r."id"=users."role_id"
                 where (lower(users."username") like $1 or lower(users."email") like $1 or lower(r."name") like $1) and "users"."deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository UserRepository) ReadBy(column, value string) (data models.User, err error) {
	statement := `select users.*,r."id",r."name" from "users"
                 inner join "roles" r on r."id"=users."role_id"
                 where "users".` + column + `=$1 and "users"."deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.UserName,
		&data.Email,
		&data.Password,
		&data.IsActive,
		&data.RoleID,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.RoleModel.ID,
		&data.RoleModel.Name,
	)

	return data, err
}

func (repository UserRepository) Edit(input viewmodel.UserVm, password string, tx *sql.Tx) (err error) {
	if password != "" {
		statement := `update "users" set "username"=$1, "email"=$2, "password"=$3, "is_active"=$4, "role_id"=$5, "updated_at"=$6 where "id"=$7`
		_, err = tx.Exec(statement, input.UserName, input.Email, password, input.IsActive, input.Role.ID, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID)
	} else {
		statement := `update "users" set "username"=$1, "email"=$2, "is_active"=$3, "role_id"=$4, "updated_at"=$5 where "id"=$6`
		_, err = tx.Exec(statement, input.UserName, input.Email, input.IsActive, input.Role.ID, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID)
	}

	return err
}

func (repository UserRepository) Add(input viewmodel.UserVm, password string, tx *sql.Tx) (res string, err error) {
	statement := `insert into "users" ("username","email","password","role_id","is_active","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7) returning "id"`
	err = tx.QueryRow(
		statement,
		input.UserName,
		input.Email,
		password,
		input.Role.ID,
		input.IsActive,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (repository UserRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "users" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3`
	_,err = tx.Exec(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID)

	return err
}

func (repository UserRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "users" where `+column+`=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,value).Scan(&res)
	}else{
		statement := `select count("id") from "users" where (`+column+`=$1 and "id"<>$2) and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}

	return res,err
}

func (repository UserRepository) CountByPk(ID string) (res int,err error){
	statement := `select count("id") from "users" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,ID).Scan(&res)

	return res,err
}
