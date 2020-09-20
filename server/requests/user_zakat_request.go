package requests

// UserZakatRequest ...
type UserZakatRequest struct {
	TransactionID     string `json:"transaction_id"`
	ContactID         string `json:"contact_id" validate:"required"`
	PaymentMethodCode int32  `json:"payment_method_code" validate:"required"`
	BankName          string `json:"bank_name" validate:"required"`
	MasterZakatID     string `json:"master_zakat_id"`
	TypeZakat         string `json:"type_zakat" validate:"required,oneof=maal penghasilan"`
	CurrentGoldPrice  int32  `json:"current_gold_price"`
	GoldNishab        int32  `json:"gold_nishab"`
	Wealth            int32  `json:"wealth" validate:"required"`
	Total             int32  `json:"total"`
}

// UserZakatPaymentRequest ...
type UserZakatPaymentRequest struct {
	PaymentMethodCode int32  `json:"payment_method_code" validate:"required"`
	BankName          string `json:"bank_name" validate:"required"`
}
