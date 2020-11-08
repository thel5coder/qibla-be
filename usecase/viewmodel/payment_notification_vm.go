package viewmodel

type PaymentNotificationVm struct {
	Response      string `json:"response"`
	TransactionID string `json:"transaction_id"`
	TrxID         string `json:"trx_id"`
	MerchantID    string `json:"merchant_id"`
	Merchant      string `json:"merchant"`
	BillNo        string `json:"bill_no"`
	ResponseCode  string `json:"response_code"`
	ResponseDesc  string `json:"response_desc"`
	ResponseDate  string `json:"response_date"`
}
