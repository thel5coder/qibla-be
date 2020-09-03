package viewmodel

type UserVm struct {
	ID             string `json:"id"`
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	MobilePhone    string `json:"mobile_phone"`
	Password       string `json:"password"`
	PIN            string `json:"pin"`
	IsActive       bool   `json:"is_active"`
	IsAdminPanel   bool   `json:"is_admin_panel"`
	RoleID         string `json:"role_id"`
	OdooUserID     int `json:"odo_user_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
