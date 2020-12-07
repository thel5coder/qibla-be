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

type FaqRepository struct{
	DB *sql.DB
}

func NewFaqRepository(DB *sql.DB) contracts.IFaqRepository{
	return &FaqRepository{DB: DB}
}

func (repository FaqRepository) Browse(search, order, sort string, limit, offset int) (data []models.Faq, count int, err error) {
	statement := `select f."id",f."faq_category_name",fi."id",fi."question",fi."answer",fi."created_at",fi."updated_at",fi."deleted_at" from "faqs" f
                 inner join "faq_lists" fi on fi."faq_id"=f."id"
                where (lower(fi."question") like $1 or lower(fi."answer") like $1) and fi."deleted_at" is null order by fi.`+order+` `+sort+` limit $2 offset $3`
	rows,err := repository.DB.Query(statement,"%"+strings.ToLower(search)+"%",limit,offset)
	if err != nil {
		return data,count,err
	}

	for rows.Next(){
		dataTemp := models.Faq{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.FaqCategoryName,
			&dataTemp.FaqListID,
			&dataTemp.Question,
			&dataTemp.Answer,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			)
		if err != nil {
			return data,count,err
		}

		data = append(data,dataTemp)
	}

	statement = `select count(fi."id") from "faqs" f
                 inner join "faq_lists" fi on fi."faq_id"=f."id"
                where (lower(fi."question") like $1 or lower(fi."answer") like $1) and fi."deleted_at" is null`
	err = repository.DB.QueryRow(statement,"%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data,count,err
	}

	return data,count,err
}

func (repository FaqRepository) ReadBy(column, value string) (data []models.Faq, err error) {
	statement := `select f."id",f."faq_category_name",fi."id",fi."question",fi."answer",fi."created_at",fi."updated_at",fi."deleted_at" from "faqs" f
                 inner join "faq_lists" fi on fi."faq_id"=f."id"
                where `+column+`=$1 and fi."deleted_at" is null`
	rows,err := repository.DB.Query(statement,value)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := models.Faq{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.FaqCategoryName,
			&dataTemp.FaqListID,
			&dataTemp.Question,
			&dataTemp.Answer,
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

func (FaqRepository) Edit(input viewmodel.FaqVm, tx *sql.Tx) (err error) {
	panic("implement me")
}

func (repository FaqRepository) Add(input viewmodel.FaqVm, webContentCategoryID string, tx *sql.Tx) (res string, err error) {
	statement := `insert into "faqs" ("web_content_category_id","faq_category_name","created_at","updated_at") values($1,$2,$3,$4) returning "id"`
	err = tx.QueryRow(statement,webContentCategoryID,input.FaqCategoryName,datetime.StrParseToTime(input.CreatedAt,time.RFC3339),datetime.StrParseToTime(input.UpdatedAt,time.RFC3339)).Scan(&res)

	return res,err
}

func (FaqRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repository FaqRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == ""{
		statement := `select count("id") from "faqs" where `+column+`=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,value).Scan(&res)
	}else{
		statement := `select count("id") from "faqs" where (`+column+`=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}

	return res,err
}

func (FaqRepository) CountByPk(ID string) (res int, err error) {
	panic("implement me")
}

