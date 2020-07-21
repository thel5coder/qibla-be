package requests

type TestimonialRequest struct {
	WebContentCategoryID string `json:"web_content_category_id"`
	FileID               string `json:"file_id"`
	CustomerName         string `json:"customer_name"`
	JobPosition          string `json:"job_position"`
	Testimony            string `json:"testimony"`
	Rating               int    `json:"rating"`
	IsActive             bool   `json:"is_active"`
}
