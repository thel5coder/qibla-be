package requests

type PaymentNotificationRequest struct {
	Request           string `json:"request"`
	TrxID             string `json:"trx_id"`
	MerchantID        string `json:"merchant_id"`
	Merchant          string `json:"merchant"`
	BillNo            string `json:"bill_no"`
	PaymentDate       string `json:"payment_date"`
	PaymentStatusCode string `json:"payment_status_code"`
	PaymentStatusDesc string `json:"payment_status_desc"`
	PaymentTotal      string `json:"payment_total"`
	PaymentChannelUid string `json:"payment_channel_uid"`
	PaymentChannel    string `json:"payment_channel"`
	Signature         string `json:"signature"`
}
