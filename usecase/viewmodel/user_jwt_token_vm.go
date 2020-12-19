package viewmodel

type UserJwtTokenVm struct {
	Token            string `json:"token"`
	ExpTime          string `json:"exp_time"`
	RefreshToken     string `json:"refresh_token"`
	ExpRefreshToken  string `json:"exp_refresh_token"`
	IsPinSet         bool   `json:"is_pin_set"`
	IsFingerPrintSet bool   `json:"is_finger_print_set"`
}
