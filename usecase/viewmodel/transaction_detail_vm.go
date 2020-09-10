package viewmodel

type TransactionDetailVm struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Unit     string  `json:"unit"`
	Fee      float32 `json:"fee"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
	SubTotal float32 `json:"sub_total"`
}
