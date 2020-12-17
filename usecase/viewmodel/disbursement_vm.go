package viewmodel

// DisbursementVm ...
type DisbursementVm struct {
	ID                           string      `json:"id"`
	ContactID                    string      `json:"contact_id"`
	ContactBranchName            string      `json:"contact_branch_name"`
	ContactTravelAgentName       string      `json:"contact_travel_agent_name"`
	ContactAddress       string      `json:"contact_address"`
	ContactPhoneNumber       string      `json:"contact_phone_number"`
	TransactionID                string      `json:"transaction_id"`
	TransactionInvoiceNumber     string      `json:"transaction_invoice_number"`
	TransactionPaymentMethodCode int32       `json:"transaction_payment_method_code"`
	TransactionPaymentStatus     string      `json:"transaction_payment_status"`
	TransactionDueDate           string      `json:"transaction_due_date"`
	TransactionVaNumber          string      `json:"transaction_va_number"`
	TransactionBankName          string      `json:"transaction_bank_name"`
	Total                        float64     `json:"total"`
	Status                       string      `json:"status"`
	DisbursementType             string      `json:"disbursement_type"`
	StartPeriod                  string      `json:"start_period"`
	EndPeriod                    string      `json:"end_period"`
	DisburseAt                   string      `json:"disburse_at"`
	AccountNumber                string      `json:"account_number"`
	AccountName                  string      `json:"account_name"`
	AccountBankName              string      `json:"account_bank_name"`
	AccountBankCode              string      `json:"account_bank_code"`
	OriginAccountNumber          string      `json:"origin_account_number"`
	OriginAccountName            string      `json:"origin_account_name"`
	OriginAccountBankName        string      `json:"origin_account_bank_name"`
	OriginAccountBankCode        string      `json:"origin_account_bank_code"`
	PaymentDetails               interface{} `json:"payment_details"`
	CreatedAt                    string      `json:"created_at"`
	UpdatedAt                    string      `json:"updated_at"`
	DeletedAt                    string      `json:"deleted_at"`
}
