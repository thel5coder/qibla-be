package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/str"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type MasterZakatRepository struct {
	DB *sql.DB
}

func NewMasterZakatRepository(DB *sql.DB) contracts.IMasterZakatRepository {
	return &MasterZakatRepository{DB: DB}
}

const(
	masterZakatSelectStatement =`select id,slug,type_zakat,name,description,amount,current_gold_price,gold_nishab,created_at,updated_at`
)

func (repository MasterZakatRepository) scanRows(rows *sql.Rows)(res models.MasterZakat,err error){
	err = rows.Scan(&res.ID, &res.Slug, &res.TypeZakat, &res.Name, &res.Description, &res.Amount, &res.CurrentGoldPrice, &res.GoldNishab,
		&res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res,err
	}

	return res,nil
}

func (repository MasterZakatRepository) scanRow(row *sql.Row)(res models.MasterZakat,err error){
	err = row.Scan(&res.ID, &res.Slug, &res.TypeZakat, &res.Name, &res.Description, &res.Amount, &res.CurrentGoldPrice, &res.GoldNishab,
		&res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return res,err
	}

	return res,nil
}

func (repository MasterZakatRepository) Browse(filters map[string]interface{}, order, sort string, limit, offset int) (data []models.MasterZakat, count int, err error) {
	filterStatement := ``
	if val,ok := filters["type_zakat"]; ok {
		filterStatement +=` and lower(cast("type_zakat" as varchar)) like '%`+val.(string)+`%'`
	}

	if val,ok := filters["name"]; ok {
		filterStatement +=` and lower("name") like '%`+val.(string)+`%'`
	}

	if val,ok := filters["description"]; ok {
		filterStatement +=` and lower("description") like '%`+val.(string)+`%'`
	}

	if val,ok := filters["updated_at"]; ok {
		filterStatement +=` and lower(cast("updated_at" as varchar)) like '%`+val.(string)+`%'`
	}

	statement := masterZakatSelectStatement+` from "master_zakats" where "deleted_at" is null`+filterStatement+` order by ` + order + ` ` + sort + ` limit $1 offset $2`
	fmt.Println(statement)
	rows, err := repository.DB.Query(statement,limit, offset)
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

	statement = `select count("id") from "master_zakats" where "deleted_at" is null`+filterStatement
	err = repository.DB.QueryRow(statement).Scan(&count)
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
		temp,err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

func (repository MasterZakatRepository) ReadBy(column, value string) (data models.MasterZakat, err error) {
	statement := `select * from "master_zakats" where ` + column + `=$1 and "deleted_at" is null`
	row := repository.DB.QueryRow(statement,value)
	data,err = repository.scanRow(row)

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
