package viewmodel

type PromotionPlatformPositionVm struct {
	ID       string   `json:"id"`
	Platform string   `json:"platform"`
	Position []string `json:"position"`
}
