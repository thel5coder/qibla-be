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

// DisbursementDetailRepository ...
type DisbursementDetailRepository struct {
	DB *sql.DB
}

// NewDisbursementDetailRepository ...
func NewDisbursementDetailRepository(DB *sql.DB) contracts.IDisbursementDetailRepository {
	return &DisbursementDetailRepository{DB: DB}
}

func (repository DisbursementDetailRepository) scanRows(rows *sql.Rows) (d models.DisbursementDetail, err error) {
	err = rows.Scan(
		&d.ID, &d.DisbursementID, &d.TransactionID, &d.Transaction.InvoiceNumber,
		&d.Transaction.PaymentMethodCode, &d.Transaction.PaymentStatus,
		&d.Transaction.DueDate, &d.Transaction.VaNumber, &d.Transaction.BankName,
	)

	return d, err
}

func (repository DisbursementDetailRepository) scanRow(row *sql.Row) (d models.DisbursementDetail, err error) {
	err = row.Scan(
		&d.ID, &d.DisbursementID, &d.TransactionID, &d.Transaction.InvoiceNumber,
		&d.Transaction.PaymentMethodCode, &d.Transaction.PaymentStatus,
		&d.Transaction.DueDate, &d.Transaction.VaNumber, &d.Transaction.BankName,
	)

	return d, err
}

// BrowseAll ...
func (repository DisbursementDetailRepository) BrowseAll(disbursementID string) (data []models.DisbursementDetail, err error) {
	statement := models.DisbursementDetailSelect + ` WHERE def."deleted_at" IS NULL
	AND def."disbursement_id" = $1`
	rows, err := repository.DB.Query(statement, disbursementID)
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

// Add ...
func (DisbursementDetailRepository) Add(input viewmodel.DisbursementDetailVm, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO "disbursement_details" ("disbursement_id", "transaction_id")
	VALUES ($1, $2) returning "id"`
	err = tx.QueryRow(statement,
		str.EmptyString(input.DisbursementID), str.EmptyString(input.TransactionID),
	).Scan(&res)

	return res, err
}

// Delete ...
func (DisbursementDetailRepository) Delete(disbursementID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `UPDATE "disbursement_details" SET "updated_at"=$1, "deleted_at"=$2
	WHERE "disbursement_id"=$3 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), disbursementID)

	return err
}
