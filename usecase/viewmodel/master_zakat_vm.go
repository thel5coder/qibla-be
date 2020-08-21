package viewmodel

type MasterZakatVm struct {
	ID               string `json:"id"`
	Slug             string `json:"slug"`
	TypeZakat        string `json:"type_zakat"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Amount           int32  `json:"amount"`
	CurrentGoldPrice int32  `json:"current_gold_price"`
	GoldNishab       int32  `json:"gold_nishab"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
}
