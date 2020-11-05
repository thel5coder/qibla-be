package viewmodel

type MenuUserVm struct {
	ID          string   `json:"id"`
	UserID      string   `json:"user_id"`
	MenuID      string   `json:"menu_id"`
	Permissions []string `json:"permissions"`
}
