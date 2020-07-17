package requests

type AddMenuRequest struct {
	MenuID   string `json:"menu_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Url      string `json:"url" validate:"required"`
	ParentID string `json:"parent_id"`
}

type MenuRequest struct {
	Menus []AddMenuRequest `json:"menus"`
}

type EditMenuRequest struct {
	ID                  string                  `json:"id"`
	Name                string                  `json:"name"`
	Url                 string                  `json:"url"`
	IsActive            bool                    `json:"is_active"`
	SelectedPermissions []MenuPermissionRequest `json:"selected_permissions"`
	DeletedPermissions  []string                `json:"deleted_permissions"`
}
