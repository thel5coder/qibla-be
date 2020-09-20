package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
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
                         t."transaction_date",t."updated_at",sum(td."price")`
const joinQuery = `left join "transaction_details" td on td."transaction_id"=t."id"`
const groupBy = `group by t."id"`

func (TransactionRepository) Browse(search, order, sort string, limit, offset int) (data []models.Transaction, count int, err error) {
	panic("implement me")
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
                  "bank_name","direction","transaction_type","transaction_date","updated_at")
                  values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) returning "id"`
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
	fmt.Println("dari uc")
	fmt.Println(updatedAt)
	fmt.Println("repo")
	now := time.Now().UTC().Format(time.RFC3339)

	fmt.Println(datetime.StrParseToTime(now, time.RFC3339))
	statement := `update "transactions" set "payment_status"=$1, "paid_date"=$2, "updated_at"=$3 where "id"=$3 returning "id"`
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

func (repository TransactionRepository) GetInvoiceCount(month int) (res int, err error) {
	statement := `select count("id") from "transactions" where EXTRACT(month from "transaction_date")=$1`
	err = repository.DB.QueryRow(statement, month).Scan(&res)

	return res, err
}
