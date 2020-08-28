package requests

type SatisfactionCategoryRequest struct {
	ParentID      string                           `json:"parent_id"`
	SubCategories []SatisfactionSubCategoryRequest `json:"sub_categories"`
}

type SatisfactionSubCategoryRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
}
