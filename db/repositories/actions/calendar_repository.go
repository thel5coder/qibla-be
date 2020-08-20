package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type CalendarRepository struct{
	DB *sql.DB
}

func NewCalendarRepository(DB *sql.DB) contracts.ICalendarRepository{
	return &CalendarRepository{DB: DB}
}

func (repository CalendarRepository) BrowseByYearMonth(year,month string) (data []models.Calendar, err error) {
	statement := `select * from "calendars" where date_part('year',timestamp,"start")=$1 and date_part('month',timestamp,'start')=$2 and "deleted_at" is null`
	rows,err := repository.DB.Query(statement,year,month)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := models.Calendar{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Title,
			&dataTemp.Start,
			&dataTemp.End,
			&dataTemp.Description,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			)
		if err != nil {
			return data,err
		}
		data = append(data,dataTemp)
	}

	return data,err
}

func (repository CalendarRepository) ReadBy(column, value string) (data models.Calendar, err error) {
	statement := `select * from "calendars" where `+column+`=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,value).Scan(
		&data.ID,
		&data.Title,
		&data.Start,
		&data.End,
		&data.Description,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		)

	return data,err
}

func (repository CalendarRepository) Edit(input viewmodel.CalendarVm) (res string, err error) {
	statement := `update "calendars" set "title"=$1, "start"=$2, "end"=$3, "description"=$4, "updated_at"=$5 where "id"=$6 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.Title,
		datetime.StrParseToTime(input.Start,time.RFC3339),
		datetime.StrParseToTime(input.End,time.RFC3339),
		input.Description,
		datetime.StrParseToTime(input.UpdatedAt,time.RFC3339),
		input.ID,
		).Scan(&res)

	return res,err
}

func (repository CalendarRepository) Add(input viewmodel.CalendarVm) (res string, err error) {
	statement := `insert into "calendars" ("title","start","end","description","created_at","updated_at") values($1,$2,$3,$4,$5,$6) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.Title,
		datetime.StrParseToTime(input.Start,time.RFC3339),
		datetime.StrParseToTime(input.End,time.RFC3339),
		input.Description,
		datetime.StrParseToTime(input.CreatedAt,time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt,time.RFC3339),
		).Scan(&res)

	return res,err
}

func (repository CalendarRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "calendars" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID).Scan(&res)

	return res,err
}

func (repository CalendarRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == ""{
		statement := `select count("id") from "calendars" where `+column+`=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,value).Scan(&res)
	}else{
		statement := `select count("id") from "calendars" where (`+column+`=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}

	return res,err
}

