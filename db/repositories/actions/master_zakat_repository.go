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

type MasterZakatRepository struct {
	DB *sql.DB
}

func NewMasterZakatRepository(DB *sql.DB) contracts.IMasterZakatRepository {
	return &MasterZakatRepository{DB: DB}
}

func (repository MasterZakatRepository) Browse(search, order, sort string, limit, offset int) (data []models.MasterZakat, count int, err error) {
	statement := `select * from "master_zakats" where (lower(cast("type_zakat" as varchar)) like $1 or lower("name") like $1 or lower("description") like $1 or cast("updated_at" as varchar) like $1) and "deleted_at" is null
                order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.MasterZakat{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Slug,
			&dataTemp.TypeZakat,
			&dataTemp.Name,
			&dataTemp.Description,
			&dataTemp.Amount,
			&dataTemp.CurrentGoldPrice,
			&dataTemp.GoldNishab,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, count, err
		}
		data = append(data, dataTemp)
	}

	statement = `select count("id") from "master_zakats" where (lower(cast("type_zakat" as varchar)) like $1 or lower("name") like $1 or lower("description") like $1 or cast("updated_at" as varchar) like $1) and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository MasterZakatRepository) BrowseAll() (data []models.MasterZakat, err error) {
	statement := `select * from "master_zakats" where "deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.MasterZakat{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Slug,
			&dataTemp.TypeZakat,
			&dataTemp.Name,
			&dataTemp.Description,
			&dataTemp.Amount,
			&dataTemp.CurrentGoldPrice,
			&dataTemp.GoldNishab,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository MasterZakatRepository) ReadBy(column, value string) (data models.MasterZakat, err error) {
	statement := `select * from "master_zakats" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.Slug,
		&data.TypeZakat,
		&data.Name,
		&data.Description,
		&data.Amount,
		&data.CurrentGoldPrice,
		&data.GoldNishab,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

func (repository MasterZakatRepository) Edit(input viewmodel.MasterZakatVm) (res string, err error) {
	statement := `update "master_zakats" set "slug"=$1, "type_zakat"=$2, "name"=$3, "description"=$4, "amount"=$5, "current_gold_price"=$6, "gold_nishab"=$7, "updated_at"=$8 where "id"=$9 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.Slug,
		input.TypeZakat,
		input.Name,
		input.Description,
		str.EmptyInt(int(input.Amount)),
		str.EmptyInt(int(input.CurrentGoldPrice)),
		str.EmptyInt(int(input.GoldNishab)),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	).Scan(&res)

	return res, err
}

func (repository MasterZakatRepository) Add(input viewmodel.MasterZakatVm) (res string, err error) {
	statement := `insert into "master_zakats" ("slug","type_zakat","name","description","amount","current_gold_price","gold_nishab","created_at","updated_at")
                 values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.Slug,
		input.TypeZakat,
		input.Name,
		input.Description,
		str.EmptyInt(int(input.Amount)),
		str.EmptyInt(int(input.CurrentGoldPrice)),
		str.EmptyInt(int(input.GoldNishab)),
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (repository MasterZakatRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "master_zakats" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository MasterZakatRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == ""{
		statement := `select count ("id") from "master_zakats" where `+column+`=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,value).Scan(&res)
	}else{
		statement := `select count ("id") from "master_zakats" where (`+column+`=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}

	return res,err
}
