package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type RoleRepository struct {
	DB *sql.DB
}

func NewRoleRepository(DB *sql.DB) contracts.IRoleRepository {
	return &RoleRepository{DB: DB}
}

func (repository RoleRepository) Browse(search, order, sort string, limit, offset int) (data []models.Role, count int, err error) {
	statement := `select * from "roles" where lower("name") like $1 and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.Role{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Name,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.Slug,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count("id") from "roles" where lower("name") like $1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository RoleRepository) ReadBy(column, value string) (data models.Role, err error) {
	statement := `select * from "roles"
                 where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.Name,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.Slug,
	)

	return data, err
}

func (repository RoleRepository) Edit(input viewmodel.RoleVm) (res string, err error) {
	statement := `update "roles" set "name"=$1, "slug"=$2, "updated_at"=$3 where "id"=$4 returning "id"`
	err = repository.DB.QueryRow(statement, input.Name, input.Slug, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID).Scan(&res)

	return res, err
}

func (repository RoleRepository) Add(input viewmodel.RoleVm) (res string, err error) {
	statement := `insert into "roles" ("name","slug","created_at","updated_at") values($1,$2,$3,$4) returning "id"`
	err = repository.DB.QueryRow(statement, input.Name,input.Slug, datetime.StrParseToTime(input.CreatedAt, time.RFC3339), datetime.StrParseToTime(input.UpdatedAt, time.RFC3339)).Scan(&res)

	return res, err
}

func (repository RoleRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "roles" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository RoleRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "roles" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "roles" where (` + column + `=$1 and "id"<>$2)and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}

func (repository RoleRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "roles" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)

	return res, err
}
