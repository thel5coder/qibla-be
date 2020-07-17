package requests

type MenuPermissionUserRequest struct {
	UserID           string `json:"user_id"`
	MenuPermissionID string `json:"menu_permission_id"`
}
