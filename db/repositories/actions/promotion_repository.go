package actions

import (
	"database/sql"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type PromotionRepository struct {
	DB *sql.DB
}

func NewPromotionRepository(DB *sql.DB) contracts.IPromotionRepository {
	return &PromotionRepository{DB: DB}
}

func (repository PromotionRepository) Browse(search, order, sort string, limit, offset int) (data []models.Promotion, count int, err error) {
	statement := `select p.*, pp."package_name" from "promotions" p
                 inner join "promotion_packages" pp on pp."id"=p."promotion_package_id"
                 where (lower(pp."package_name") like $1 or lower(p."package_promotion") like $1 or cast(p."price" as varchar) like $1 or lower("description") like $1) 
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
                 inner join "promotion_packages" pp on pp."id"=p."promotion_package_id"
                 where (lower(pp."package_name") like $1 or lower(p."package_promotion") like $1 or cast(p."price" as varchar) like $1 or lower("description") like $1) 
                 and p."deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository PromotionRepository) ReadBy(column, value string) (data models.Promotion, err error) {
	statement := `select p.*, pp."package_name" from "promotions" p 
                 inner join "promotion_packages" pp on pp."id"=p."promotion_package_id"
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

func (PromotionRepository) Edit(input viewmodel.PromotionVm, tx *sql.Tx) (res string, err error) {
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

func (PromotionRepository) Add(input viewmodel.PromotionVm, tx *sql.Tx) (res string, err error) {
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

func (PromotionRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (res string, err error) {
	statement := `update "promotions" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return res, err
}

func (repository PromotionRepository) CountBy(ID, promotionPackageID, column, value string) (res int, err error) {
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
