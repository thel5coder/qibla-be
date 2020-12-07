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

type TermConditionRepository struct {
	DB *sql.DB
}

func NewTermConditionRepository(DB *sql.DB) contracts.ITermConditionRepository {
	return &TermConditionRepository{DB: DB}
}

func (repository TermConditionRepository) Browse(search, order, sort string, limit, offset int) (data []models.TermConditions, count int, err error) {
	statement := `select * from "term_conditions" where (lower("term_name") like $1 or lower("description") like $1) and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.TermConditions{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.TermName,
			&dataTemp.TermType,
			&dataTemp.Description,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count("id") from "term_conditions" where (lower("term_name") like $1 or lower("description") like $1) and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository TermConditionRepository) ReadBy(column, value string) (data models.TermConditions, err error) {
	statement := `select * from "term_conditions" where `+column+`=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,value).Scan(
		&data.ID,
		&data.TermName,
		&data.TermType,
		&data.Description,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		)

	return data,err
}

func (repository TermConditionRepository) Edit(input viewmodel.TermConditionVm) (res string, err error) {
	statement := `update "term_conditions" set "term_name"=$1, "term_type"=$2, "description"=$3, "updated_at"=$4 where "id"=$5 returning "id"`
	err = repository.DB.QueryRow(statement,input.TermName,input.TermType,input.Description,datetime.StrParseToTime(input.UpdatedAt,time.RFC3339), input.ID).Scan(&res)

	return res,err
}

func (repository TermConditionRepository) Add(input viewmodel.TermConditionVm) (res string, err error) {
	statement := `insert into "term_conditions" ("term_name","term_type","description","created_at","updated_at") values($1,$2,$3,$4,$5) returning "id"`
	err = repository.DB.QueryRow(statement,input.TermName,input.TermType,input.Description,datetime.StrParseToTime(input.CreatedAt,time.RFC3339),datetime.StrParseToTime(input.UpdatedAt,time.RFC3339)).Scan(&res)

	return res,err
}

func (repository TermConditionRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "term_conditions" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID).Scan(&res)

	return res,err
}

func (repository TermConditionRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == ""{
		statement := `select count("id") from "term_conditions" where `+column+`=$1`
		err = repository.DB.QueryRow(statement,value).Scan(&res)
	}else{
		statement := `select count("id") from "term_conditions" where `+column+`=$1 and "id"<>$2`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}

	return res,err
}

func (repository TermConditionRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "term_conditions" where "id"=$1`
	err = repository.DB.QueryRow(statement,ID).Scan(&res)

	return res,err
}
