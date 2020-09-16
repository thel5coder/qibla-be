package viewmodel

type JamaahVm struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	MobilePhone    string `json:"mobile_phone"`
	ProfilePicture string `json:"profile_picture"`
	RoleID         string `json:"role_id"`
	RoleName       string `json:"role_name"`
	IsActive       bool   `json:"is_active"`
	IsPinSet       bool   `json:"is_pin_set"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
