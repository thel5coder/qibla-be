package models

import "database/sql"

// Disbursement ...
type Disbursement struct {
	ID               string          `db:"id"`
	TransactionID    string          `db:"transaction_id"`
	Transaction      Transaction     `db:"transaction"`
	Total            sql.NullFloat64 `db:"total"`
	Status           sql.NullString  `db:"status"`
	DisbursementType sql.NullString  `db:"disbursement_type"`
	StartPeriod      sql.NullString  `db:"start_period"`
	EndPeriod        sql.NullString  `db:"end_period"`
	DisburseAt       sql.NullString  `db:"disburse_at"`
	CreatedAt        string          `db:"created_at"`
	UpdatedAt        string          `db:"updated_at"`
	DeletedAt        sql.NullString  `db:"deleted_at"`
}

var (
	// DisbursementSelect ...
	DisbursementSelect = `SELECT def."id", def."transaction_id", def."total", def."status",
	def."disbursement_type", def."start_period", def."end_period", def."disburse_at",
	def."created_at", def."updated_at", def."deleted_at",
	t."invoice_number", t."payment_method_code", t."payment_status",
	t."due_date", t."va_number", t."bank_name"
	FROM "disbursements" def
	LEFT JOIN "transactions" t ON t."id" = def."transaction_id"`
)
