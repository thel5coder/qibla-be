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

type SettingPromotionRepository struct {
	DB *sql.DB
}

func NewSettingPromotionRepository(DB *sql.DB) contracts.ISettingPromotionRepository {
	return &SettingPromotionRepository{DB: DB}
}

const (
	promotionSelectStatement = `select p.id,pp.id,pp.name,pp.slug,p.package_promotion,p.start_date,p.end_date,p.price,p.description,
                               array_to_string(array_agg(plat.id || ':' || plat.platform || ':'),','),
                               array_to_string(array_agg(pos.id || ':' || pos.promotion_platform_id || ':' || pos.position),',')`
	promotionJoinStatement = `inner join "master_promotions" pp on pp."id"=p."promotion_package_id"
                              inner join "promotion_platforms" plat on plat.promotion_id=p.id
                              inner join "promotion_positions" pos on pos.promotion_platform_id=plat.id`
	promotionGroupByStatement = `GROUP BY p."id",pp."id"`
)

func (repository SettingPromotionRepository) scanRows(rows *sql.Rows) (res models.Promotion,err error){
	err = rows.Scan(&res.ID,&res.PromotionPackageID,&res.PackageName,&res.PackagePromotionSlug,&res.PackagePromotion,&res.StartDate,&res.EndDate,&res.Price,&res.Description,&res.Platform,
		&res.Position)
	if err != nil {
		return res,err
	}

	return res,nil
}

func (repository SettingPromotionRepository) scanRow(row *sql.Row) (res models.Promotion,err error){
	err = row.Scan(&res.ID,&res.PromotionPackageID,&res.PackagePromotion,&res.PackagePromotionSlug,&res.PackageName,&res.StartDate,&res.EndDate,&res.Price,&res.Description,&res.Platform,
		&res.Position)
	if err != nil {
		return res,err
	}

	return res,nil
}

func (repository SettingPromotionRepository) Browse(search, order, sort string, limit, offset int) (data []models.Promotion, count int, err error) {
	statement := `select p.*, pp."name" from "promotions" p
                 inner join "master_promotions" pp on pp."id"=p."promotion_package_id"
                 where (lower(pp."name") like $1 or lower(p."package_promotion") like $1 or cast(p."price" as varchar) like $1 or lower("description") like $1) 
                 and p."deleted_at" is null 
                 order by ` + order + " " + sort + " limit $2 offset $3"
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.Promotion{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.PromotionPackageID,
			&dataTemp.PackagePromotion,
			&dataTemp.StartDate,
			&dataTemp.EndDate,
			&dataTemp.Price,
			&dataTemp.Description,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
			&dataTemp.IsActive,
			&dataTemp.PackageName,
		)
		if err != nil {
			return data, count, err
		}
		data = append(data, dataTemp)
	}

	statement = `select count(p."id") from "promotions" p 
                 inner join master_promotions pp on pp."id"=p."promotion_package_id"
                 where (lower(pp.name) like $1 or lower(p."package_promotion") like $1 or cast(p."price" as varchar) like $1 or lower("description") like $1) 
                 and p."deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository SettingPromotionRepository) BrowseAll(filters map[string]interface{}) (data []models.Promotion, err error) {
	var filterStatement string
	if val,ok := filters["position"];ok {
		filterStatement += ` and lower(pos.position) = '`+val.(string)+`'`
	}

	if val,ok := filters["platform"];ok {
		platforms := strings.Split(val.(string),",")
		if len(platforms) > 1 {
			filterStatement += ` and (lower(plat.platform)='`+platforms[0]+`' or lower(plat.platform)='`+platforms[1]+`')`
		}else{
			filterStatement += ` and lower(plat.platform)='`+platforms[0]+`'`
		}

	}

	if val,ok := filters["startDate"];ok {
		filterStatement += ` and p.start_date < '`+val.(string)+`'`
	}

	if val,ok := filters["endDate"];ok {
		filterStatement += ` and p.end_date > '`+val.(string)+`'`
	}

	if val,ok := filters["type"];ok {
		filterStatement += ` and lower(p.package_promotion) = '`+val.(string)+`'`
	}

	statement := promotionSelectStatement +` from "promotions" p `+promotionJoinStatement+` where p."deleted_at" is null `+filterStatement+` `+promotionGroupByStatement
	rows,err := repository.DB.Query(statement)
	if err != nil {
		return data,err
	}
	for rows.Next() {
		temp,err := repository.scanRows(rows)
		if err != nil {
			return data,err
		}
		data = append(data,temp)
	}

	return data,err
}

func (repository SettingPromotionRepository) ReadBy(column, value string) (data models.Promotion, err error) {
	statement := `select p.*, pp."name" from "promotions" p 
                 inner join "master_promotions" pp on pp."id"=p."promotion_package_id"
                 where ` + column + `=$1
                 and p."deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.PromotionPackageID,
		&data.PackagePromotion,
		&data.StartDate,
		&data.EndDate,
		&data.Price,
		&data.Description,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
		&data.IsActive,
		&data.PackageName,
	)

	return data, err
}

