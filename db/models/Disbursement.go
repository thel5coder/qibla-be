package models

import "database/sql"

// Disbursement ...
type Disbursement struct {
	ID                    string          `db:"id"`
	ContactID             string          `db:"contact_id"`
	Contact               Contact         `db:"contact"`
	TransactionID         sql.NullString  `db:"transaction_id"`
	Transaction           Transaction     `db:"transaction"`
	Total                 sql.NullFloat64 `db:"total"`
	Status                sql.NullString  `db:"status"`
	DisbursementType      sql.NullString  `db:"disbursement_type"`
	StartPeriod           sql.NullString  `db:"start_period"`
	EndPeriod             sql.NullString  `db:"end_period"`
	DisburseAt            sql.NullString  `db:"disburse_at"`
	AccountNumber         sql.NullString  `db:"account_number"`
	AccountName           sql.NullString  `db:"account_name"`
	AccountBankName       sql.NullString  `db:"account_bank_name"`
	AccountBankCode       sql.NullString  `db:"account_bank_code"`
	OriginAccountNumber   sql.NullString  `db:"origin_account_number"`
	OriginAccountName     sql.NullString  `db:"origin_account_name"`
	OriginAccountBankName sql.NullString  `db:"origin_account_bank_name"`
	OriginAccountBankCode sql.NullString  `db:"origin_account_bank_code"`
	PaymentDetails        sql.NullString  `db:"payment_details"`
	CreatedAt             string          `db:"created_at"`
	UpdatedAt             string          `db:"updated_at"`
	DeletedAt             sql.NullString  `db:"deleted_at"`
}

var (
	// DisbursementSelect ...
	DisbursementSelect = `SELECT def."id", def."contact_id", def."transaction_id", def."total", def."status",
	def."disbursement_type", def."start_period", def."end_period", def."disburse_at",
	def."account_number", def."account_name", def."account_bank_name", def."account_bank_code",
	def."origin_account_number", def."origin_account_name", def."origin_account_bank_name", def."origin_account_bank_code",
	def."payment_details", def."created_at", def."updated_at", def."deleted_at",
	t."invoice_number", t."payment_method_code", t."payment_status",
	t."due_date", t."va_number", t."bank_name",
	c."branch_name", c."travel_agent_name", c."account_bank_name", c."address", c."phone_number"
	FROM "disbursements" def
	LEFT JOIN "transactions" t ON t."id" = def."transaction_id"
	LEFT JOIN "contacts" c ON c."id" = def."contact_id"`
)
