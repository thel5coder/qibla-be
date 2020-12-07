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

type GlobalInfoCategoryRepository struct {
	DB *sql.DB
}

func NewGlobalInfoCategoryRepository(DB *sql.DB) contracts.IGlobalInfoCategoryRepository {
	return &GlobalInfoCategoryRepository{DB: DB}
}

func (repository GlobalInfoCategoryRepository) Browse(search, order, sort string, limit, offset int) (data []models.GlobalInfoCategory, count int, err error) {
	statement := `select * from "global_info_categories" where lower("name") like $1 and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.GlobalInfoCategory{}

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

	statement = `select count("id") from "global_info_categories" where lower("name") like $1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository GlobalInfoCategoryRepository) ReadBy(column, value string) (data models.GlobalInfoCategory, err error) {
	statement := `select * from "global_info_categories" where ` + column + `=$1 and "deleted_at" is null`
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

func (repository GlobalInfoCategoryRepository) Edit(input viewmodel.GlobalInfoCategoryVm) (res string, err error) {
	statement := `update "global_info_categories" set "name"=$1, "updated_at"=$2, "slug"=$3 where "id"=$4 returning "id"`
	err = repository.DB.QueryRow(statement, input.Name, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.Slug, input.ID).Scan(&res)

	return res, err
}

func (repository GlobalInfoCategoryRepository) Add(input viewmodel.GlobalInfoCategoryVm) (res string, err error) {
	statement := `insert into "global_info_categories" ("name","slug","created_at","updated_at") values($1,$2,$3,$4) returning "id"`
	err = repository.DB.QueryRow(statement, input.Name, input.Slug, datetime.StrParseToTime(input.CreatedAt, time.RFC3339), datetime.StrParseToTime(input.UpdatedAt, time.RFC3339)).Scan(&res)

	return res, err
}

func (repository GlobalInfoCategoryRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "global_info_categories" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository GlobalInfoCategoryRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID != "" {
		statement := `select count("id") from "global_info_categories" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	} else {
		statement := `select count("id") from "global_info_categories" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	}

	return res, err
}
