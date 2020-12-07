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

type WebComprofCategoryRepository struct{
	DB *sql.DB
}

func NewWebComprofCategoryRepository(DB *sql.DB) contracts.IWebComprofCategoryRepository{
	return &WebComprofCategoryRepository{DB: DB}
}

func (repository WebComprofCategoryRepository) Browse(search, order, sort string, limit, offset int) (data []models.WebComprofCategory, count int, err error) {
	statement := `select * from "web_comprof_categories" where lower("name") like $1 and "deleted_at" is null order by `+order+` `+sort+` limit $2 offset $3`
	rows,err := repository.DB.Query(statement,"%"+strings.ToLower(search)+"%",limit,offset)
	if err != nil {
		return data,count,err
	}

	for rows.Next(){
		dataTemp := models.WebComprofCategory{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Slug,
			&dataTemp.Name,
			&dataTemp.CategoryType,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			)
		if err != nil {
			return data,count,err
		}

		data = append(data,dataTemp)
	}

	statement = `select count("id") from "web_comprof_categories" where lower("name") like $1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,"%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data,count,err
	}

	return data,count,err
}

func (repository WebComprofCategoryRepository) ReadBy(column, value string) (data models.WebComprofCategory, err error) {
	statement := `select * from "web_comprof_categories" where `+column+`=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,value).Scan(
		&data.ID,
		&data.Slug,
		&data.Name,
		&data.CategoryType,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		)

	return data,err
}

func (repository WebComprofCategoryRepository) Edit(input viewmodel.WebComprofCategoryVm) (res string, err error) {
	statement := `update "web_comprof_categories" set "slug"=$1,"name"=$2, "category_type"=$3, "updated_at"=$4 where "id"=$5 returning "id"`
	err = repository.DB.QueryRow(statement,input.Slug,input.Name,input.CategoryType,datetime.StrParseToTime(input.UpdatedAt,time.RFC3339),input.ID).Scan(&res)

	return res,err
}

func (repository WebComprofCategoryRepository) Add(input viewmodel.WebComprofCategoryVm) (res string, err error) {
	statement := `insert into "web_comprof_categories" ("slug","name","category_type","created_at","updated_at") values($1,$2,$3,$4,$5) returning "id"`
	err = repository.DB.QueryRow(statement,input.Slug,input.Name,input.CategoryType,datetime.StrParseToTime(input.CreatedAt,time.RFC3339),datetime.StrParseToTime(input.UpdatedAt,time.RFC3339)).Scan(&res)

	return res,err
}

func (repository WebComprofCategoryRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "web_comprof_categories" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID).Scan(&res)

	return res,err
}

func (repository WebComprofCategoryRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID != ""{
		statement := `select count("id") from "web_comprof_categories" where (`+column+`=$1 and "deleted_at" is null) and "id"<> $2`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}else{
		statement := `select count("id") from "web_comprof_categories" where `+column+`=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,value).Scan(&res)
	}

	return res,err
}

