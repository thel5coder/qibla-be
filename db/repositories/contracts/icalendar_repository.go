package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ICalendarRepository interface {
	BrowseByYearMonth(yearMonth string) (data []models.Calendar, err error)

	ReadBy(column, value string) (data models.Calendar, err error)

	Edit(input viewmodel.CalendarVm) (res string, err error)

	Add(input viewmodel.CalendarVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
