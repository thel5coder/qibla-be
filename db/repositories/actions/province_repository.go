package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type ProvinceRepository struct {
	DB *sql.DB
}

func NewProvinceRepository(DB *sql.DB) contracts.IProvinceRepository {
	return &ProvinceRepository{DB: DB}
}

const (
	provinceSelectStatement = `select id,name`
)

func (repository ProvinceRepository) scanRows(rows *sql.Rows) (res models.Province, err error) {
	err = rows.Scan(&res.ID, &res.Name)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository ProvinceRepository) BrowseAll() (data []models.Province, err error) {
	statement := provinceSelectStatement + ` from "master_provinces" order by name asc`
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

	return data, nil
}
