package models

type TransactionDetails struct {
	ID            string `db:"id"`
	TransactionID string `db:"transaction_id"`
	Name          string `db:"name"`
	Fee           string `db:"fee"`
	Price         string `db:"price"`
	Unit          string `db:"unit"`
	Quantity      string `db:"quantity"`
}
