package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/helpers/enums"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(DB *sql.DB) contracts.ITransactionRepository {
	return &TransactionRepository{DB: DB}
}

const transactionSelect = `select t."id",t."user_id",t."invoice_number",t."trx_id",t."due_date",t."due_date_period",t."payment_status",
                         t."payment_method_code",t."va_number",t."bank_name",t."direction",t."transaction_type",t."paid_date",
                         t."transaction_date",t."updated_at",t."total",t."fee_qibla",t."is_disburse",t."is_disburse_allowed"`
const joinQuery = `left join "transaction_details" td on td."transaction_id"=t."id"`
const groupBy = `group by t."id"`

func (TransactionRepository) Browse(search, order, sort string, limit, offset int) (data []models.Transaction, count int, err error) {
	panic("implement me")
}

func (repository TransactionRepository) BrowseAllZakatDisbursement(contactID string) (data []models.Transaction, err error) {
	statement := `SELECT def."id", def."user_id", def."invoice_number", def."trx_id",
	def."due_date", def."due_date_period", def."payment_status", def."payment_method_code",
	def."va_number", def."bank_name", def."direction", def."transaction_type", def."paid_date",
	def."transaction_date", def."updated_at", def."total", def."fee_qibla", def."is_disburse",
	def."is_disburse_allowed"
	FROM "transactions" def
	JOIN "user_zakats" uz ON uz."transaction_id" = def."id"
	WHERE def."deleted_at" IS NULL AND def."transaction_type" = $1 AND def."payment_status" = $2
	AND def."is_disburse_allowed" = $3 AND def."is_disburse" = $4 AND uz."contact_id" = $5
	GROUP BY def."id"`
	rows, err := repository.DB.Query(statement,
		enums.KeyTransactionType1, enums.KeyPaymentStatus3, true, false, contactID,
	)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		d := models.Transaction{}
		err = rows.Scan(
			&d.ID, &d.UserID, &d.InvoiceNumber, &d.TrxID, &d.DueDate,
			&d.DueDatePeriod, &d.PaymentStatus, &d.PaymentMethodCode, &d.VaNumber,
			&d.BankName, &d.Direction, &d.TransactionType, &d.PaidDate,
			&d.TransactionDate, &d.UpdatedAt, &d.Total, &d.FeeQibla,
			&d.IsDisburse, &d.IsDisburseAllowed,
		)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}

	return data, err
}

func (repository TransactionRepository) ReadBy(column, value, operator string) (data models.Transaction, err error) {
	statement := transactionSelect + ` from "transactions" t ` + joinQuery + ` where ` + column + `` + operator + `$1 ` + groupBy
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.UserID,
		&data.InvoiceNumber,
		&data.TrxID,
		&data.DueDate,
		&data.DueDatePeriod,
		&data.PaymentStatus,
		&data.PaymentMethodCode,
		&data.VaNumber,
		&data.BankName,
		&data.Direction,
		&data.TransactionType,
		&data.PaidDate,
		&data.TransactionDate,
		&data.UpdatedAt,
		&data.Total,
	)

	return data, err
}

func (TransactionRepository) Add(input viewmodel.TransactionVm, tx *sql.Tx) (res string, err error) {
	statement := `insert into "transactions" ("user_id","invoice_number","trx_id","due_date","due_date_period","payment_status","payment_method_code","va_number",
                  "bank_name","direction","transaction_type","transaction_date","updated_at","total","fee_qibla","is_disburse","is_disburse_allowed","invoice_status")
                  values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18) returning "id"`
	err = tx.QueryRow(
		statement,
		input.UserID,
		input.InvoiceNumber,
		input.TrxID,
		datetime.StrParseToTime(input.DueDate, "2006-01-02"),
		input.DueDatePeriod,
		input.PaymentStatus,
		input.PaymentMethodCode,
		input.VaNumber,
		input.BankName,
		input.Direction,
		input.TransactionType,
		datetime.StrParseToTime(input.TransactionDate, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.Total,
		input.FeeQibla,
		input.IsDisburse,
		input.IsDisburseAllowed,
		input.InvoiceStatus,
	).Scan(&res)

	return res, err
}

func (repository TransactionRepository) EditDueDate(ID, dueDate, updatedAt string, dueDatePeriod int) (res string, err error) {
	statement := `update "transactions" set "due_date"=$1, "due_date_period"=$2, "updated_at"=$3 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		datetime.StrParseToTime(dueDate, "2006-01-02"),
		dueDatePeriod,
		datetime.StrParseToTime(updatedAt, time.RFC3339),
		ID,
	).Scan(&res)

	return res, err
}

func (repository TransactionRepository) EditStatus(ID, paymentStatus, paidDate, updatedAt string, tx *sql.Tx) (err error) {
	statement := `update "transactions" set "payment_status"=$1, "paid_date"=$2, "updated_at"=$3 where "id"=$4 returning "id"`
	_, err = tx.Exec(
		statement,
		paymentStatus,
		datetime.StrParseToTime(paidDate, "2006-01-02 15:04:05"),
		datetime.StrParseToTime(updatedAt, time.RFC3339),
		ID,
	)

	return err
}

func (repository TransactionRepository) EditTrxID(ID, trxID string, updatedAt string, tx *sql.Tx) (err error) {
	statement := `update "transactions" set "trx_id"=$1, "updated_at"=$2 where "id"=$3 returning "id"`
	_, err = tx.Exec(
		statement,
		trxID,
		datetime.StrParseToTime(updatedAt, time.RFC3339),
		ID,
	)

	return err
}

func (repository TransactionRepository) EditIsDisburse(ID, updatedAt string, tx *sql.Tx) (err error) {
	statement := `update "transactions" set "is_disburse" = $1, "updated_at" = $2 where "id" = $3 returning "id"`
	_, err = tx.Exec(statement,
		true, datetime.StrParseToTime(updatedAt, time.RFC3339), ID,
	)

	return err
}

func (repository TransactionRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "transactions" where ` + column + `=$1`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "transactions" where ` + column + `=$1 and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}

func (repository TransactionRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "transactions" set "updated_at"=$1, "deleted_at"=$2 where "id"=$1`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository TransactionRepository) GetInvoiceCount(month int) (res int, err error) {
	statement := `select count("id") from "transactions" where EXTRACT(month from "transaction_date")=$1`
	err = repository.DB.QueryRow(statement, month).Scan(&res)

	return res, err
}
