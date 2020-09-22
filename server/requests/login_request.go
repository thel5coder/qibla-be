package requests

type LoginRequest struct {
	UserName       string `json:"user_name" validate:"required"`
	Password       string `json:"password" validate:"min=6"`
	FcmDeviceToken string `json:"fcm_device_token"`
}
