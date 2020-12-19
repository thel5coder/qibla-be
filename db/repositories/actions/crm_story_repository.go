package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type CrmStoryRepository struct{
	DB *sql.DB
}

func NewCrmStoryRepository(db *sql.DB) contracts.ICrmStoryRepository{
	return &CrmStoryRepository{DB: db}
}

func (repository CrmStoryRepository) BrowseAll() (data []models.CrmStory, err error) {
	statement := `select * from "crm_stories" where "deleted_at" is null order by created_at asc`
	rows,err := repository.DB.Query(statement)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := models.CrmStory{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Slug,
			&dataTemp.Name,
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

func (repository CrmStoryRepository) ReadBy(column, value string) (data models.CrmStory, err error) {
	statement := `select * from "crm_stories" where `+column+`=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,value).Scan(
		&data.ID,
		&data.Slug,
		&data.Name,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		)

	return data,err
}

func (repository CrmStoryRepository) Edit(input viewmodel.CrmStoryVm) (res string, err error) {
	statement := `update "crm_stories" set "name"=$1,"slug"=$2,"updated_at"=$3 where "id"=$4 returning "id"`
	err = repository.DB.QueryRow(statement,input.Name,input.Slug,datetime.StrParseToTime(input.UpdatedAt,time.RFC3339),input.ID).Scan(&res)

	return res,err
}

func (repository CrmStoryRepository) Add(input viewmodel.CrmStoryVm) (res string, err error) {
	statement := `insert into "crm_stories" ("name","slug","created_at","updated_at") values($1,$2,$3,$4) returning "id"`
	err = repository.DB.QueryRow(statement,input.Name,input.Slug,datetime.StrParseToTime(input.CreatedAt,time.RFC3339),datetime.StrParseToTime(input.UpdatedAt,time.RFC3339)).Scan(&res)

	return res,err
}

func (repository CrmStoryRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "crm_stories" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID).Scan(&res)

	return res,err
}

func (repository CrmStoryRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == ""{
		statement := `select count("id") from "crm_stories" where `+column+`=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,value).Scan(&res)
	}else{
		statement := `select count("id") from "crm_stories" where (`+column+`=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}

	return res,err
}

