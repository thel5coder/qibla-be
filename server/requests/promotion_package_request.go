package requests

type PromotionPackageRequest struct {
	PackageName string `json:"package_name" validate:"required"`
	IsActive    bool   `json:"is_active"`
}
