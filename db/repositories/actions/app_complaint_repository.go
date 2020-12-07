package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type AppComplaintRepository struct {
	DB *sql.DB
}

func NewAppComplaintRepository(DB *sql.DB) contracts.IAppComplaintRepository {
	return &AppComplaintRepository{DB: DB}
}

func (repository AppComplaintRepository) Browse(search, order, sort string, limit, offset int) (data []models.AppComplaint, count int, err error) {
	statement := `select * from "app_complaints" where (cast("created_at" as varchar) like $1 or lower("full_name") like $1 or lower("email") like $1 or "ticket_number" like $1 or lower("complaint_type") like $1 or lower("complaint") like $1 
                 or lower("solution") like $1 or cast("status" as varchar) like $1) and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		fmt.Println(err.Error())
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.AppComplaint{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.FullName,
			&dataTemp.Email,
			&dataTemp.TicketNumber,
			&dataTemp.ComplaintType,
			&dataTemp.Complaint,
			&dataTemp.Solution,
			&dataTemp.Status,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		data = append(data, dataTemp)
	}

	statement = `select count("id") from "app_complaints" where (cast("created_at" as varchar) like $1 or lower("full_name") like $1 or lower("email") like $1 or "ticket_number" like $1 or lower("complaint_type") like $1 or lower("complaint") like $1 
                 or lower("solution") like $1 or cast("status" as varchar) like $1) and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository AppComplaintRepository) ReadBy(column, value string) (data models.AppComplaint, err error) {
	statement := `select * from "app_complaints" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.FullName,
		&data.Email,
		&data.TicketNumber,
		&data.ComplaintType,
		&data.Complaint,
		&data.Solution,
		&data.Status,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

func (repository AppComplaintRepository) Edit(input viewmodel.AppComplaintVm) (res string, err error) {
	statement := `update "app_complaints" set "complaint"=$1, "solution"=$2, "status"=$3, "updated_at"=$4 where "id"=$5 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.Complaint,
		input.Solution,
		input.Status,
		datetime.StrParseToTime(input.UpdatedAt,time.RFC3339),
		input.ID,
		).Scan(&res)

	return res,err
}

func (repository AppComplaintRepository) Add(input viewmodel.AppComplaintVm) (res string, err error) {
	statement := `insert into "app_complaints" ("full_name","email","ticket_number","complaint_type","complaint","solution","status","created_at","updated_at")
                values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.FullName,
		input.Email,
		input.TicketNumber,
		input.ComplaintType,
		input.Complaint,
		input.Solution,
		input.Status,
		datetime.StrParseToTime(input.CreatedAt,time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt,time.RFC3339),
		).Scan(&res)

	return res,err
}

func (repository AppComplaintRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "app_complaints" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID).Scan(&res)

	return res,err
}

func (repository AppComplaintRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == ""{
		statement := `select count("id") from "app_complaints" where `+column+`=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,value).Scan(&res)
	}else{
		statement := `select count("id") from "app_complaints" where (`+column+`=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}

	return res,err
}

func (repository AppComplaintRepository) CountAll() (res int,err error){
	statement := `select count("id") from "app_complaints"`
	err = repository.DB.QueryRow(statement).Scan(&res)

	return res,err
}
