package actions

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/interfacepkg"
	"qibla-backend/pkg/str"
	"qibla-backend/usecase/viewmodel"
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
		&d.ID, &d.ContactID, &d.TransactionID, &d.Total, &d.Status, &d.DisbursementType,
		&d.StartPeriod, &d.EndPeriod, &d.DisburseAt, &d.AccountNumber, &d.AccountName,
		&d.AccountBankName, &d.AccountBankCode, &d.OriginAccountNumber, &d.OriginAccountName,
		&d.OriginAccountBankName, &d.OriginAccountBankCode, &d.PaymentDetails, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.Transaction.InvoiceNumber, &d.Transaction.PaymentMethodCode,
		&d.Transaction.PaymentStatus, &d.Transaction.DueDate, &d.Transaction.VaNumber,
		&d.Transaction.BankName, &d.Contact.BranchName, &d.Contact.TravelAgentName,
		&d.Contact.AccountBankName, &d.Contact.Address, &d.Contact.PhoneNumber,
	)

	return d, err
}

func (repository DisbursementRepository) scanRow(row *sql.Row) (d models.Disbursement, err error) {
	err = row.Scan(
		&d.ID, &d.ContactID, &d.TransactionID, &d.Total, &d.Status, &d.DisbursementType,
		&d.StartPeriod, &d.EndPeriod, &d.DisburseAt, &d.AccountNumber, &d.AccountName,
		&d.AccountBankName, &d.AccountBankCode, &d.OriginAccountNumber, &d.OriginAccountName,
		&d.OriginAccountBankName, &d.OriginAccountBankCode, &d.PaymentDetails, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.Transaction.InvoiceNumber, &d.Transaction.PaymentMethodCode,
		&d.Transaction.PaymentStatus, &d.Transaction.DueDate, &d.Transaction.VaNumber,
		&d.Transaction.BankName, &d.Contact.BranchName, &d.Contact.TravelAgentName,
		&d.Contact.AccountBankName, &d.Contact.Address, &d.Contact.PhoneNumber,
	)

	return d, err
}

