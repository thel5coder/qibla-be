package viewmodel

type CrmStoryVm struct {
	ID          string       `json:"id"`
	Slug        string       `json:"slug"`
	Name        string       `json:"name"`
	ProfitCount float32      `json:"profit_count"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	CrmBoards   []CrmBoardVm `json:"crm_boards"`
}
