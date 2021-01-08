package actions

import (
	"database/sql"
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

func (repository SatisfactionCategoryRepository) scanRows(rows *sql.Rows) (res models.SatisfactionCategory, err error) {
	err = rows.Scan(&res.ID, &res.ParentID, &res.Slug, &res.Name, &res.IsActive, &res.Description, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository SatisfactionCategoryRepository) scanRow(row *sql.Row) (res models.SatisfactionCategory, err error) {
	err = row.Scan(&res.ID, &res.ParentID, &res.Slug, &res.Name, &res.IsActive, &res.Description, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository SatisfactionCategoryRepository) BrowseAllBy(filters map[string]interface{}, order, sort string) (data []models.SatisfactionCategory, err error) {
	filterStatement := ``

	if val, ok := filters["name"]; ok {
		filterStatement += ` and lower(sc.name) like '%` + val.(string) + `%'`
	}

	if val, ok := filters["description"]; ok {
		filterStatement += ` and lower(sc.description) like '%` + val.(string) + `%'`
	}

	if val, ok := filters["updated_at"]; ok {
		filterStatement += ` and cast(sc.updated_at as varchar)='` + val.(string) + `'`
	}

	statement := `WITH RECURSIVE satisfactions AS (
                  select id,parent_id,slug,name,is_active,description,created_at,updated_at
	              from satisfaction_categories
	              where parent_id is null
	              UNION 
		          SELECT sc.id,sc.parent_id,sc.slug,sc.name,sc.is_active,sc.description,sc.created_at,sc.updated_at
                  from satisfaction_categories sc inner join satisfactions s on s.id=sc.parent_id
				  where sc.deleted_at is null ` + filterStatement + `
                  ) select * from satisfactions order by ` + order + ` ` + sort
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

	return data, err
}

func (repository SatisfactionCategoryRepository) ReadBy(column, value string) (data []models.SatisfactionCategory, err error) {
	statement := `WITH RECURSIVE satisfactions AS (
                  select id,parent_id,slug,name,is_active,description,created_at,updated_at
	              from satisfaction_categories
	              where parent_id is null and ` + column + `=$1
	              UNION 
		          SELECT sc.id,sc.parent_id,sc.slug,sc.name,sc.is_active,sc.description,sc.created_at,sc.updated_at
                  from satisfaction_categories sc inner join satisfactions s on s.id=sc.parent_id
				  where sc.deleted_at is null
                  )select * from satisfactions`
	rows, err := repository.DB.Query(statement,value)
	if err != nil {
		return data, err
	}
	for rows.Next() {
		temp, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, temp)
	}

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

func (SatisfactionCategoryRepository) EditUpdatedAt(model models.SatisfactionCategory, tx *sql.Tx) (err error) {
	statement := `update satisfaction_categories set updated_at=$1 where id=$2`
	_, err = tx.Exec(statement, datetime.StrParseToTime(model.UpdatedAt, time.RFC3339), model.ID)
	if err != nil {
		return err
	}

	return nil
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
