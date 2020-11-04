package requests

type MenuUserRequest struct {
	MenuID          string   `json:"menu_id"`
	MenuPermissions []string `json:"menu_permissions"`
}
