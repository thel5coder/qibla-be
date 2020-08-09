package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type PromotionPackageRepository struct {
	DB *sql.DB
}

func NewPromotionPackageRepository(DB *sql.DB) contracts.IPackagePromotionRepository {
	return PromotionPackageRepository{DB: DB}
}

func (repository PromotionPackageRepository) Browse(search, order, sort string, limit, offset int) (data []models.PromotionPackage, count int, err error) {
	statement := `select * from "promotion_packages" 
                  where lower("package_name") like $1 and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		dataTemp := models.PromotionPackage{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Slug,
			&dataTemp.PackageName,
			&dataTemp.IsActive,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, count, err
		}

		data = append(data, dataTemp)
	}

	statement = `select count("id") from "promotion_packages" 
                  where lower("package_name") like $1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

func (repository PromotionPackageRepository) ReadBy(column, value string) (data models.PromotionPackage, err error) {
	statement := `select * from "promotion_packages" where ` + column + `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(
		&data.ID,
		&data.Slug,
		&data.PackageName,
		&data.IsActive,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

func (repository PromotionPackageRepository) Edit(input viewmodel.PromotionPackageVm) (res string, err error) {
	statement := `update "promotion_packages" set "package_name"=$1, "slug"=$2, "updated_at"=$3 where "id"=$4 returning "id"`
	err = repository.DB.QueryRow(statement, input.PackageName, input.Slug, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID).Scan(&res)

	return res, err
}

func (repository PromotionPackageRepository) Add(input viewmodel.PromotionPackageVm) (res string, err error) {
	statement := `insert into "promotion_packages" ("package_name","slug","created_at","updated_at") values($1,$2,$3,$4) returning "id"`
	err = repository.DB.QueryRow(
		statement,
		input.PackageName,
		input.Slug,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

func (repository PromotionPackageRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "promotion_packages" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (repository PromotionPackageRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `select count("id") from "promotion_packages" where ` + column + `=$1 and "deleted_at" is null`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `select count("id") from "promotion_packages" where (` + column + `=$1 and "deleted_at" is null) and "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
