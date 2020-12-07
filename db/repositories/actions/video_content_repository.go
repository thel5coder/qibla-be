package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type VideoContentRepository struct {
	DB *sql.DB
}

func NewVideoContentRepository(DB *sql.DB) contracts.IVideoContentRepository {
	return &VideoContentRepository{DB: DB}
}

func (repository VideoContentRepository) Browse(order, sort string, limit, offset int) (data []models.VideoContent, count int, err error) {
	statement := `select * from "video_contents" where "deleted_at" is null order by ` + order + ` ` + sort + ` limit $1 offset $2`
	rows, err := repository.DB.Query(statement, limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.VideoContent{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Channel,
			&dataTemp.Links,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.IsActive,
			&dataTemp.ChannelID,
		)
		if err != nil {
			return data, count, err
		}
		data = append(data, dataTemp)
	}

	statement = `select count("id") from "video_contents"`
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository VideoContentRepository) BrowseAll() (data []models.VideoContent, err error) {
	statement := `select * from "video_contents" where "deleted_at" is null order by "created_at"`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.VideoContent{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Channel,
			&dataTemp.Links,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.IsActive,
			&dataTemp.ChannelID,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository VideoContentRepository) ReadBy(column, value string) (data models.VideoContent, err error) {
	statement := `select * from "video_contents" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.Channel,
		&data.Links,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.IsActive,
		&data.ChannelID,
	)

	return data, err
}

func (repository VideoContentRepository) Edit(input viewmodel.VideoContentVm) (res string, err error) {
	statement := `update "video_contents" set "channel"=$1, "links"=$2, "is_active"=$3, "updated_at"=$4,"channel_id"=$5 where "id"=$6 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.Channel,
		input.Link,
		input.IsActive,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ChannelID,
		input.ID,
		).Scan(&res)

	return res, err
}

func (repository VideoContentRepository) Add(input viewmodel.VideoContentVm) (res string, err error) {
	statement := `insert into "video_contents" ("channel","channel_id","links","is_active","created_at","updated_at") values ($1,$2,$3,$4,$5,$6) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.Channel,
		input.ChannelID,
		input.Link,
		input.IsActive,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		).Scan(&res)

	return res, err
}

func (repository VideoContentRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "video_contents" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository VideoContentRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "video_contents" where lower(` + column + `)=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "video_contents" where (lower(` + column + `)=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
