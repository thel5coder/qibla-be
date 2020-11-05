package requests

type UserRequest struct {
	Name             string            `json:"name"`
	UserName         string            `json:"user_name"`
	Email            string            `json:"email"`
	MobilePhone      string            `json:"mobile_phone"`
	Password         string            `json:"password"`
	RoleID           string            `json:"role_id"`
	IsActive         bool              `json:"is_active"`
	ProfilePictureID string            `json:"profile_picture_id"`
	OdoUserID        string            `json:"odo_user_id"`
	IsAdminPanel     bool              `json:"is_admin_panel"`
	MenuUsers        []MenuUserRequest `json:"menu_users"`
}
