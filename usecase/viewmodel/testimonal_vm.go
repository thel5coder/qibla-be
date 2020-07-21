package viewmodel

type TestimonialVm struct {
	ID                   string `json:"id"`
	WebContentCategoryID string `json:"web_content_category_id"`
	FileID               string `json:"file_id"`
	Path                 string `json:"file_path"`
	CustomerName         string `json:"customer_name"`
	JobPosition          string `json:"job_position"`
	Testimony            string `json:"testimony"`
	Rating               int    `json:"rating"`
	IsActive             bool   `json:"is_active"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	DeletedAt            string `json:"deleted_at"`
}
