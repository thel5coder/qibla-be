package models

type TransactionHistory struct {
	ID            string `db:"id"`
	TransactionID string `db:"transaction_id"`
	Status        string `db:"status"`
	Response      string `db:"response"`
	CreatedAt     string `db:"created_at"`
	UpdatedAt     string `db:"updated_at"`
}
