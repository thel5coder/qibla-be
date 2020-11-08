package requests

// DisbursementDetailRequest ...
type DisbursementDetailRequest struct {
	DisbursementID string `json:"disbursement_id"`
	TransactionID  string `json:"transaction_id"`
}
