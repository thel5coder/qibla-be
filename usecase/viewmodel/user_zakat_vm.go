package viewmodel

// UserZakatVm ...
type UserZakatVm struct {
	ID                           string `json:"id"`
	UserID                       string `json:"user_id"`
	UserEmail                    string `json:"user_email"`
	UserName                     string `json:"user_name"`
	TransactionID                string `json:"transaction_id"`
	TransactionInvoiceNumber     string `json:"transaction_invoice_number"`
	TransactionPaymentMethodCode int32  `json:"transaction_payment_method_code"`
	TransactionPaymentStatus     string `json:"transaction_payment_status"`
	TransactionDueDate           string `json:"transaction_due_date"`
	TransactionVaNumber          string `json:"transaction_va_number"`
	TransactionBankName          string `json:"transaction_bank_name"`
	ContactID                    string `json:"contact_id"`
	ContactBranchName            string `json:"contact_branch_name"`
	ContactTravelAgentName       string `json:"contact_travel_agent_name"`
	MasterZakatID                string `json:"master_zakat_id"`
	TypeZakat                    string `json:"type_zakat"`
	CurrentGoldPrice             int32  `json:"current_gold_price"`
	GoldNishab                   int32  `json:"gold_nishab"`
	Wealth                       int32  `json:"wealth"`
	Total                        int32  `json:"total"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
	DeletedAt                    string `json:"deleted_at"`
}
