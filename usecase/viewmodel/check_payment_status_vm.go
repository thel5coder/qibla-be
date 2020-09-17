package viewmodel

type CheckPaymentStatus struct {
	Request    string `json:"request"`
	TrxID      string `json:"trx_id"`
	MerchantID string `json:"merchant_id"`
	BillNo     string `json:"bill_no"`
	Signature  string `json:"signature"`
}
