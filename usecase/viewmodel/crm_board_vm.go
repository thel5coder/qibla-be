package viewmodel

type CrmBoardVm struct {
	ID                string  `json:"id"`
	CrmStoryID        string  `json:"crm_story_id"`
	ContactID         string  `json:"contact_id"`
	Opportunity       float32 `json:"opportunity"`
	ProfitExpectation float32 `json:"profit_expectation"`
	Star              int     `json:"star"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
}
