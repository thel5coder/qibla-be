package requests

type EditJamaahStatusRequest struct {
	JamaahID int `json:"jamaah_id"`
	Status string `json:"status"`
}