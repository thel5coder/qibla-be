package viewmodel

type TermConditionVm struct {
	ID          string `json:"id"`
	TermName    string `json:"term_name"`
	TermType    string `json:"term_type"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
