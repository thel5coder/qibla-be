package viewmodel

// DisbursementDetailVm ...
type DisbursementDetailVm struct {
	ID                           string `json:"id"`
	DisbursementID               string `json:"disbursement_id"`
	TransactionID                string `json:"transaction_id"`
	TransactionInvoiceNumber     string `json:"transaction_invoice_number"`
	TransactionPaymentMethodCode int32  `json:"transaction_payment_method_code"`
	TransactionPaymentStatus     string `json:"transaction_payment_status"`
	TransactionDueDate           string `json:"transaction_due_date"`
	TransactionVaNumber          string `json:"transaction_va_number"`
	TransactionBankName          string `json:"transaction_bank_name"`
}
