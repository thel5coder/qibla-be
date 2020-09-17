package requests

type EditProfileRequest struct {
	FullName         string `json:"full_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	MobilePhone      string `json:"mobile_phone"`
	ProfilePictureID string `json:"profile_picture_id"`
}
