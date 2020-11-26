package flip

// Bank ...
type Bank struct {
	BankCode string  `json:"bank_code"`
	Fee      float64 `json:"fee"`
	Name     string  `json:"name"`
	Queue    float64 `json:"queue"`
	Status   string  `json:"status"`
}
