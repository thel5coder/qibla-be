package viewmodel

type AdminVm struct {
	User       UserVm                  `json:"user"`
	MenuAccess []AdminUserAccessMenuVm `json:"menu_access"`
}

type AdminUserAccessMenuVm struct {
	MenuID     string   `json:"menu_id"`
	Permission []string `json:"permission"`
}
