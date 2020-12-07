package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type PrayRepository struct {
	DB *sql.DB
}

func NewPrayRepository(DB *sql.DB) contracts.IPrayRepository {
	return &PrayRepository{DB: DB}
}

func (repository PrayRepository) Browse(search, order, sort string, limit, offset int) (data []models.Pray, count int, err error) {
	statement := `select * from "prays" where (lower("name") like $1 or cast("updated_at" as varchar) like $1) and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.Pray{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Name,
			&dataTemp.FileID,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		data = append(data, dataTemp)
	}

	statement = `select count("id") from "prays" where (lower("name") like $1 or cast("updated_at" as varchar) like $1) and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository PrayRepository) ReadBy(column, value string) (data models.Pray, err error) {
	statement := `select * from "prays" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.Name,
		&data.FileID,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

func (repository PrayRepository) Edit(input viewmodel.PrayVm) (res string, err error) {
	statement := `update "prays" set "name"=$1,"file_id"=$2,"updated_at"=$3,"is_active"=$4 where "id"=$5 returning "id"`
	err = repository.DB.QueryRow(statement, input.Name, input.FileID, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.IsActive, input.ID).Scan(&res)

	return res, err
}

func (repository PrayRepository) Add(input viewmodel.PrayVm) (res string, err error) {
	statement := `insert into "prays" ("name","file_id","is_active","created_at","updated_at") values($1,$2,$3,$4,$5) returning "id"`
	err = repository.DB.QueryRow(statement, input.Name, input.FileID, input.IsActive, datetime.StrParseToTime(input.CreatedAt, time.RFC3339), datetime.StrParseToTime(input.UpdatedAt, time.RFC3339)).Scan(&res)

	return res, err
}

func (repository PrayRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "prays" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository PrayRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "prays" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "prays" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
