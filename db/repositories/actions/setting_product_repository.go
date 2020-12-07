package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/str"
	"qibla-backend/usecase/viewmodel"
	"time"
)

// SettingProductRepository ...
type SettingProductRepository struct {
	DB *sql.DB
}

// NewSettingProductRepository ...
func NewSettingProductRepository(DB *sql.DB) contracts.ISettingProductRepository {
	return &SettingProductRepository{DB: DB}
}

func (repository SettingProductRepository) scanRows(rows *sql.Rows) (d models.SettingProduct, err error) {
	err = rows.Scan(
		&d.ID, &d.ProductID, &d.Price, &d.PriceUnit, &d.MaintenancePrice, &d.Discount,
		&d.DiscountType, &d.DiscountPeriodStart, &d.DiscountPeriodEnd, &d.Description,
		&d.Sessions, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt, &d.ProductName,
		&d.ProductType, &d.Features, &d.Periods,
	)

	return d, err
}

func (repository SettingProductRepository) scanRow(row *sql.Row) (d models.SettingProduct, err error) {
	err = row.Scan(
		&d.ID, &d.ProductID, &d.Price, &d.PriceUnit, &d.MaintenancePrice, &d.Discount,
		&d.DiscountType, &d.DiscountPeriodStart, &d.DiscountPeriodEnd, &d.Description,
		&d.Sessions, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt, &d.ProductName,
		&d.ProductType, &d.Features, &d.Periods,
	)

	return d, err
}

// Browse ...
func (repository SettingProductRepository) Browse(search, order, sort string, limit, offset int) (data []models.SettingProduct, count int, err error) {
	statement := models.SettingProductSelect + ` where sp."deleted_at" is null
		` + models.SettingProductGroup + ` order by sp.` + order + ` ` + sort + `
		limit $1 offset $2`
	rows, err := repository.DB.Query(statement, limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		d, err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, d)
	}

	statement = `select count(sp."id") from "setting_products" sp where sp."deleted_at" is null`
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

// BrowseBy ...
func (repository SettingProductRepository) BrowseBy(column, value, operator string) (data []models.SettingProduct, err error) {
	statement := models.SettingProductSelect + ` where ` + column + `` + operator + `$1
	and sp."deleted_at" is null ` + models.SettingProductGroup + ` order by sp."id" asc`
	rows, err := repository.DB.Query(statement, value)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		d, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}

	return data, err
}

// BrowseAll ...
func (repository SettingProductRepository) BrowseAll() (data []models.SettingProduct, err error) {
	statement := models.SettingProductSelect + ` where sp."deleted_at" is null ` + models.SettingProductGroup
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		d, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}

	return data, err
}

// ReadBy ...
func (repository SettingProductRepository) ReadBy(column, value string) (data models.SettingProduct, err error) {
	statement := models.SettingProductSelect + ` where ` + column + `=$1
	and sp."deleted_at" is null ` + models.SettingProductGroup
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, err
}

// Add ...
func (SettingProductRepository) Add(input viewmodel.SettingProductVm, tx *sql.Tx) (res string, err error) {
	statement := `insert into "setting_products" (
		"product_id","price","price_unit","maintenance_price","discount","discount_type",
		"discount_period_start","discount_period_end","description","sessions",
		"created_at","updated_at"
	) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning "id"`
	err = tx.QueryRow(
		statement, input.ProductID, input.Price, str.EmptyString(input.PriceUnit),
		str.EmptyInt(int(input.MaintenancePrice)), str.EmptyInt(int(input.Discount)),
		str.EmptyString(input.DiscountType), datetime.StrParseToTime(input.DiscountPeriodStart, "2006-01-02"),
		datetime.StrParseToTime(input.DiscountPeriodEnd, "2006-01-02"), input.Description,
		str.EmptyString(input.Sessions), datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

// Edit ...
func (SettingProductRepository) Edit(input viewmodel.SettingProductVm, tx *sql.Tx) (err error) {
	statement := `update "setting_products" set "product_id"=$1, "price"=$2,
		"price_unit"=$3, "maintenance_price"=$4, "discount"=$5, "discount_type"=$6,
		"discount_period_start"=$7, "discount_period_end"=$8, "description"=$9,
		"sessions"=$10, "updated_at"=$11 where "id"=$12 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		input.ProductID, input.Price, str.EmptyString(input.PriceUnit),
		str.EmptyInt(int(input.MaintenancePrice)), str.EmptyInt(int(input.Discount)),
		str.EmptyString(input.DiscountType), datetime.StrParseToTime(input.DiscountPeriodStart, "2006-01-02"),
		datetime.StrParseToTime(input.DiscountPeriodEnd, "2006-01-02"), input.Description,
		str.EmptyString(input.Sessions), datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	)

	return err
}

// Delete ...
func (SettingProductRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "setting_products" set "updated_at"=$1,"deleted_at"=$2
	where "id"=$3 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return err
}

// CountBy ...
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
