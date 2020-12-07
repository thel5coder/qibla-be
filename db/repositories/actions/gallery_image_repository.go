package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"time"
)

type GalleryImageRepository struct {
	DB *sql.DB
}

func NewGalleryImageRepository(DB *sql.DB) contracts.IGalleryImageRepository {
	return &GalleryImageRepository{DB: DB}
}

func (repository GalleryImageRepository) Browse(galleryID string) (data []models.GalleryImages, err error) {
	statement := `select gi.*,file."path" from "gallery_images" gi
                 inner join "files" file on file."id"=gi."file_id"
                 where gi."gallery_id"=$1 and gi."deleted_at" is null`
	rows, err := repository.DB.Query(statement, galleryID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.GalleryImages{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.GalleryID,
			&dataTemp.FileID,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.Path,
		)
		if err != nil {
			return data, err
		}
		fmt.Print(dataTemp)

		data = append(data, dataTemp)
	}

	return data, err
}

func (repository GalleryImageRepository) ReadBy(column, value string) (data models.GalleryImages, err error) {
	statement := `select gi.*,file."path" from "gallery_images" gi
                 inner join "files" file on file."id"=gi."file_id" and "deleted_at" is null
                 where gi.` + column + `=$1 and gi."deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.GalleryID,
		&data.FileID,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.Path,
	)

	return data,err
}

func (GalleryImageRepository) Edit(ID, fileID, updatedAt string, tx *sql.Tx) (err error) {
	statement := `update "gallery_images" set "file_id"=$1, "updated_at"=$2 where "id"=$3`
	_, err = tx.Exec(statement, fileID, datetime.StrParseToTime(updatedAt, time.RFC3339), ID)

	return err
}

func (GalleryImageRepository) Add(galleryID, fileID, createdAt, updatedAt string, tx *sql.Tx) (err error) {
	statement := `insert into "gallery_images" ("gallery_id","file_id","created_at","updated_at") values($1,$2,$3,$4)`
	_, err = tx.Exec(statement, galleryID, fileID, datetime.StrParseToTime(createdAt, time.RFC3339), datetime.StrParseToTime(updatedAt, time.RFC3339))

	return err
}

func (GalleryImageRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "gallery_images" set "updated_at"=$1,"deleted_at"=$2 where "id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return err
}

func (GalleryImageRepository) DeleteByGalleryID(galleryID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "gallery_images" set "updated_at"=$1,"deleted_at"=$2 where "gallery_id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), galleryID)

	return err
}

func (repository GalleryImageRepository) CountBy(ID, galleryID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "gallery_images" where ` + column + `=$1 and "deleted_at" is null and "gallery_id"=$2`
		err = repository.DB.QueryRow(statement, value, galleryID).Scan(&res)
	} else {
		statement := `select count("id") from "gallery_images" where (` + column + `=$1 and "deleted_at" is null and "gallery_id"=$2) and "id" <> $3`
		err = repository.DB.QueryRow(statement, value, galleryID, ID).Scan(&res)
	}

	return res, err
}