func (SettingPromotionRepository) Edit(input viewmodel.PromotionTodayVm, tx *sql.Tx) (res string, err error) {
	statement := `update "promotions" set "promotion_package_id"=$1, "package_promotion"=$2, "start_date"=$3, "end_date"=$4, "price"=$5, "description"=$6, "updated_at"=$7, "is_active"=$8
                 where "id"=$9 returning "id"`
	_, err = tx.Exec(
		statement,
		input.PromotionPackageID,
		input.PackagePromotion,
		datetime.StrParseToTime(input.StartDate, "2006-01-02"),
		datetime.StrParseToTime(input.EndDate, "2006-01-02"),
		input.Price,
		input.Description,
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.IsActive,
		input.ID,
	)

	return res, err
}

func (SettingPromotionRepository) Add(input viewmodel.PromotionTodayVm, tx *sql.Tx) (res string, err error) {
	statement := `insert into "promotions" ("promotion_package_id","package_promotion","start_date","end_date","price","description","created_at","updated_at","is_active")
                 values($1,$2,$3,$4,$5,$6,$7,$8,$9) returning "id"`
	err = tx.QueryRow(
		statement,
		input.PromotionPackageID,
		input.PackagePromotion,
		datetime.StrParseToTime(input.StartDate, "2006-01-02"),
		datetime.StrParseToTime(input.EndDate, "2006-01-02"),
		input.Price,
		input.Description,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
		input.IsActive,
	).Scan(&res)

	return res, err
}

func (SettingPromotionRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (res string, err error) {
	statement := `update "promotions" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return res, err
}

func (repository SettingPromotionRepository) CountBy(ID, promotionPackageID, column, value string) (res int, err error) {
	if ID == "" {
		if promotionPackageID == ""{
			statement := `select count(p."id") from "promotions" p 
                 inner join "promotion_packages" pp on pp."id"=p."promotion_package_id"
                 where ` + column + `=$1 and p."deleted_at" is null`
			err = repository.DB.QueryRow(statement, value).Scan(&res)
			fmt.Print("ini")
		}else{
			fmt.Println(promotionPackageID)
			statement := `select count(p."id") from "promotions" p 
                 inner join "promotion_packages" pp on pp."id"=p."promotion_package_id"
                 where ` + column + `=$1 and p."deleted_at" is null and p."promotion_package_id"=$2`
			err = repository.DB.QueryRow(statement, value, promotionPackageID).Scan(&res)
		}
	} else {
		if promotionPackageID == ""{
			statement := `select count(p."id") from "promotions" p 
                 inner join "promotion_packages" pp on pp."id"=p."promotion_package_id"
                 where (` + column + `=$1 and p."deleted_at" is null) and p."id"<>$3`
			err = repository.DB.QueryRow(statement, value, promotionPackageID, ID).Scan(&res)
			fmt.Print("inu")
		}else{
			statement := `select count(p."id") from "promotions" p 
                 inner join "promotion_packages" pp on pp."id"= p."promotion_package_id"
                 where (` + column + `=$1 and p."deleted_at" is null and p."promotion_package_id"=$2) and p."id"<>$3`
			fmt.Println("ono")
			err = repository.DB.QueryRow(statement, value, promotionPackageID, ID).Scan(&res)
		}
	}

	return res, err
}
