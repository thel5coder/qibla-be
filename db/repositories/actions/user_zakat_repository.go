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

// UserZakatRepository ...
type UserZakatRepository struct {
	DB *sql.DB
}

// NewUserZakatRepository ...
func NewUserZakatRepository(DB *sql.DB) contracts.IUserZakatRepository {
	return &UserZakatRepository{DB: DB}
}

func (repository UserZakatRepository) scanRows(rows *sql.Rows) (d models.UserZakat, err error) {
	err = rows.Scan(
		&d.ID, &d.UserID, &d.TransactionID, &d.ContactID, &d.MasterZakatID, &d.TypeZakat,
		&d.CurrentGoldPrice, &d.GoldNishab, &d.Wealth, &d.Total, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.User.Email, &d.User.Name, &d.Transaction.InvoiceNumber, &d.Contact.BranchName,
		&d.Contact.TravelAgentName,
	)

	return d, err
}

func (repository UserZakatRepository) scanRow(row *sql.Row) (d models.UserZakat, err error) {
	err = row.Scan(
		&d.ID, &d.UserID, &d.TransactionID, &d.ContactID, &d.MasterZakatID, &d.TypeZakat,
		&d.CurrentGoldPrice, &d.GoldNishab, &d.Wealth, &d.Total, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.User.Email, &d.User.Name, &d.Transaction.InvoiceNumber, &d.Contact.BranchName,
		&d.Contact.TravelAgentName,
	)

	return d, err
}

// Browse ...
func (repository UserZakatRepository) Browse(search, order, sort string, limit, offset int) (data []models.UserZakat, count int, err error) {
	statement := models.UserZakatSelect + ` WHERE uz."deleted_at" IS NULL
		ORDER BY uz.` + order + ` ` + sort + ` LIMIT $1 OFFSET $2`
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

	statement = `SELECT COUNT(uz."id") FROM "user_zakats" uz WHERE uz."deleted_at" IS NULL`
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

// BrowseBy ...
func (repository UserZakatRepository) BrowseBy(column, value, operator string) (data []models.UserZakat, err error) {
	statement := models.UserZakatSelect + ` WHERE ` + column + `` + operator + `$1
	AND uz."deleted_at" IS NULL ORDER BY sp."id" ASC`
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
func (repository UserZakatRepository) BrowseAll() (data []models.UserZakat, err error) {
	statement := models.UserZakatSelect + ` WHERE uz."deleted_at" IS NULL`
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
func (repository UserZakatRepository) ReadBy(column, value string) (data models.UserZakat, err error) {
	statement := models.UserZakatSelect + ` WHERE ` + column + `=$1
	AND uz."deleted_at" IS NULL`
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, err
}

// Add ...
func (UserZakatRepository) Add(input viewmodel.UserZakatVm, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO "user_zakats" (
		"user_id","transaction_id","contact_id","master_zakat_id","type_zakat","current_gold_price",
		"gold_nishab","wealth","total", "created_at","updated_at"
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) returning "id"`
	err = tx.QueryRow(statement,
		str.EmptyString(input.UserID), str.EmptyString(input.TransactionID),
		str.EmptyString(input.ContactID), str.EmptyString(input.MasterZakatID), input.TypeZakat,
		input.CurrentGoldPrice, input.GoldNishab, input.Wealth, input.Total,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339), datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

// Edit ...
func (UserZakatRepository) Edit(input viewmodel.UserZakatVm, tx *sql.Tx) (err error) {
	statement := `UPDATE "user_zakats" set "user_id"=$1,"transaction_id"=$2,"contact_id"=$3,
		"master_zakat_id"=$4,"type_zakat"=$5,"current_gold_price"=$6, "gold_nishab"=$7,"wealth"=$8,
		"total"=$9, "updated_at"=$10 WHERE "id"=$11 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		str.EmptyString(input.UserID), str.EmptyString(input.TransactionID),
		str.EmptyString(input.ContactID), str.EmptyString(input.MasterZakatID), input.TypeZakat,
		input.CurrentGoldPrice, input.GoldNishab, input.Wealth, input.Total,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339), datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	)

	return err
}

// Delete ...
func (UserZakatRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `UPDATE "user_zakats" SET "updated_at"=$1,"deleted_at"=$2
	WHERE "id"=$3 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return err
}

// CountBy ...
func (repository UserZakatRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `SELECT count("id") FROM "user_zakats" WHERE ` + column + `=$1 AND "deleted_at" IS NULL`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `SELECT count("id") FROM "user_zakats" WHERE (` + column + `=$1 AND "deleted_at" IS NULL) AND "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
