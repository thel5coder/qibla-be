package requests

type RegisterByOauthRequest struct {
	Type           string `json:"type"`
	Token          string `json:"token"`
	FcmDeviceToken string `json:"fcm_device_token"`
}
