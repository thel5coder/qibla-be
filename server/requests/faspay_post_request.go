package requests

type FaspayPostRequest struct {
	RequestTransaction string              `json:"request_transaction"`
	InvoiceNumber      string              `json:"invoice_number"`
	TransactionDate    string              `json:"transaction_date"`
	DueDate            string              `json:"due_date"`
	TransactionDesc    string              `json:"transaction_desc"`
	UserID             string              `json:"user_id"`
	Total              float32             `json:"total"`
	Item               []FaspayItemRequest `json:"item"`
	PaymentChannel     int                 `json:"payment_channel"`
}

type FaspayItemRequest struct {
	Product     string `json:"product"`
	Amount      int    `json:"amount"`
	Qty         int    `json:"qty"`
	PaymentPlan string `json:"payment_plan"`
	Tenor       string `json:"tenor"`
	MerchantID  string `json:"merchant_id"`
}
