package viewmodel

type SatisfactionCategoryVm struct {
	ID          string                   `json:"id"`
	ParentID    string                   `json:"parent_id"`
	Slug        string                   `json:"slug"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	IsActive    bool                     `json:"is_active"`
	CreatedAt   string                   `json:"created_at"`
	UpdatedAt   string                   `json:"updated_at"`
	Child       []SatisfactionCategoryVm `json:"child"`
}
