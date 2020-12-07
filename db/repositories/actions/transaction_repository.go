package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/enums"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(DB *sql.DB) contracts.ITransactionRepository {
	return &TransactionRepository{DB: DB}
}

const (
	transactionSelect = `select t."id",t."user_id",t."invoice_number",t."trx_id",t."due_date",t."due_date_period",t."payment_status",
                         t."payment_method_code",t."va_number",t."bank_name",t."direction",t."transaction_type",t."paid_date",t."invoice_status",
                         t."transaction_date",t."updated_at",t."total",t."fee_qibla",t."is_disburse",t."is_disburse_allowed",t."jamaah_count"`
	jointTransaction   = `left`
	groupByTransaction = `group by t."id"`
)

var (
	// DefaultDirectionInvoice ...
	DefaultDirectionInvoice = "in"

	// TransactionFieldDate ...
	TransactionFieldDate = "transaction_date"
	// TransactionFieldInvoice ...
	TransactionFieldInvoice = "invoice_number"
	// TransactionFieldTrxID ...
	TransactionFieldTrxID = "trx_id"
	// TransactionFieldDueDate ...
	TransactionFieldDueDate = "due_date"
	// TransactionFieldDueDatePeriod ...
	TransactionFieldDueDatePeriod = "due_date_period"
	// TransactionFieldPaymentStatus ...
	TransactionFieldPaymentStatus = "payment_status"
	// TransactionFieldPaymentMethodCode ...
	TransactionFieldPaymentMethodCode = "payment_method_code"
	// TransactionFieldVaNumber ...
	TransactionFieldVaNumber = "va_number"
	// TransactionFieldBankName ...
	TransactionFieldBankName = "bank_name"
	// TransactionFieldTransactionType ...
	TransactionFieldTransactionType = "transaction_type"
	// TransactionFieldPaidDate ...
	TransactionFieldPaidDate = "paid_date"
	// TransactionFieldTransactionDate ...
	TransactionFieldTransactionDate = "transaction_date"
	// TransactionFieldTotal ...
	TransactionFieldTotal = "total"
	// TransactionFieldFeeQibla ...
	TransactionFieldFeeQibla = "fee_qibla"
	// TransactionFieldIsDisburse ...
	TransactionFieldIsDisburse = "is_disburse"

	// DefaultSortAsc ...
	DefaultSortAsc = "asc"
	// DefaultSortDesc ...
	DefaultSortDesc = "desc"
)

