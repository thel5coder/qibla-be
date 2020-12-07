package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/enums"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type MasterProductRepository struct {
	DB *sql.DB
}

func NewMasterProductRepository(DB *sql.DB) contracts.IMasterProductRepository {
	return &MasterProductRepository{DB: DB}
}

func (repository MasterProductRepository) Browse(search, order, sort string, limit, offset int) (data []models.MasterProduct, count int, err error) {
	statement := `select * from "master_products" 
                  where (lower("name") like $1 or cast("updated_at" as varchar) like $1) and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.MasterProduct{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Slug,
			&dataTemp.Name,
			&dataTemp.SubscriptionType,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count("id") from "master_products" 
                  where (lower("name") like $1 or cast("updated_at" as varchar) like $1) and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository MasterProductRepository) BrowseAll() (data []models.MasterProduct,err error) {
	statement := `select * from "master_products" where "deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.MasterProduct{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Slug,
			&dataTemp.Name,
			&dataTemp.SubscriptionType,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}

	return data,err
}

func (repository MasterProductRepository) BrowseExtraProducts() (data []models.MasterProduct,err error) {
	statement := `select * from "master_products" where "subscription_type"<>$1 and "deleted_at" is null`
	rows, err := repository.DB.Query(statement,enums.KeySubscriptionEnum1)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.MasterProduct{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Slug,
			&dataTemp.Name,
			&dataTemp.SubscriptionType,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}

	return data,err
}

func (repository MasterProductRepository) ReadBy(column, value string) (data models.MasterProduct, err error) {
	statement := `select * from "master_products" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.Slug,
		&data.Name,
		&data.SubscriptionType,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

func (repository MasterProductRepository) Edit(input viewmodel.MasterProductVm) (res string, err error) {
	statement := `update "master_products" set "name"=$1, "subscription_type"=$2, "slug"=$3, "updated_at"=$4 where "id"=$5 returning "id"`
	err = repository.DB.QueryRow(statement, input.Name, input.SubscriptionType, input.Slug, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID).Scan(&res)

	return res, err
}

func (repository MasterProductRepository) Add(input viewmodel.MasterProductVm) (res string, err error) {
	fmt.Println(input.UpdatedAt)
	fmt.Println(input.UpdatedAt, time.RFC3339)
	statement := `insert into "master_products" ("name","subscription_type","slug","created_at","updated_at") values($1,$2,$3,$4,$5) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.Name,
		input.SubscriptionType,
		input.Slug,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (repository MasterProductRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "master_products" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository MasterProductRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "master_products" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "master_products" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
