package viewmodel

type PromotionVm struct {
	ID                   string                        `json:"id"`
	PromotionPackageID   string                        `json:"promotion_package_id"`
	PromotionPackageName string                        `json:"promotion_package_name"`
	PackagePromotion     string                        `json:"package_promotion"`
	StartDate            string                        `json:"start_date"`
	EndDate              string                        `json:"end_date"`
	Positions            []PromotionPlatformPositionVm `json:"positions"`
	Platform             []PromotionPlatformVm         `json:"platform"`
	Price                int                           `json:"price"`
	Description          string                        `json:"description"`
	IsActive             bool                          `json:"is_active"`
	CreatedAt            string                        `json:"created_at"`
	UpdatedAt            string                        `json:"updated_at"`
	DeletedAt            string                        `json:"deleted_at"`
}

type PromotionPlatformVm struct {
	ID       string               `json:"id"`
	Platform string               `json:"platform"`
	Position []PlatformPositionVm `json:"position"`
}

type PlatformPositionVm struct {
	ID       string `json:"id"`
	Position string `json:"position"`
}
