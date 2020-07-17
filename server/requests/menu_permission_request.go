package requests

type MenuPermissionRequest struct {
	ID         string `json:"id"`
	MenuID     string `json:"menu_id"`
	Permission string `json:"permission"`
}
