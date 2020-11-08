package viewmodel

// DisbursementVm ...
type DisbursementVm struct {
	ID                           string  `json:"id"`
	ContactID                    string  `json:"contact_id"`
	TransactionID                string  `json:"transaction_id"`
	TransactionInvoiceNumber     string  `json:"transaction_invoice_number"`
	TransactionPaymentMethodCode int32   `json:"transaction_payment_method_code"`
	TransactionPaymentStatus     string  `json:"transaction_payment_status"`
	TransactionDueDate           string  `json:"transaction_due_date"`
	TransactionVaNumber          string  `json:"transaction_va_number"`
	TransactionBankName          string  `json:"transaction_bank_name"`
	Total                        float64 `json:"total"`
	Status                       string  `json:"status"`
	DisbursementType             string  `json:"disbursement_type"`
	StartPeriod                  string  `json:"start_period"`
	EndPeriod                    string  `json:"end_period"`
	DisburseAt                   string  `json:"disburse_at"`
	AccountNumber                string  `json:"account_number"`
	AccountName                  string  `json:"account_name"`
	AccountBankName              string  `json:"account_bank_name"`
	AccountBankCode              string  `json:"account_bank_code"`
	CreatedAt                    string  `json:"created_at"`
	UpdatedAt                    string  `json:"updated_at"`
	DeletedAt                    string  `json:"deleted_at"`
}
