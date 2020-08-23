package viewmodel

type PrayVm struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	FileID    string `json:"file_id"`
	IsActive  bool   `json:"is_active"`
	File      FileVm `json:"file"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
