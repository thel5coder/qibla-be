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

// DisbursementRepository ...
type DisbursementRepository struct {
	DB *sql.DB
}

// NewDisbursementRepository ...
func NewDisbursementRepository(DB *sql.DB) contracts.IDisbursementRepository {
	return &DisbursementRepository{DB: DB}
}

func (repository DisbursementRepository) scanRows(rows *sql.Rows) (d models.Disbursement, err error) {
	err = rows.Scan(
		&d.ID, &d.TransactionID, &d.Total, &d.Status, &d.DisbursementType,
		&d.StartPeriod, &d.EndPeriod, &d.DisburseAt, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.Transaction.InvoiceNumber, &d.Transaction.PaymentMethodCode,
		&d.Transaction.PaymentStatus, &d.Transaction.DueDate, &d.Transaction.VaNumber,
		&d.Transaction.BankName,
	)

	return d, err
}

func (repository DisbursementRepository) scanRow(row *sql.Row) (d models.Disbursement, err error) {
	err = row.Scan(
		&d.ID, &d.TransactionID, &d.Total, &d.Status, &d.DisbursementType,
		&d.StartPeriod, &d.EndPeriod, &d.DisburseAt, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.Transaction.InvoiceNumber, &d.Transaction.PaymentMethodCode,
		&d.Transaction.PaymentStatus, &d.Transaction.DueDate, &d.Transaction.VaNumber,
		&d.Transaction.BankName,
	)

	return d, err
}

// Browse ...
func (repository DisbursementRepository) Browse(search, order, sort string, limit, offset int) (data []models.Disbursement, count int, err error) {
	statement := models.DisbursementSelect + ` WHERE def."deleted_at" IS NULL
		ORDER BY def.` + order + ` ` + sort + ` LIMIT $1 OFFSET $2`
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

	statement = `SELECT COUNT(def."id") FROM "disbursements" uz WHERE def."deleted_at" IS NULL`
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

// BrowseBy ...
func (repository DisbursementRepository) BrowseBy(column, value, operator string) (data []models.Disbursement, err error) {
	statement := models.DisbursementSelect + ` WHERE ` + column + `` + operator + `$1
	AND def."deleted_at" IS NULL ORDER BY def."id" ASC`
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
func (repository DisbursementRepository) BrowseAll() (data []models.Disbursement, err error) {
	statement := models.DisbursementSelect + ` WHERE def."deleted_at" IS NULL`
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
func (repository DisbursementRepository) ReadBy(column, value string) (data models.Disbursement, err error) {
	statement := models.DisbursementSelect + ` WHERE ` + column + `=$1
	AND def."deleted_at" IS NULL`
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, err
}

// Add ...
func (DisbursementRepository) Add(input viewmodel.DisbursementVm, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO "disbursements" (
		"transaction_id", "total", "status", "disbursement_type", "start_period",
		"end_period", "disburse_at", "created_at","updated_at"
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $8) returning "id"`
	err = tx.QueryRow(statement,
		str.EmptyString(input.TransactionID), input.Total, input.Status, input.DisbursementType,
		str.EmptyString(input.StartPeriod), str.EmptyString(input.EndPeriod),
		str.EmptyString(input.DisburseAt), datetime.StrParseToTime(input.CreatedAt, time.RFC3339), datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

// Edit ...
func (DisbursementRepository) Edit(input viewmodel.DisbursementVm, tx *sql.Tx) (err error) {
	statement := `UPDATE "disbursements" set "transaction_id" = $1, "total" = $2,
	"status" = $3, "disbursement_type" = $4, "start_period" = $5, "end_period" = $6,
	"disburse_at = $7, "updated_at"=$8 WHERE "id"=$9 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		str.EmptyString(input.TransactionID), input.Total, input.Status, input.DisbursementType,
		str.EmptyString(input.StartPeriod), str.EmptyString(input.EndPeriod),
		str.EmptyString(input.DisburseAt), datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID,
	)

	return err
}

// Delete ...
func (DisbursementRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `UPDATE "disbursements" SET "updated_at"=$1,"deleted_at"=$2
	WHERE "id"=$3 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return err
}

// CountBy ...
func (repository DisbursementRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `SELECT count("id") FROM "disbursements" WHERE ` + column + `=$1 AND "deleted_at" IS NULL`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `SELECT count("id") FROM "disbursements" WHERE (` + column + `=$1 AND "deleted_at" IS NULL) AND "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
