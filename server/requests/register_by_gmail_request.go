package requests

type RegisterByOauthRequest struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}
