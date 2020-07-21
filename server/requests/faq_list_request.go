package requests

type FaqListRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
