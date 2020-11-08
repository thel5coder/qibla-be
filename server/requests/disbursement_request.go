package requests

// DisbursementRequest ...
type DisbursementRequest struct {
	TransactionID    string                      `json:"transaction_id"`
	Total            float64                     `json:"total"`
	Status           string                      `json:"status"`
	DisbursementType string                      `json:"disbursement_type"`
	StartPeriod      string                      `json:"start_period"`
	EndPeriod        string                      `json:"end_period"`
	DisburseAt       string                      `json:"disburse_at"`
	Details          []DisbursementDetailRequest `json:"details"`
}
