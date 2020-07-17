package viewmodel

type MenuVm struct {
	ID              string   `json:"id"`
	MenuID          string   `json:"menu_id"`
	Name            string   `json:"name"`
	Url             string   `json:"url"`
	ParentID        string   `json:"parent_id"`
	IsActive        bool     `json:"is_active"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
	DeletedAt       string   `json:"deleted_at"`
	MenuPermissions []string `json:"menu_permissions"`
	ChildMenus      []MenuVm `json:"child_menus"`
}
