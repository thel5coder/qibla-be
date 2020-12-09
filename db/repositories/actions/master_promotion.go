package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helper/master_promotion_helper"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type MasterPromotionRepository struct {
	DB *sql.DB
}

func NewMasterPromotionRepository(DB *sql.DB) contracts.IMasterPromotionRepository {
	return &MasterPromotionRepository{DB: DB}
}

const(
	masterPromotionSelect = `select id,slug,name,is_active,created_at,updated_at`
)

func(repository MasterPromotionRepository) scanRows(rows *sql.Rows) (res models.MasterPromotion,err error){
	err = rows.Scan(&res.ID,&res.Slug,&res.Name,&res.IsActive,&res.CreatedAt,&res.UpdatedAt)
	if err != nil{
		return res,err
	}

	return res,nil
}

func(repository MasterPromotionRepository) scanRow(row *sql.Row) (res models.MasterPromotion,err error){
	err = row.Scan(&res.ID,&res.Slug,&res.Name,&res.IsActive,&res.CreatedAt,&res.UpdatedAt)
	if err != nil{
		return res,err
	}

	return res,nil
}

func (repository MasterPromotionRepository) Browse(filters map[string]interface{}, order, sort string, limit, offset int) (data []models.MasterPromotion, count int, err error) {
	filterStatement := master_promotion_helper.ParseFilterStatement(filters)

	statement := masterPromotionSelect+` from "master_promotions" where "deleted_at" is null `+filterStatement+` order by ` + order + ` ` + sort + ` limit $1 offset $2`
	rows, err := repository.DB.Query(statement, limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		temp,err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}

		data = append(data, temp)
	}

	statement = `select count("id") from "master_promotions" where "deleted_at" is null`+filterStatement
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository MasterPromotionRepository) BrowseAll() (data []models.MasterPromotion,err error){
	statement := masterPromotionSelect+` from master_promotions`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		temp,err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}

		data = append(data, temp)
	}

	return data,err
}

func (repository MasterPromotionRepository) ReadBy(column, value string) (data models.MasterPromotion, err error) {
	statement := masterPromotionSelect+` from "master_promotions" where ` + column + `=$1 and "deleted_at" is null`
	row := repository.DB.QueryRow(statement,value)
	data,err = repository.scanRow(row)

	return data, err
}

func (repository MasterPromotionRepository) Edit(input viewmodel.MasterPromotionVm) (res string, err error) {
	statement := `update "master_promotions" set name=$1, "slug"=$2, "updated_at"=$3, "is_active"=$4 where "id"=$5 returning "id"`
	err = repository.DB.QueryRow(statement, input.Name, input.Slug, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.IsActive, input.ID).Scan(&res)

	return res, err
}

func (repository MasterPromotionRepository) Add(input viewmodel.MasterPromotionVm) (res string, err error) {
	statement := `insert into "master_promotions" (name,"slug","created_at","updated_at","is_active") values($1,$2,$3,$4,$5) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.Name,
		input.Slug,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.IsActive,
	).Scan(&res)

	return res, err
}

func (repository MasterPromotionRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "master_promotions" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository MasterPromotionRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "master_promotions" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "master_promotions" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
