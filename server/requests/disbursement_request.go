package requests

// DisbursementRequest ...
type DisbursementRequest struct {
	ContactID        string                      `json:"contact_id"`
	TransactionID    string                      `json:"transaction_id"`
	Total            float64                     `json:"total"`
	Status           string                      `json:"status"`
	DisbursementType string                      `json:"disbursement_type"`
	StartPeriod      string                      `json:"start_period"`
	EndPeriod        string                      `json:"end_period"`
	DisburseAt       string                      `json:"disburse_at"`
	AccountNumber    string                      `json:"account_number"`
	AccountName      string                      `json:"account_name"`
	AccountBankName  string                      `json:"account_bank_name"`
	AccountBankCode  string                      `json:"account_bank_code"`
	Details          []DisbursementDetailRequest `json:"details"`
}
