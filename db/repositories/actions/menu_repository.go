package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/str"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type MenuRepository struct {
	DB *sql.DB
}

func NewMenuRepository(DB *sql.DB) contracts.IMenuRepository {
	menuSelectStatementParams = []interface{}{}

	return &MenuRepository{DB: DB}
}

const (
	menuSelect  = `select m."id",m."menu_id",m."name",m."url",m."parent_id",m."is_active",m."created_at",m."updated_at",array_to_string(array_agg(mp."id" || ':' || mp."permission"),',')`
	menuJoin    = `left join "menu_permissions" mp on mp."menu_id"=m."id"`
	menuGroupBy = `group by m."id"`
)

var (
	menuSelectStatementParams = []interface{}{}
	menuWhereStatement        = `where (m."menu_id" like $1 or lower(m."name") like $1 or lower(m."url") like $1 or cast(m."updated_at" as varchar) like $1) and m."deleted_at" is null`
)

func (repository MenuRepository) scanRow(row *sql.Row) (res models.Menu, err error) {
	err = row.Scan(&res.ID, &res.MenuID, &res.Name, &res.Url, &res.ParentID, &res.IsActive, &res.CreatedAt, &res.UpdatedAt, &res.Permissions)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository MenuRepository) scanRows(rows *sql.Rows) (res models.Menu, err error) {
	err = rows.Scan(&res.ID, &res.MenuID, &res.Name, &res.Url, &res.ParentID, &res.IsActive, &res.CreatedAt, &res.UpdatedAt, &res.Permissions)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository MenuRepository) Browse(parentID, search, order, sort string, limit, offset int) (data []models.Menu, count int, err error) {
	menuSelectStatementParams = append(menuSelectStatementParams, []interface{}{"%" + strings.ToLower(search) + "%", limit, offset}...)
	if parentID != "" {
		menuSelectStatementParams = append(menuSelectStatementParams, parentID)
		menuWhereStatement += ` and m."parent_id"=$4`
	} else {
		menuWhereStatement += ` and m."parent_id" is null`
	}
	statement := menuSelect + ` from "menus" m ` + menuJoin + ` ` + menuWhereStatement + ` ` + menuGroupBy + ` order by m.` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, menuSelectStatementParams...)
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

	statement = `select count(distinct m."id") from "menus" m ` + menuJoin + ` ` + menuWhereStatement
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository MenuRepository) BrowseAllBy(column, value, operator string,isActive bool) (data []models.Menu, err error) {
	menuSelectStatementParams = []interface{}{}
	whereStatement := `where ` + column + ` is null and m."deleted_at" is null`
	if value != "" {
		menuSelectStatementParams = []interface{}{value}
		whereStatement = `where ` + column + `` + operator + `$1 and m."deleted_at" is null`
	}

	if isActive == true {
		whereStatement +=` and is_active=true`
	}

	statement := menuSelect + ` from "menus" m ` + menuJoin + ` ` + whereStatement + ` ` + menuGroupBy
	rows, err := repository.DB.Query(statement, menuSelectStatementParams...)
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

func (repository MenuRepository) ReadBy(column, value, operator string) (data models.Menu, err error) {
	menuSelectStatementParams = []interface{}{}
	whereStatement := `where ` + column + ` is null and m."deleted_at" is null`
	if value != "" {
		menuSelectStatementParams = []interface{}{value}
		whereStatement = `where ` + column + `` + operator + `$1 and m."deleted_at" is null`
	}
	statement := menuSelect + ` from "menus" m ` + menuJoin + ` ` + whereStatement + ` ` + menuGroupBy
	row := repository.DB.QueryRow(statement, menuSelectStatementParams...)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (repository MenuRepository) Edit(input viewmodel.MenuVm, tx *sql.Tx) (err error) {
	statement := `update "menus" set "name"=$1, "url"=$2, "is_active"=$3, "updated_at"=$4 where "id"=$5 and "deleted_at" is null returning "id"`
	_, err = tx.Exec(
		statement,
		input.Name,
		input.Url,
		input.IsActive,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.ID,
	)

	return err
}

func (repository MenuRepository) Add(input viewmodel.MenuVm, tx *sql.Tx) (res string, err error) {
	statement := `insert into "menus" ("menu_id","name","url","parent_id","is_active","created_at","updated_at") values($1,$2,$3,$4,$5,$6,$7) returning "id"`
	err = tx.QueryRow(
		statement,
		input.MenuID,
		input.Name,
		input.Url,
		str.EmptyString(input.ParentID),
		input.IsActive,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (repository MenuRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "menus" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return err
}

func (repository MenuRepository) DeleteChild(parentID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "menus" set "updated_at"=$1, "deleted_at"=$2 where "parent_id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), parentID)

	return err
}

func (repository MenuRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID != "" {
		statement := `select count("id") from "menus" where ` + column + `=$1 and "deleted_at" is null and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	} else {
		if value == "" {
			statement := `select count("id") from "menus" where ` + column + ` is null and "deleted_at" is null`
			err = repository.DB.QueryRow(statement).Scan(&res)
		} else {
			statement := `select count("id") from "menus" where ` + column + `=$1 and "deleted_at" is null`
			err = repository.DB.QueryRow(statement, value).Scan(&res)
		}
	}

	return res, err
}

//query count by pk
func (repository MenuRepository) CountByPk(ID string) (res int, err error) {
	statement := `select count("id") from "menus" where "id"=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, ID).Scan(&res)

	return res, err
}
