package requests

type FaspayPostRequest struct {
	RequestTransaction  string              `json:"request_transaction"`
	InvoiceNumber       string              `json:"invoice_number"`
	TransactionDate     string              `json:"transaction_date"`
	DueDate             string              `json:"due_date"`
	TransactionDesc     string              `json:"transaction_desc"`
	UserID              string              `json:"user_id"`
	CustomerName        string              `json:"customer_name"`
	CustomerEmail       string              `json:"customer_email"`
	CustomerPhoneNumber string              `json:"customer_phone_number"`
	Total               float32             `json:"total"`
	PaymentChannel      int32               `json:"payment_channel"`
	Item                []FaspayItemRequest `json:"item"`
}

type FaspayItemRequest struct {
	Product     string `json:"product"`
	Amount      int    `json:"amount"`
	Qty         int    `json:"qty"`
	PaymentPlan string `json:"payment_plan"`
	Tenor       string `json:"tenor"`
	MerchantID  string `json:"merchant_id"`
}
