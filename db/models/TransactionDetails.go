package models

type TransactionDetails struct {
	ID            string  `db:"id"`
	TransactionID string  `db:"transaction_id"`
	Name          string  `db:"name"`
	Fee           float32 `db:"fee"`
	Price         float32 `db:"price"`
	Unit          string  `db:"unit"`
	Quantity      int     `db:"quantity"`
}
