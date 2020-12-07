package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type GlobalInfoCategorySettingRepository struct {
	DB *sql.DB
}

func NewGlobalInfoCategorySettingRepository(DB *sql.DB) contracts.IGlobalInfoCategorySettingRepository {
	return &GlobalInfoCategorySettingRepository{DB: DB}
}

func (repository GlobalInfoCategorySettingRepository) Browse(globalInfoCategory, search, order, sort string, limit, offset int) (data []models.GlobalInfoCategorySetting, count int, err error) {
	var rows *sql.Rows
	if globalInfoCategory == "" {
		statement := `select gics.*,gic."name" from "global_info_category_settings" gics
                 inner join "global_info_categories" gic on gic."id"=gics."global_info_category_id" and gic."deleted_at" is null
                 where (lower(gic."name") like $1 or lower(gics."description") like $1) and gics."deleted_at" is null order by gic.` + order + ` ` + sort + ` limit $2 offset $3`
		rows, err = repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	} else {
		statement := `select gics.*,gic."name" from "global_info_category_settings" gics
                 inner join "global_info_categories" gic on gic."id"=gics."global_info_category_id" and gic."deleted_at" is null
                 where gic."slug"=$1 and gics."deleted_at" is null order by gic.` + order + ` ` + sort + ` limit $2 offset $3`
		rows, err = repository.DB.Query(statement, globalInfoCategory, limit, offset)
	}
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.GlobalInfoCategorySetting{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.GlobalInfoCategoryID,
			&dataTemp.Description,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.GlobalInfoCategoryName,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	if globalInfoCategory == "" {
		statement := `select count(gics."id") from "global_info_category_settings" gics
                 inner join "global_info_categories" gic on gic."id"=gics."global_info_category_id" and gic."deleted_at" is null
                 where (lower(gic."name") like $1 or lower(gics."description") like $1) and gics."deleted_at" is null`
		err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	} else {
		statement := `select count(gics."id") from "global_info_category_settings" gics
                 inner join "global_info_categories" gic on gic."id"=gics."global_info_category_id" and gic."deleted_at" is null
                 where gic."slug"=$1 and gics."deleted_at" is null`
		err = repository.DB.QueryRow(statement, globalInfoCategory).Scan(&count)
	}
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository GlobalInfoCategorySettingRepository) Browse2(globalInfoCategory, search, order, sort string, limit, offset int) (data []models.GlobalInfoCategorySetting, count int, err error) {
	var rows *sql.Rows

	filterQuery := ``
	filterAttr := []interface{}{}
	if globalInfoCategory == "" {
		filterQuery += `and (lower(gic."name") like $1 or lower(gics."description") like $1)`
		filterAttr = append(filterAttr, "%"+strings.ToLower(search)+"%")
	} else {
		filterQuery += `and gic."slug"=$1`
		filterAttr = append(filterAttr, globalInfoCategory)
	}
	paginationAttr := append(filterAttr, []interface{}{limit, offset}...)
	statement := `select gics.*,gic."name" from "global_info_category_settings" gics
                inner join "global_info_categories" gic on gic."id"=gics."global_info_category_id" and gic."deleted_at" is null
                 where gics."deleted_at" is null ` + filterQuery + ` order by gic.` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err = repository.DB.Query(statement, paginationAttr...)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.GlobalInfoCategorySetting{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.GlobalInfoCategoryID,
			&dataTemp.Description,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.GlobalInfoCategoryName,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count(gics."id") from "global_info_category_settings" gics
                 inner join "global_info_categories" gic on gic."id"=gics."global_info_category_id" and gic."deleted_at" is null
                 where gics."deleted_at" is null ` + filterQuery
	err = repository.DB.QueryRow(statement, filterAttr).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository GlobalInfoCategorySettingRepository) ReadBy(column, value string) (data models.GlobalInfoCategorySetting, err error) {
	statement := `select gics.*,gic."name" from "global_info_category_settings" gics
                 inner join "global_info_categories" gic on gic."id"=gics."global_info_category_id" and gic."deleted_at" is null
                 where gics.` + column + `=$1 and gics."deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.GlobalInfoCategoryID,
		&data.Description,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.GlobalInfoCategoryName,
	)
	fmt.Print(data)

	return data, err
}

func (repository GlobalInfoCategorySettingRepository) Edit(input viewmodel.GlobalInfoCategorySettingVm) (res string, err error) {
	statement := `update "global_info_category_settings" set "global_info_category_id"=$1, "description"=$2, "is_active"=$3, "updated_at"=$4 where "id"=$5 returning "id"`
	err = repository.DB.QueryRow(statement, input.GlobalInfoCategoryID, input.Description, input.IsActive, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID).Scan(&res)

	return res, err
}

func (repository GlobalInfoCategorySettingRepository) Add(input viewmodel.GlobalInfoCategorySettingVm) (res string, err error) {
	statement := `insert into "global_info_category_settings" ("global_info_category_id","description","is_active","created_at","updated_at") values($1,$2,$3,$4,$5) returning "id"`
	err = repository.DB.QueryRow(statement, input.GlobalInfoCategoryID, input.Description, input.IsActive, datetime.StrParseToTime(input.CreatedAt, time.RFC3339), datetime.StrParseToTime(input.UpdatedAt, time.RFC3339)).Scan(&res)

	return res, err
}

func (repository GlobalInfoCategorySettingRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "global_info_category_settings" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository GlobalInfoCategorySettingRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID != "" {
		statement := `select count("id") from "global_info_category_settings" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	} else {
		statement := `select count("id") from "global_info_category_settings" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	}

	return res, err
}
