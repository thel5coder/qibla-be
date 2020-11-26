package models

import "database/sql"

type Transaction struct {
	ID                string          `db:"id"`
	UserID            string          `db:"user_id"`
	InvoiceNumber     sql.NullString  `db:"invoice_number"`
	TrxID             sql.NullString  `db:"trx_id"`
	DueDate           sql.NullString  `db:"due_date"`
	DueDatePeriod     sql.NullInt32   `db:"due_date_period"`
	PaymentStatus     sql.NullString  `db:"payment_status"`
	InvoiceStatus     sql.NullString  `db:"invoice_status"`
	PaymentMethodCode sql.NullInt32   `db:"payment_method_code"`
	VaNumber          sql.NullString  `db:"va_number"`
	BankName          sql.NullString  `db:"bank_name"`
	Direction         string          `db:"direction"`
	TransactionType   string          `db:"transaction_type"`
	PaidDate          sql.NullString  `db:"paid_date"`
	TransactionDate   string          `db:"transaction_date"`
	UpdatedAt         string          `db:"updated_at"`
	Total             sql.NullFloat64 `db:"total"`
	FeeQibla          sql.NullFloat64 `db:"fee_qibla"`
	IsDisburse        sql.NullBool    `db:"is_disburse"`
	IsDisburseAllowed sql.NullBool    `db:"is_disburse_allowed"`
	Details           sql.NullString  `json:"details"`
	PartnerName       sql.NullString  `db:"partner_name"`
	TravelAgentName   sql.NullString  `db:"travel_agent_name"`
}
