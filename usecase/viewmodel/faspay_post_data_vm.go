package viewmodel

type FaspayPostDataVm struct {
	RequestDescription string                 `json:"request_description"`
	MerchantID         string                 `json:"merchant_id"`
	Merchant           string                 `json:"merchant"`
	BillNo             string                 `json:"bill_no"`
	BillDate           string                 `json:"bill_date"`
	BillExpired        string                 `json:"bill_expired"`
	BillDesc           string                 `json:"bill_desc"`
	BillCurrency       string                 `json:"bill_currency"`
	BillTotal          int                    `json:"bill_total"`
	PaymentChannel     int                    `json:"payment_channel"`
	PayType            int                    `json:"pay_type"`
	CustNo             string                 `json:"cust_no"`
	CustName           string                 `json:"cust_name"`
	Msisdn             string                 `json:"msisdn"`
	Email              string                 `json:"email"`
	Terminal           string                 `json:"terminal"`
	Signature          string                 `json:"signature"`
	Item               []ItemFaspayPostDataVm `json:"item"`
}

type ItemFaspayPostDataVm struct {
	Product     string `json:"product"`
	Amount      int    `json:"amount"`
	Qty         int    `json:"qty"`
	PaymentPlan string `json:"payment_plan"`
	Tenor       string `json:"tenor"`
	MerchantID  string `json:"merchant_id"`
}
