package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"strconv"
	"strings"
)

type VideoContentDetailRepository struct {
	DB *sql.DB
}

func NewVideContentDetailRepository(DB *sql.DB) contracts.IVideoContentDetailRepository {
	return &VideoContentDetailRepository{DB: DB}
}

const (
	videoContentDetailSelectStatement = `select id,video_content_id,title,channel_title,embedded_url,thumbnails,description,published_at,created_at,updated_at`
)

func (repository VideoContentDetailRepository) scanRows(rows *sql.Rows) (res models.VideoContentDetails, err error) {
	err = rows.Scan(&res.ID, &res.VideoContentID, &res.Title, &res.ChannelTitle, &res.EmbeddedUrl, &res.Thumbnails, &res.Description, &res.PublishedAt, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository VideoContentDetailRepository) scanRow(row *sql.Row) (res models.VideoContentDetails, err error) {
	err = row.Scan(&res.ID, &res.VideoContentID, &res.Title, &res.ChannelTitle, &res.EmbeddedUrl, &res.Thumbnails, &res.Description, &res.PublishedAt, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository VideoContentDetailRepository) BrowseByVideoContentID(videoContentID, search, order, sort string, limit, offset int) (data []models.VideoContentDetails, count int, err error) {
	whereStatement := `where (lower(title) like $1 or lower(channel_title) like $1) and deleted_at is null`
	whereParams := []interface{}{"%" + strings.ToLower(search) + "%"}
	if videoContentID != "" {
		whereStatement += ` and video_content_id=$2`
		whereParams = append(whereParams, videoContentID)
	}

	statement := videoContentDetailSelectStatement + ` from video_content_details ` + whereStatement + ` order by ` + order + ` ` + sort +
		` limit ` + strconv.Itoa(limit) + ` offset ` + strconv.Itoa(offset)
	rows, err := repository.DB.Query(statement, whereParams...)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, temp)
	}

	statement = `select count(id) from video_content_details ` + whereStatement
	err = repository.DB.QueryRow(statement, whereParams...).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, nil
}

func (repository VideoContentDetailRepository) Read(ID string) (data models.VideoContentDetails, err error) {
	statement := videoContentDetailSelectStatement + ` from video_content_details where id=$1 and deleted_at is null`
	row := repository.DB.QueryRow(statement, ID)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository VideoContentDetailRepository) Add(model models.VideoContentDetails) (res string, err error) {
	statement := `insert into video_content_details (video_content_id,title,channel_title,embedded_url,thumbnails,description,published_at,created_at,updated_at) 
                  values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`
	err = repository.DB.QueryRow(statement, model.VideoContentID, model.Title, model.ChannelTitle, model.EmbeddedUrl, model.Thumbnails, model.Description, model.PublishedAt,
		model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository VideoContentDetailRepository) DeleteBy(column, value, operator string, model models.VideoContentDetails) (res string, err error) {
	statement := `update video_content_details set updated_at=$1, deleted_at=$2 where ` + column + `` + operator + `$3 returning id`
	err = repository.DB.QueryRow(statement, model.UpdatedAt, model.DeletedAt.Time, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository VideoContentDetailRepository) CountBy(column, value, operator string) (res int, err error) {
	statement := `select count (id) from video_content_details where ` + column + `` + operator + `$1 and deleted_at is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
