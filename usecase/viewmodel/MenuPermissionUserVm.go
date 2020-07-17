package viewmodel

type MenuPermissionUserVm struct {
	MenuID           string `json:"menu_id"`
	MenuName         string `json:"menu_name"`
	MenuPermissionID string `json:"menu_permission_id"`
	Permission       string `json:"permission"`
}
