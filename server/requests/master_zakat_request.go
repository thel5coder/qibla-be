package requests

type MasterZakatRequest struct {
	TypeZakat        string `json:"type_zakat"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Amount           int32  `json:"amount"`
	CurrentGoldPrice int32  `json:"current_gold_price"`
	GoldNishab       int32  `json:"gold_nishab"`
}
