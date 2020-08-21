package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type VideoContentRepository struct{
	DB *sql.DB
}

func NewVideoContentRepository(DB *sql.DB) contracts.IVideoContentRepository{
	return &VideoContentRepository{DB: DB}
}

func (repository VideoContentRepository) Browse(order,sort string,limit,offset int) (data []models.VideoContent, count int, err error) {
	statement := `select * from "video_contents" order by `+order+` `+sort+` limit $1 offset $2`
	rows,err := repository.DB.Query(statement,limit,offset)
	if err != nil {
		return data,count,err
	}

	for rows.Next(){
		dataTemp := models.VideoContent{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Channel,
			&dataTemp.Links,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			)
		if err != nil {
			return data,count,err
		}
		data = append(data,dataTemp)
	}

	statement = `select count("id") from "video_contents"`
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data,count,err
	}

	return data,count,err
}

func (repository VideoContentRepository) Add(input viewmodel.VideoContentVm) (res string, err error) {
	statement := `insert into "video_contents" ("channel","links","created_at","updated_at") values ($1,$2,$3,$4) returning "id"`
	err = repository.DB.QueryRow(statement,input.Channel,input.Link,datetime.StrParseToTime(input.CreatedAt,time.RFC3339),datetime.StrParseToTime(input.UpdatedAt,time.RFC3339)).Scan(&res)

	return res,err
}

