package requests

type PartnerStatusRequest struct {
	Password string `json:"password"`
	Reason   string `json:"reason"`
	IsActive bool   `json:"is_active"`
}
