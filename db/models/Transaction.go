package models

type Transaction struct {
	ID        string `json:"id"`
	InvoiceID string `json:"invoice_id"`
}
