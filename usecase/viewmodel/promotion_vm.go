package viewmodel

type PromotionVm struct {
	ID                 string `json:"id"`
	PromotionPackageID string `json:"promotion_package_id"`
	PackagePromotion   string `json:"package_promotion"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	Platform           string `json:"platform"`
	Position           string `json:"position"`
	Price              string `json:"price"`
	Description        string `json:"description"`
	IsActive           bool   `json:"is_active"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	DeletedAt          string `json:"deleted_at"`
}
