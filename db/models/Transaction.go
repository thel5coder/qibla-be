package models

import "database/sql"

type Transaction struct {
	ID                string         `json:"id"`
	UserID            string         `json:"user_id"`
	InvoiceNumber     sql.NullString `json:"invoice_number"`
	TrxID             sql.NullString `json:"trx_id"`
	DueDate           sql.NullString `json:"due_date"`
	DueDatePeriod     sql.NullInt32  `json:"due_date_period"`
	PaymentStatus     sql.NullString `json:"payment_status"`
	PaymentMethodCode sql.NullInt32  `json:"payment_method_code"`
	VaNumber          sql.NullString `json:"va_number"`
	BankName          sql.NullString `json:"bank_name"`
	Direction         string         `json:"direction"`
	TransactionType   string         `json:"transaction_type"`
	PaidDate          sql.NullString `json:"paid_date"`
	TransactionDate   string         `json:"transaction_date"`
	UpdatedAt         string         `json:"updated_at"`
	Total             float32        `json:"total"`
}
