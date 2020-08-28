package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type ICalendarParticipantRepository interface {
	BrowseByCalendarID(calendarID string) (data []models.CalendarParticipant,err error)

	Add(calendarID ,participantEmail string,tx *sql.Tx) (err error)

	Delete(calendarID string,tx *sql.Tx) (err error)
}