func (repository TransactionRepository) scanRow(row *sql.Row) (res models.Transaction, err error) {
	err = row.Scan(&res.ID, &res.UserID, &res.InvoiceNumber, &res.TrxID, &res.DueDate, &res.DueDatePeriod, &res.PaymentStatus, &res.PaymentMethodCode, &res.VaNumber, &res.BankName,
		&res.Direction, &res.TransactionType, &res.PaidDate, &res.InvoiceStatus, &res.TransactionDate, &res.UpdatedAt, &res.Total, &res.FeeQibla, &res.IsDisburse, &res.IsDisburseAllowed)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository TransactionRepository) scanRows(rows *sql.Rows) (res models.Transaction, err error) {
	err = rows.Scan(&res.ID, &res.UserID, &res.InvoiceNumber, &res.TrxID, &res.DueDate, &res.DueDatePeriod, &res.PaymentStatus, &res.PaymentMethodCode, &res.VaNumber, &res.BankName,
		&res.Direction, &res.TransactionType, &res.PaidDate, &res.InvoiceStatus, &res.TransactionDate, &res.UpdatedAt, &res.Total, &res.FeeQibla, &res.IsDisburse, &res.IsDisburseAllowed)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository TransactionRepository) Browse(search, order, sort string, limit, offset int) (data []models.Transaction, count int, err error) {
	return data, count, err
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

// BrowseInvoices ...
func (repository TransactionRepository) BrowseInvoices(filters map[string]interface{}, order, sort string, limit, offset int) (data []models.Transaction, count int, err error) {
	var filterStatement string

	if val, ok := filters["name"]; ok {
		filterStatement += ` AND (lower(tp."full_name")='` + val.(string) + `' or lower(c."travel_agent_name")='` + val.(string) + `')`
	}
	if val, ok := filters["transaction_type"]; ok {
		filterStatement += `  AND inv.transaction_type = '` + val.(string) + `'`
	}
	if val, ok := filters["invoice_number"]; ok {
		filterStatement += `  AND inv.invoice_number = '` + val.(string) + `'`
	}
	if val, ok := filters["wumber_of_worshipers"]; ok {
		filterStatement += `  AND inv.wumber_of_worshipers = '` + val.(string) + `'`
	}
	if val, ok := filters["transaction_date"]; ok {
		filterStatement += `  AND inv.transaction_date = '` + val.(string) + `'`
	}
	if val, ok := filters["fee_qibla"]; ok {
		filterStatement += `  AND inv.fee_qibla = '` + val.(string) + `'`
	}
	if val, ok := filters["due_date"]; ok {
		filterStatement += `  AND inv.due_date = '` + val.(string) + `'`
	}
	if val, ok := filters["due_date_period"]; ok {
		filterStatement += `  AND inv.due_date_period = '` + val.(string) + `'`
	}
	if val, ok := filters["total"]; ok {
		filterStatement += `  AND inv.total = '` + val.(string) + `'`
	}
	if val, ok := filters["payment_status"]; ok {
		filterStatement += `  AND inv.payment_status = '` + val.(string) + `'`
	}
	if val, ok := filters["startDate"]; ok {
		filterStatement += ` and p.start_date < '` + val.(string) + `'`
	}

	statement := `SELECT inv."id", inv."transaction_type", inv."invoice_number", inv."fee_qibla",
	inv."total", inv."due_date", inv."due_date_period", inv."payment_status", inv."paid_date", inv."direction",
	inv."transaction_date", inv."updated_at", tp."full_name" as partner_name, c."travel_agent_name"
	FROM transactions inv
	LEFT JOIN user_tour_purchase_transactions tt ON tt.transaction_id = inv."id"
	LEFT JOIN user_tour_purchase_participants tp ON tp.user_tour_purchase_id = tt.user_tour_purchase_id
	LEFT JOIN partners p on p.user_id = inv.user_id
	LEFT JOIN contacts c on c."id" = p.contact_id
	WHERE inv."direction" = $1 AND inv."deleted_at" IS NULL ` + filterStatement + `
	ORDER BY inv.` + order + ` ` + sort + ` limit $2 offset $3`

	rows, err := repository.DB.Query(statement, DefaultDirectionInvoice, limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.Transaction{}

		err = rows.Scan(&dataTemp.ID, &dataTemp.TransactionType, &dataTemp.InvoiceNumber,
			&dataTemp.FeeQibla, &dataTemp.Total, &dataTemp.DueDate, &dataTemp.DueDatePeriod,
			&dataTemp.PaymentStatus, &dataTemp.PaidDate, &dataTemp.Direction, &dataTemp.TransactionDate,
			&dataTemp.UpdatedAt, &dataTemp.PartnerName, &dataTemp.TravelAgentName,
		)

		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `SELECT COUNT(1) FROM transactions inv
		LEFT JOIN user_tour_purchase_transactions tt ON tt.transaction_id = inv."id"
		LEFT JOIN user_tour_purchase_participants tp ON tp.user_tour_purchase_id = tt.user_tour_purchase_id
		LEFT JOIN partners p on p.user_id = inv.user_id
		LEFT JOIN contacts c on c."id" = p.contact_id
		WHERE inv."direction" = $1 AND inv."deleted_at" IS NULL ` + filterStatement

	err = repository.DB.QueryRow(statement, DefaultDirectionInvoice).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository TransactionRepository) ReadBy(column, value, operator string) (data models.Transaction, err error) {
	statement := transactionSelect + ` from "transactions" t ` + jointTransaction + ` where ` + column + `` + operator + `$1 ` + groupByTransaction
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
