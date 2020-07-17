package requests

type GlobalInfoCategorySettingRequest struct {
	GlobalInfoCategoryID string `json:"global_info_category_id" validate:"required"`
	Description          string `json:"description" validate:"required"`
	IsActive             bool   `json:"is_active"`
}
