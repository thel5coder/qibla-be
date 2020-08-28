package models

type CalendarParticipant struct {
	ID         string `db:"id"`
	CalendarID string `db:"calendar_id"`
	Email      string `db:"email"`
}
