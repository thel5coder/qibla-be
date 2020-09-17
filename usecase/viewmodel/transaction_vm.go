package viewmodel

type TransactionVm struct {
	ID                string                 `json:"id"`
	UserID            string                 `json:"user_id"`
	InvoiceNumber     string                 `json:"invoice_number"`
	TrxID             string                 `json:"trx_id"`
	DueDate           string                 `json:"due_date"`
	DueDatePeriod     int32                  `json:"due_date_period"`
	PaymentStatus     string                 `json:"payment_status"`
	PaymentMethodCode int32                  `json:"payment_method_code"`
	VaNumber          string                 `json:"va_number"`
	BankName          string                 `json:"bank_name"`
	Direction         string                 `json:"direction"`
	TransactionType   string                 `json:"transaction_type"`
	PaidDate          string                 `json:"paid_date"`
	TransactionDate   string                 `json:"transaction_date"`
	UpdatedAt         string                 `json:"updated_at"`
	Total             float32                `json:"total"`
	Details           []TransactionDetailVm  `json:"details"`
	FaspayResponse    map[string]interface{} `json:"faspay_response"`
}
