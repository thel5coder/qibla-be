package requests

type CalendarRequest struct {
	Title       string `json:"title" validate:"required"`
	Start       string `json:"start" validate:"required"`
	End         string `json:"end"`
	Description string `json:"description"`
}
