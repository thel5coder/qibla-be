package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/str"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type SatisfactionCategoryRepository struct {
	DB *sql.DB
}

func NewSatisfactionCategoryRepository(db *sql.DB) contracts.ISatisfactionCategoryRepository {
	return &SatisfactionCategoryRepository{DB: db}
}

const(
	satisfactionCategorySelectStatement = ``
)

func (repository SatisfactionCategoryRepository) BrowseAllBy(column, value string) (data []models.SatisfactionCategory, err error) {
	var rows *sql.Rows
	if value == "" && column == "parent_id" {
		statement := `select * from "satisfaction_categories" where "parent_id" is null and "deleted_at" is null`
		rows, err = repository.DB.Query(statement)
		fmt.Println(rows)
	} else {
		statement := `select * from "satisfaction_categories" where ` + column + `=$1 and "deleted_at" is null`
		rows, err = repository.DB.Query(statement, value)
	}
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.SatisfactionCategory{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.ParentID,
			&dataTemp.Slug,
			&dataTemp.Name,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.Description,
		)
		if err != nil {
			return data, err
		}
		data = append(data, dataTemp)
	}

	return data, err
}

func (repository SatisfactionCategoryRepository) ReadBy(column, value string) (data models.SatisfactionCategory, err error) {
	statement := `select * from "satisfaction_categories" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.ParentID,
		&data.Slug,
		&data.Name,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.Description,
	)

	return data, err
}

func (SatisfactionCategoryRepository) Edit(input viewmodel.SatisfactionCategoryVm, tx *sql.Tx) (err error) {
	statement := `update "satisfaction_categories" set "parent_id"=$1, "slug"=$2, "name"=$3, "description"=$4, "is_active"=$5, "updated_at"=$6 where "id"=$7`
	_, err = tx.Exec(
		statement,
		str.EmptyString(input.ParentID),
		input.Slug,
		input.Name,
		str.EmptyString(input.Description),
		input.IsActive,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	)

	return err
}

func (SatisfactionCategoryRepository) Add(input viewmodel.SatisfactionCategoryVm, tx *sql.Tx) (err error) {
	statement := `insert into "satisfaction_categories" ("parent_id","slug","name","description","is_active","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7)`
	_, err = tx.Exec(
		statement,
		str.EmptyString(input.ParentID),
		input.Slug,
		input.Name,
		str.EmptyString(input.Description),
		input.IsActive,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	)

	return err
}

func (SatisfactionCategoryRepository) DeleteBy(column, value, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "satisfaction_categories" set "updated_at"=$1, "deleted_at"=$2 where ` + column + `=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), value)

	return err
}

func (repository SatisfactionCategoryRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "satisfaction_categories" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "satisfaction_categories" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}

func (repository SatisfactionCategoryRepository) CountByParentID(parentID, ID, slug string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "satisfaction_categories" where parent_id=$1 and "slug"=$2 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, parentID, slug).Scan(&res)
	} else {
		statement := `select count("id") from "satisfaction_categories" where (parent_id=$1 and "slug"=$2 and "deleted_at" is null) and "id"<>$3`
		err = repository.DB.QueryRow(statement, parentID, slug, ID).Scan(&res)
	}

	return res, err
}
