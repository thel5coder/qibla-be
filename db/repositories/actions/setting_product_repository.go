package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/helpers/str"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type SettingProductRepository struct {
	DB *sql.DB
}

func NewSettingProductRepository(DB *sql.DB) contracts.ISettingProductRepository {
	return &SettingProductRepository{DB: DB}
}

func (repository SettingProductRepository) Browse(search, order, sort string, limit, offset int) (data []models.SettingProduct, count int, err error) {
	statement := `select sp.*,mp."name" from "setting_products" sp 
                 inner join "master_products" mp on mp."id"=sp."product_id" and mp."deleted_at" is null
                order by sp.` + order + ` ` + sort + ` limit $1 offset $2`
	rows, err := repository.DB.Query(statement, limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.SettingProduct{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.ProductID,
			&dataTemp.Price,
			&dataTemp.PriceUnit,
			&dataTemp.MaintenancePrice,
			&dataTemp.Discount,
			&dataTemp.DiscountType,
			&dataTemp.DiscountPeriodStart,
			&dataTemp.DiscountPeriodEnd,
			&dataTemp.Description,
			&dataTemp.Sessions,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.ProductName,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count(sp."id") from "setting_products" sp 
                 inner join "master_products" mp on mp."id"=sp."product_id" and mp."deleted_at" is null`
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository SettingProductRepository) BrowseAll() (data []models.SettingProduct, err error) {
	statement := `select sp.*,mp."name" from "setting_products" sp 
                 inner join "master_products" mp on mp."id"=sp."product_id" and mp."deleted_at" is null`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.SettingProduct{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.ProductID,
			&dataTemp.Price,
			&dataTemp.PriceUnit,
			&dataTemp.MaintenancePrice,
			&dataTemp.Discount,
			&dataTemp.DiscountType,
			&dataTemp.DiscountPeriodStart,
			&dataTemp.DiscountPeriodEnd,
			&dataTemp.Description,
			&dataTemp.Sessions,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.ProductName,
		)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}

	return data, err
}

func (repository SettingProductRepository) ReadBy(column, value string) (data models.SettingProduct, err error) {
	statement := `select sp.*,mp."name" from "setting_products" sp 
                 inner join "master_products" mp on mp."id"=sp."product_id" and mp."deleted_at" is null
                 where ` + column + `=$1`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.ProductID,
		&data.Price,
		&data.PriceUnit,
		&data.MaintenancePrice,
		&data.Discount,
		&data.DiscountType,
		&data.DiscountPeriodStart,
		&data.DiscountPeriodEnd,
		&data.Description,
		&data.Sessions,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.ProductName,
	)

	return data, err
}

func (SettingProductRepository) Edit(input viewmodel.SettingProductVm, tx *sql.Tx) (err error) {
	statement := `update "setting_products" set "product_id"=$1, "price"=$2, "price_unit"=$3, "maintenance_price"=$4, "discount"=$5, "discount_type"=$6, "discount_period_start"=$7, "discount_period_end"=$8,
                  "description"=$9, "sessions"=$10, "updated_at"=$11 where "id"=$12`
	_, err = tx.Exec(
		statement,
		input.ProductID,
		input.Price,
		str.EmptyString(input.PriceUnit),
		str.EmptyInt(int(input.MaintenancePrice)),
		str.EmptyInt(int(input.Discount)),
		str.EmptyString(input.DiscountType),
		datetime.StrParseToTime(input.DiscountPeriodStart, time.RFC3339),
		datetime.StrParseToTime(input.DiscountPeriodEnd, time.RFC3339),
		input.Description,
		str.EmptyString(input.Sessions),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	)

	return err
}

func (SettingProductRepository) Add(input viewmodel.SettingProductVm, tx *sql.Tx) (res string, err error) {
	statement := `insert into "setting_products" 
                  ("product_id","price","price_unit","maintenance_price","discount","discount_type","discount_period_start","discount_period_end","description","sessions","created_at","updated_at")
                values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning "id"`
	err = tx.QueryRow(
		statement,
		input.ProductID,
		input.Price,
		str.EmptyString(input.PriceUnit),
		str.EmptyInt(int(input.MaintenancePrice)),
		str.EmptyInt(int(input.Discount)),
		str.EmptyString(input.DiscountType),
		datetime.StrParseToTime(input.DiscountPeriodStart, time.RFC3339),
		datetime.StrParseToTime(input.DiscountPeriodEnd, time.RFC3339),
		input.Description,
		str.EmptyString(input.Sessions),
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (SettingProductRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "setting_products" set "updated_at"=$1,"deleted_at"=$2 where "id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return err
}

func (repository SettingProductRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "setting_products" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "setting_products" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
