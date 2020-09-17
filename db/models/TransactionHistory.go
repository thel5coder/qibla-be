package models

type TransactionHistory struct {
	ID        string `json:"id"`
	TrxID     string `json:"trx_id"`
	Status    string `json:"status"`
	Response  string `json:"response"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
