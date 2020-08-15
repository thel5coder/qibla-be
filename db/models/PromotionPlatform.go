package models

type PromotionPlatform struct {
	ID          string `json:"id"`
	PromotionID string `json:"promotion_id"`
	Platform    string `json:"platform"`
}
