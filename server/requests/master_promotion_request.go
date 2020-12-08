package requests

type MasterPromotionRequest struct {
	PackageName string `json:"package_name" validate:"required"`
	IsActive    bool   `json:"is_active"`
}
