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

type TestimonialRepository struct {
	DB *sql.DB
}

func NewTestimonialRepository(DB *sql.DB) contracts.ITestimonialRepository {
	return &TestimonialRepository{DB: DB}
}

func (repository TestimonialRepository) Browse(search, order, sort string, limit, offset int) (data []models.Testimonial, count int, err error) {
	statement := `select t.*, file."path" from "testimonials" t
                 left join "files" file on file."id"=t."file_id"
                where (lower(t."customer_name") like $1 or lower(t."job_position") like $1 or lower(t."testimony") like $1)
                and t."deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.Testimonial{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.WebContentCategoryID,
			&dataTemp.FileID,
			&dataTemp.CustomerName,
			&dataTemp.JobPosition,
			&dataTemp.Testimony,
			&dataTemp.Rating,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.Path,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count(t."id") from "testimonials" t
                 left join "files" file on file."id"=t."file_id"
                where (lower(t."customer_name") like $1 or lower(t."job_position") like $1 or lower(t."testimony") like $1)
                and t."deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository TestimonialRepository) ReadBy(column, value string) (data models.Testimonial, err error) {
	statement := `select t.*,file."path" from "testimonials" t
                 left join "files" file on file."id"=t."file_id"
                where t.` + column + `=$1 and t."deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.WebContentCategoryID,
		&data.FileID,
		&data.CustomerName,
		&data.JobPosition,
		&data.Testimony,
		&data.Rating,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.Path,
	)

	return data, err
}

func (repository TestimonialRepository) Edit(input viewmodel.TestimonialVm) (res string, err error) {
	statement := `update "testimonials" set "file_id"=$1, "customer_name"=$2, "job_position"=$3, "testimony"=$4, "rating"=$5, "is_active"=$6, "updated_at"=$7 where "id"=$8 returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.FileID,
		input.CustomerName,
		input.JobPosition,
		input.Testimony,
		input.Rating,
		input.IsActive,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	).Scan(&res)

	return res, err
}

func (repository TestimonialRepository) Add(input viewmodel.TestimonialVm) (res string, err error) {
	statement := `insert into "testimonials" ("web_content_category_id","file_id","customer_name","job_position","testimony","rating","is_active","created_at","updated_at")
                 values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.WebContentCategoryID,
		input.FileID,
		input.CustomerName,
		input.JobPosition,
		input.Testimony,
		input.Rating,
		input.IsActive,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (repository TestimonialRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "testimonials" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository TestimonialRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "testimonials" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "testimonials" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
