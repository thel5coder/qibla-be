package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type CityRepository struct {
	DB *sql.DB
}

func NewCityRepository(DB *sql.DB) contracts.ICityRepository {
	return CityRepository{DB: DB}
}

const (
	citySelectStatement = `select id,name,province_id`
)

func (repository CityRepository) scanRows(rows *sql.Rows) (res models.City, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.ProvinceID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository CityRepository) BrowseAllByProvince(provinceID string) (data []models.City, err error) {
	statement := citySelectStatement + ` from "master_cities" where province_id=$1 order by name asc`
	rows, err := repository.DB.Query(statement, provinceID)
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

	return data, nil
}
