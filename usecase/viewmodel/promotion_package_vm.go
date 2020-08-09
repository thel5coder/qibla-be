package viewmodel

type PromotionPackageVm struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	PackageName string `json:"package_name"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
