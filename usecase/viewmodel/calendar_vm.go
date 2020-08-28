package viewmodel

type CalendarVm struct {
	ID           string   `json:"id"`
	Title        string   `json:"title"`
	Start        string   `json:"start"`
	End          string   `json:"end"`
	Description  string   `json:"description"`
	Participants []string `json:"participants"`
	Remember     int      `json:"remember"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
	DeletedAt    string   `json:"deleted_at"`
}
