package viewmodel

type FaqVm struct {
	ID                   string `json:"id"`
	FaqCategoryName      string `json:"faq_category_name"`
	FaqListID            string `json:"faq_list_id"`
	Question             string `json:"question"`
	Answer               string `json:"answer"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	DeletedAt            string `json:"deleted_at"`
}