// Browse ...
func (repository DisbursementRepository) Browse(filters map[string]interface{}, order, sort string, limit, offset int) (data []models.Disbursement, count int, err error) {
	var conditionString string
	if val, ok := filters["contact_travel_agent_name"]; ok {
		conditionString += ` AND LOWER(c."travel_agent_name") LIKE '%` + strings.ToLower(val.(string)) + `%'`
	}
	if val, ok := filters["contact_branch_name"]; ok {
		conditionString += ` AND LOWER(c."branch_name") LIKE '%` + strings.ToLower(val.(string)) + `%'`
	}
	if val, ok := filters["total"]; ok {
		conditionString += ` AND def."total"::TEXT LIKE '%` + val.(string) + `%'`
	}

	if val, ok := filters["start_date"]; ok {
		conditionString += ` AND def."start_period" <= '` + val.(string) + `'`
	}

	if val,ok := filters["end_date"]; ok {
		conditionString += `AND def."end_period" >= '`+val.(string)+`'`
	}

	if val, ok := filters["contact_account_bank_name"]; ok {
		conditionString += ` AND LOWER(c."account_bank_name") LIKE '%` + strings.ToLower(val.(string)) + `%'`
	}
	if val, ok := filters["status"]; ok {
		conditionString += ` AND lower(cast(def."status" as varchar)) = '` + val.(string) + `'`
	}
	if val, ok := filters["disburse_at"]; ok {
		conditionString += ` AND cast(def."disburse_at" as varchar) LIKE '%` + val.(string) + `%'`
	}
	if val, ok := filters["origin_account_bank_name"]; ok {
		conditionString += ` AND LOWER(def."origin_account_bank_name") LIKE '%` + strings.ToLower(val.(string)) + `%'`
	}

	statement := models.DisbursementSelect + ` WHERE def."deleted_at" IS NULL ` + conditionString + `
		ORDER BY def.` + order + ` ` + sort + ` LIMIT $1 OFFSET $2`
	fmt.Println(statement)
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

	statement = `SELECT COUNT(def."id") FROM "disbursements" def
	LEFT JOIN "transactions" t ON t."id" = def."transaction_id"
	LEFT JOIN "contacts" c ON c."id" = def."contact_id"
	WHERE def."deleted_at" IS NULL ` + conditionString
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
func (repository DisbursementRepository) BrowseAll(status string) (data []models.Disbursement, err error) {
	conditionString := ``
	if status != "" {
		conditionString += ` AND def."status" = '` + status + `'`
	}

	statement := models.DisbursementSelect + ` WHERE def."deleted_at" IS NULL ` + conditionString
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
	statement := models.DisbursementSelect + ` WHERE def."deleted_at" IS NULL
	AND def."id" = $1
	ORDER BY def."created_at" DESC LIMIT 1`
	fmt.Println(statement)
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, err
}

// ReadByPaymentID ...
func (repository DisbursementRepository) ReadByPaymentID(paymentID int) (data models.Disbursement, err error) {
	statement := models.DisbursementSelect + ` WHERE def."deleted_at" IS NULL
	AND def."payment_details"::json ->> 'id' = $1 ORDER BY def."created_at" DESC LIMIT 1`
	row := repository.DB.QueryRow(statement, paymentID)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, err
}

// Add ...
func (DisbursementRepository) Add(input viewmodel.DisbursementVm, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO "disbursements" (
		"contact_id", "transaction_id", "total", "status", "disbursement_type", "start_period",
		"end_period", "disburse_at", "account_number", "account_name", "account_bank_name",
		"account_bank_code", "origin_account_number", "origin_account_name", "origin_account_bank_name",
		"origin_account_bank_code", "created_at","updated_at"
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) returning "id"`
	err = tx.QueryRow(statement,
		str.EmptyString(input.ContactID), str.EmptyString(input.TransactionID), input.Total, input.Status, input.DisbursementType,
		str.EmptyString(input.StartPeriod), str.EmptyString(input.EndPeriod),
		str.EmptyString(input.DisburseAt), input.AccountNumber, input.AccountName,
		input.AccountBankName, input.AccountBankCode, input.OriginAccountNumber, input.OriginAccountName,
		input.OriginAccountBankName, input.OriginAccountBankCode, datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

// Edit ...
func (DisbursementRepository) Edit(input viewmodel.DisbursementVm, tx *sql.Tx) (err error) {
	statement := `UPDATE "disbursements" set "contact_id" = $1, "transaction_id" = $2, "total" = $3,
	"status" = $4, "disbursement_type" = $5, "start_period" = $6, "end_period" = $7,
	"disburse_at = $8, "account_number" = $9, "account_name" = $10, "account_bank_name" = $11,
	"account_bank_code" = $12, "account_number" = $13, "account_name" = $14,
	"account_bank_name" = $15, "account_bank_code" = $16, "updated_at" = $17
	WHERE "id" = $18 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		str.EmptyString(input.ContactID), str.EmptyString(input.TransactionID), input.Total, input.Status, input.DisbursementType,
		str.EmptyString(input.StartPeriod), str.EmptyString(input.EndPeriod),
		str.EmptyString(input.DisburseAt), input.AccountNumber, input.AccountName,
		input.AccountBankName, input.AccountBankCode, input.OriginAccountNumber, input.OriginAccountName,
		input.OriginAccountBankName, input.OriginAccountBankCode, datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID,
	)

	return err
}

// EditPaymentDetails ...
func (DisbursementRepository) EditPaymentDetails(input viewmodel.DisbursementVm, tx *sql.Tx) (err error) {
	statement := `UPDATE "disbursements" set "payment_details" = $1, "status" = $2, "updated_at" = $3
	WHERE "id" = $4 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		interfacepkg.Marshall(input.PaymentDetails), input.Status,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID,
	)

	return err
}

// EditStatus ...
func (DisbursementRepository) EditStatus(input viewmodel.DisbursementVm, tx *sql.Tx) (err error) {
	statement := `UPDATE "disbursements" set "status" = $1, "updated_at" = $2
	WHERE "id" = $3 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		input.Status, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID,
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
