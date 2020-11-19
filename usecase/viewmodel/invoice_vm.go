package viewmodel

// InvoiceVM ...
type InvoiceVm struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	TransactionType  string  `json:"transaction_type"`
	InvoiceNumber    string  `json:"invoice_number"`
	NumberWorshipers int32   `json:"number_worshipers"`
	FeeQibla         float64 `json:"fee_qibla"`
	Total            float64 `json:"total"`
	DueDate          string  `json:"due_date"`
	BillingStatus    string  `json:"billing_status"`
	DueDatePeriod    int32   `json:"due_date_period"`
	PaymentStatus    string  `json:"payment_status"`
	PaidDate         string  `json:"paid_date"`
	InvoiceStatus    string  `json:"invoice_status"`
	Direction        string  `json:"direction"`
	TransactionDate  string  `json:"transaction_date"`
	UpdatedAt        string  `json:"updated_at"`
}
