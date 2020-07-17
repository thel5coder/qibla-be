package viewmodel

type GlobalInfoCategorySettingVm struct {
	ID                     string `json:"id"`
	GlobalInfoCategoryID   string `json:"global_info_category_id"`
	GlobalInfoCategoryName string `json:"global_info_category_name"`
	Description            string `json:"description"`
	IsActive               bool   `json:"is_active"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
	DeletedAt              string `json:"deleted_at"`
}
