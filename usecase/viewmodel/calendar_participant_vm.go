package viewmodel

type CalendarParticipantVm struct {
	ID         string `json:"id"`
	CalendarID string `json:"calendar_id"`
	Email      string `json:"email"`
}
