package viewmodel

type UserVm struct {
	ID              string                 `json:"id"`
	UserName        string                 `json:"user_name"`
	Email           string                 `json:"email"`
	IsActive        bool                   `json:"is_active"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
	DeletedAt       string                 `json:"deleted_at"`
	Role            RoleVm                 `json:"role"`
	MenuPermissions []MenuPermissionUserVm `json:"menu_permissions"`
}
