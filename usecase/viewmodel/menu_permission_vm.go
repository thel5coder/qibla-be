package viewmodel

type MenuPermissionVm struct {
	ID         string `json:"id"`
	MenuID     string `json:"menu_id"`
	Permission string `json:"permission"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

type SelectedMenuPermissionVm struct {
	ID         string `json:"id"`
	Permission string `json:"permission"`
}
