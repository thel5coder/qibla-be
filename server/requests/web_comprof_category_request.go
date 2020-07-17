package requests

type WebComprofCategoryRequest struct {
	Name         string `json:"name" validate:"required"`
	CategoryType string `json:"category_type" validate:"required"`
}
