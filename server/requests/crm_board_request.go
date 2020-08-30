package requests

type CrmBoardRequest struct {
	CrmStoryID        string  `json:"crm_story_id"`
	ContactID         string  `json:"contact_id"`
	Opportunity       float32 `json:"opportunity"`
	ProfitExpectation float32 `json:"profit_expectation"`
	Star              int     `json:"star"`
}

type CrmBoardEditStoryRequest struct {
	CrmStoryID string `json:"crm_story_id"`
}
