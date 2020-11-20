package models

import "database/sql"

// Invoice ...
type Invoice struct {
	ID              string          `db:"id"`
	Name            sql.NullString  `db:"name"`
	TransactionType string          `db:"transaction_type"`
	InvoiceNumber   string          `db:"invoice_number"`
	FeeQibla        sql.NullFloat64 `db:"fee_qibla"`
	Total           sql.NullFloat64 `db:"total"`
	DueDate         sql.NullString  `db:"due_date"`
	DueDatePeriod   sql.NullInt32   `db:"due_date_period"`
	PaymentStatus   sql.NullString  `db:"payment_status"`
	PaidDate        sql.NullString  `db:"paid_date"`
	InvoiceStatus   sql.NullString  `db:"invoice_status"`
	Direction       string          `db:"direction"`
	TransactionDate string          `db:"transaction_date"`
	UpdatedAt       string          `db:"updated_at"`
}
