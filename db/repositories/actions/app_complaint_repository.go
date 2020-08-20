package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type AppComplaintRepository struct {
	DB *sql.DB
}

func NewAppComplaintRepository(DB *sql.DB) contracts.IAppComplaintRepository {
	return &AppComplaintRepository{DB: DB}
}

func (repository AppComplaintRepository) Browse(search, order, sort string, limit, offset int) (data []models.AppComplaint, count int, err error) {
	statement := `select * from "app_complaints" where (cast("created_at" as varchar) like $1 or "full_name" like $1 or "email" like $1 or "ticket_number" like $1 or "complaint_type" like $1 or "complaint" like $1 
                 or "solution" like $1 or "status" like $1) and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+search+"%", limit, offset)
	if err != nil {
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

	statement = `select count("id") from "app_complaints" where (cast("created_at" as varchar) like $1 or "full_name" like $1 or "email" like $1 or "ticket_number" like $1 or "complaint_type" like $1 or "complaint" like $1 
                 or "solution" like $1 or "status" like $1) and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+search+"%").Scan(&count)
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
	statement := `update "app_complaints" set "full_name"=$1, "email"=$2, "ticket_number"=$3, "complaint_type"=$4, "complaint"=$4, "solution"=$5, "status"=$6, "updated_at"=$7 where "id"=$8 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.FullName,
		input.Email,
		input.TicketNumber,
		input.ComplaintType,
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
