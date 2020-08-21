package requests

type PrayRequest struct {
	Name      string `json:"name" validate:"required"`
	FileID    string `json:"file_id"`
	IsActive  bool   `json:"is_active"`
}
