package requests

type PromotionPlatformPositionRequest struct {
	Platform string   `json:"platform"`
	Position []string `json:"position"`
}
