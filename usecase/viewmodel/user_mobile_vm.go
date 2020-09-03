package viewmodel

type JamaahVm struct {
	ID             string `json:"id"`
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	MobilePhone    string `json:"mobile_phone"`
	ProfilePicture string `json:"profile_picture"`
	IsActive       bool   `json:"is_active"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
