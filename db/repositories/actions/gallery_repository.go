package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"strings"
	"time"
)

type GalleryRepository struct{
	DB *sql.DB
}

func NewGalleryRepository(DB *sql.DB) contracts.IGalleryRepository{
	return &GalleryRepository{DB: DB}
}

func (repository GalleryRepository) Browse(search, order, sort string, limit, offset int) (data []models.Gallery, count int, err error) {
	statement := `select * from "galleries" where lower("gallery_name") like $1 and "deleted_at" is null order by `+order+` `+sort+` limit $2 offset $3`
	rows,err := repository.DB.Query(statement,"%"+strings.ToLower(search)+"%",limit,offset)
	if err != nil {
		return data,count,err
	}

	for rows.Next(){
		dataTemp := models.Gallery{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.WebContentCategoryID,
			&dataTemp.GalleryName,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			)
		if err != nil {
			return data,count,err
		}

		data = append(data,dataTemp)
	}

	statement = `select count("id") from "galleries" where lower("gallery_name") like $1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,"%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data,count,err
	}

	return data,count,err
}

func (repository GalleryRepository) ReadBy(column, value string) (data models.Gallery, err error) {
	statement := `select * from "galleries" where `+column+`=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,value).Scan(
		&data.ID,
		&data.WebContentCategoryID,
		&data.GalleryName,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		)

	return data,err
}

func (GalleryRepository) Edit(ID,galleryName, updatedAt string, tx *sql.Tx) (err error) {
	statement := `update "galleries" set "gallery_name"=$1, "updated_at"=$2 where "id"=$3`
	_,err = tx.Exec(statement,galleryName,datetime.StrParseToTime(updatedAt,time.RFC3339),ID)

	return err
}

func (GalleryRepository) Add(webContentCategoryID,galleryName, createdAt, updatedAt string, tx *sql.Tx) (res string, err error) {
	statement := `insert into "galleries" ("web_content_category_id","gallery_name","created_at","updated_at") values($1,$2,$3,$4) returning "id"`
	err = tx.QueryRow(statement, webContentCategoryID,galleryName,datetime.StrParseToTime(createdAt,time.RFC3339),datetime.StrParseToTime(updatedAt,time.RFC3339)).Scan(&res)

	return res,err
}

func (GalleryRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "galleries" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3`
	_,err = tx.Exec(statement,datetime.StrParseToTime(updatedAt,time.RFC3339),datetime.StrParseToTime(deletedAt,time.RFC3339),ID)

	return err
}

func (repository GalleryRepository) CountBy(ID, column, value string) (res int, err error) {
	fmt.Print(ID)
	fmt.Println(value)
	if ID == "" {
		statement := `select count("id") from "galleries" where `+column+`=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,value).Scan(&res)
	}else{
		statement := `select count("id") from "galleries" where (`+column+`=$1 and "deleted_at" is null) and "id" <> $2`
		err = repository.DB.QueryRow(statement,value,ID).Scan(&res)
	}

	return res,err
}

