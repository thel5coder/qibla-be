package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type FileRepository struct{
	DB *sql.DB
}

func NewFileRepository(DB *sql.DB) contracts.IFileRepository{
	return &FileRepository{DB: DB}
}

func (repository FileRepository) ReadBy(column, value string) (data models.File, err error) {
	statement := `select * from "files" where `+column+`=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,value).Scan(
		&data.ID,
		&data.Name,
		&data.Path,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.UserID,
		&data.Type,
		)

	return data,err
}

func (repository FileRepository) Add(input viewmodel.FileVm) (res string, err error) {
	statement := `insert into "files" ("name","path","created_at","updated_at") values($1,$2,$3,$4) returning "id"`
	err = repository.DB.QueryRow(statement,input.Name,input.Path,datetime.StrParseToTime(input.CreatedAt,time.RFC3339),datetime.StrParseToTime(input.UpdatedAt,time.RFC3339)).Scan(&res)

	return res,err
}

func (repository FileRepository) Delete(ID,updatedAt,deletedAt string) (res string, err error) {
	statement := `update "files" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID).Scan(&res)

	return res,err
}

func (repository FileRepository) CountBy(column,value string) (res int,err error){
	statement := `select count("id") from "files" where `+column+`=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,value).Scan(&res)

	return res,err
}

