package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"time"
)

type TourPackageRepository struct {
	DB *sql.DB
}

func NewTourPackageRepository(DB *sql.DB) contracts.ITourPackageRepository {
	return &TourPackageRepository{DB: DB}
}

const (
	tourPackageSelectStatement = `select tp.id,tp.odoo_package_id,tp.name,tp.departure_date,tp.return_date,tp.package_program,tp.odoo_package_program_id,tp.program_day,tp.description,tp.notes,
                                  tp.created_at,tp.updated_at,array_to_string(array_agg(tph.id || ':' || tph.odoo_product_template_id || ':' || tph.name || ':' || coalesce(tph.rating,0) 
                                  || ':' || coalesce(tph.location,null)),','), array_to_string(array_agg(tpm.id || ':' || tpm.odoo_product_template_id || ':' || tpm.name),','), 
                                  array_to_string(array_agg(tpb.id || ':' || tpb.odoo_product_template_id || ':' || tpb.name),','), array_to_string(array_agg(tpa.id || ':' || 
                                  tpa.odoo_product_template_id || ':' || tpa.name),','),array_to_string(array_agg(tpp.id || ':' || tpp.room_type || ':' || tpp.room_capacity || ':' 
                                  || tpp.price || ':' || coalesce(tpp.price_promo,null) || ':' || tpp.airline_class),',')`
	tourPackageJoinStatement = `inner join "tour_package_hotels" tph on tph.tour_package_id = tp.id 
                                inner join "tour_package_meals" tpm tpm.tour_package_id = tp.id
                                inner join "tour_package_busses" tpb tpb.tour_package_id = tp.id
                                inner join "tour_package_airlines" tpa tpa.tour_package_id = tp.id
                                inner join "tour_package_prices" tpp tpp.tour_package_id = tp.id`
	tourPackageGroupByStatement = `group by tp.id`
)

func (repository TourPackageRepository) scanRows(rows *sql.Rows) (res models.TourPackage, err error) {


	return res, nil
}

func (repository TourPackageRepository) scanRow(row *sql.Row) (res models.TourPackage, err error) {


	return res, nil
}

func (repository TourPackageRepository) BrowseAllBy(column, value, operator string) (data []models.TourPackage, err error) {
	statement := tourPackageSelectStatement + ` from "tour_packages" from tp ` + tourPackageJoinStatement + ` where ` + column + `` + operator + `$1 ` + tourPackageGroupByStatement
	rows, err := repository.DB.Query(statement, value)
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

func (repository TourPackageRepository) ReadBy(column, value, operator string) (data models.TourPackage, err error) {
	statement := tourPackageSelectStatement + ` from "tour_packages" from tp ` + tourPackageJoinStatement + ` where ` + column + `` + operator + `$1 ` + tourPackageGroupByStatement
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (TourPackageRepository) Edit(input models.TourPackage, tx *sql.Tx) (err error) {


	return nil
}

func (TourPackageRepository) Add(input models.TourPackage, tx *sql.Tx) (res string, err error) {


	return res, nil
}

func (repository TourPackageRepository) DeleteBy(column, value, operator string, updatedAt, deletedAt time.Time, tx *sql.Tx) (err error) {
	statement := `update "tour_packages" set "updated_at"=$1, "deleted_at"=$2 where ` + column + `` + operator + `$3`
	_, err = tx.Exec(statement, updatedAt, deletedAt, value)
	if err != nil {
		return err
	}

	return nil
}
