package requests

type PromotionRequest struct {
	PromotionPackageID string                             `json:"promotion_package_id"`
	PackagePromotion   string                             `json:"package_promotion"`
	StartDate          string                             `json:"start_date"`
	EndDate            string                             `json:"end_date"`
	Position           []PromotionPlatformPositionRequest `json:"position"`
	Price              int                                `json:"price"`
	Description        string                             `json:"description"`
	IsActive           bool                               `json:"is_active"`
}
