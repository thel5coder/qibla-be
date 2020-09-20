package requests

type TransactionRequest struct {
	UserID             string                     `json:"user_id"`
	DueDate            string                     `json:"due_date"`
	DueDateAging       int32                      `json:"due_date_aging"`
	BankName           string                     `json:"bank_name"`
	PaymentMethodeCode int32                      `json:"payment_methode_code"`
	TransactionDetail  []TransactionDetailRequest `json:"transaction_detail"`
	FaspayBody         FaspayPostRequest          `json:"faspay_body"`
}
