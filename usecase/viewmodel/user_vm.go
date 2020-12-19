package viewmodel

type UserVm struct {
	ID               string `json:"id"`
	UserName         string `json:"user_name"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	MobilePhone      string `json:"mobile_phone"`
	PIN              string `json:"pin"`
	IsActive         bool   `json:"is_active"`
	IsAdminPanel     bool   `json:"is_admin_panel"`
	IsPINSet         bool   `json:"is_pin_set"`
	IsFingerPrintSet bool   `json:"is_finger_print_set"`
	OdooUserID       int32  `json:"odo_user_id"`
	FcmDeviceToken   string `json:"fcm_device_token"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	Role             RoleVm `json:"role"`
	File             FileVm `json:"file"`
}
