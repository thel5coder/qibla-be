package models

// DisbursementDetail ...
type DisbursementDetail struct {
	ID             string      `db:"id"`
	DisbursementID string      `db:"disbursement_id"`
	TransactionID  string      `db:"transaction_id"`
	Transaction    Transaction `db:"transaction"`
}

var (
	// DisbursementDetailSelect ...
	DisbursementDetailSelect = `SELECT def."id", def."disbursement_id", def."transaction_id",
	t."invoice_number", t."payment_method_code", t."payment_status",
	t."due_date", t."va_number", t."bank_name"
	FROM "disbursement_details" def
	LEFT JOIN "transactions" t ON t."id" = def."transaction_id"`
)
