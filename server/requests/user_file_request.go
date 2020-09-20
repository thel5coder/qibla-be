package requests

// UserZakatRequest ...
type UserZakatRequest struct {
	UserID           string `json:"user_id"`
	TransactionID    string `json:"transaction_id"`
	ContactID        string `json:"contact_id" validate:"required"`
	MasterZakatID    string `json:"master_zakat_id"`
	TypeZakat        string `json:"type_zakat" validate:"required,oneof=maal penghasilan"`
	CurrentGoldPrice int32  `json:"current_gold_price"`
	GoldNishab       int32  `json:"gold_nishab"`
	Wealth           int32  `json:"wealth" validate:"required"`
	Total            int32  `json:"total"`
}
