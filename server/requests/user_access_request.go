package requests

type UserAccessRequest struct {
	MenuID      string   `json:"menu_id"`
	Permissions []string `json:"permissions"`
}
