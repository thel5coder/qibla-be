package requests

// DisbursementRequest ...
type DisbursementRequest struct {
	ContactID             string                      `json:"contact_id"`
	TransactionID         string                      `json:"transaction_id"`
	Total                 float64                     `json:"total"`
	Status                string                      `json:"status"`
	DisbursementType      string                      `json:"disbursement_type"`
	StartPeriod           string                      `json:"start_period"`
	EndPeriod             string                      `json:"end_period"`
	DisburseAt            string                      `json:"disburse_at"`
	AccountNumber         string                      `json:"account_number"`
	AccountName           string                      `json:"account_name"`
	AccountBankName       string                      `json:"account_bank_name"`
	AccountBankCode       string                      `json:"account_bank_code"`
	OriginAccountNumber   string                      `json:"origin_account_number"`
	OriginAccountName     string                      `json:"origin_account_name"`
	OriginAccountBankName string                      `json:"origin_account_bank_name"`
	OriginAccountBankCode string                      `json:"origin_account_bank_code"`
	Details               []DisbursementDetailRequest `json:"details"`
}

// DisbursementReqRequest ...
type DisbursementReqRequest struct {
	Data []DisbursementReqIDRequest `json:"data"`
}

// DisbursementReqIDRequest ...
type DisbursementReqIDRequest struct {
	ID string `json:"id"`
}
