package requests

type addMenuRequest struct {
	Name     string `json:"name" validate:"required"`
	Url      string `json:"url" validate:"required"`
	ParentID string `json:"parent_id"`
}

type AddMenuRequest struct {
	Menus []addMenuRequest `json:"menus"`
}

type EditMenuRequest struct {
	Menus               []editMenuRequest       `json:"menus"`
	SelectedPermissions []MenuPermissionRequest `json:"selected_permissions"`
	DeletedPermissions  []string                `json:"deleted_permissions"`
}

type editMenuRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	IsActive bool   `json:"is_active"`
}
