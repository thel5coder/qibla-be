package viewmodel

type AdminVm struct {
	ID              string                 `json:"id"`
	UserName        string                 `json:"user_name"`
	Email           string                 `json:"email"`
	Name            string                 `json:"name"`
	MobilePhone     string                 `json:"mobile_phone"`
	ProfilePicture  string                 `json:"profile_picture"`
	IsActive        bool                   `json:"is_active"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
	DeletedAt       string                 `json:"deleted_at"`
	OdooUserID      int                    `json:"odoo_user_id"`
	Role            RoleVm                 `json:"role"`
	MenuPermissions []MenuPermissionUserVm `json:"menu_permissions"`
}
