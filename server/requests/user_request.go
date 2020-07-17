package requests

type UserRequest struct {
	UserName   string              `json:"user_name"`
	Email      string              `json:"email"`
	Password   string              `json:"password"`
	RoleID     string              `json:"role_id"`
	IsActive   bool                `json:"is_active"`
	UserAccess []UserAccessRequest `json:"user_access"`
}