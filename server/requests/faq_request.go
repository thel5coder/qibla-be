package requests

type FaqRequest struct {
	WebContentCategoryID string           `json:"web_content_category_id"`
	FaqCategoryName      string           `json:"faq_category_name"`
	FaqLists             []FaqListRequest `json:"faq_lists"`
}
