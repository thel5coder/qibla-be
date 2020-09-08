package viewmodel

type PartnerExtraProductVm struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	Price            float32 `json:"price"`
	PriceUnit        string  `json:"price_unit"`
	Session          string  `json:"session"`
}
