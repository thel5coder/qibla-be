package viewmodel

import "database/sql"

type CalendarVm struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Start       string         `json:"start"`
	End         string         `json:"end"`
	Description string         `json:"description"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	DeletedAt   sql.NullString `json:"deleted_at"`
}
