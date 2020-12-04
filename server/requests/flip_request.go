package requests

// FlipDisbursementCallbackRequest ...
type FlipDisbursementCallbackRequest struct {
	Token string                  `json:"token"`
	Data  FlipDisbursementRequest `json:"data"`
}

// FlipDisbursementRequest ...
type FlipDisbursementRequest struct {
	ID            int     `json:"id"`
	Fee           float64 `json:"fee"`
	Amount        float64 `json:"amount"`
	Remark        string  `json:"remark"`
	Status        string  `json:"status"`
	Receipt       string  `json:"receipt"`
	UserID        int     `json:"user_id"`
	BankCode      string  `json:"bank_code"`
	BundleID      int     `json:"bundle_id"`
	Direction     string  `json:"direction"`
	Timestamp     string  `json:"timestamp"`
	CompanyID     int     `json:"company_id"`
	SenderBank    string  `json:"sender_bank"`
	AccountNumber string  `json:"account_number"`
	RecipientCity int     `json:"recipient_city"`
	RecipientName string  `json:"recipient_name"`
}
