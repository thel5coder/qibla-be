package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type CalendarParticipantRepository struct{
	DB *sql.DB
}

func NewCalendarParticipantRepository(db *sql.DB) contracts.ICalendarParticipantRepository{
	return &CalendarParticipantRepository{DB: db}
}

func (repository CalendarParticipantRepository) BrowseByCalendarID(calendarID string) (data []models.CalendarParticipant, err error) {
	statement := `select * from "calendar_participants" where "calendar_id"=$1`
	rows,err := repository.DB.Query(statement,calendarID)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := models.CalendarParticipant{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.CalendarID,
			&dataTemp.Email,
			)
		if err != nil {
			return data,err
		}
		data = append(data,dataTemp)
	}

	return data,err
}

func (CalendarParticipantRepository) Add(calendarID , participantEmail string, tx *sql.Tx) (err error) {
	statement := `insert into "calendar_participants" ("calendar_id","email") values($1,$2)`
	_,err = tx.Exec(statement,calendarID,participantEmail)

	return err
}

func (CalendarParticipantRepository) Delete(calendarID string, tx *sql.Tx) (err error) {
	statement := `delete from "calendar_participants" where "calendar_id"=$1`
	_,err = tx.Exec(statement,calendarID)

	return err
}

