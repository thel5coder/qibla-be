package models

type PromotionPosition struct {
	ID                  string `json:"id"`
	PromotionPlatformID string `json:"promotion_platform_id"`
	Position            string `json:"position"`
}
