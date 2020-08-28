package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ICalendarRepository interface {
	BrowseByYearMonth(yearMonth string) (data []models.Calendar, err error)

	ReadBy(column, value string) (data models.Calendar, err error)

	Edit(input viewmodel.CalendarVm,tx *sql.Tx) (err error)

	Add(input viewmodel.CalendarVm,tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string,tx *sql.Tx) (err error)

	CountBy(ID, column, value string) (res int, err error)
}
