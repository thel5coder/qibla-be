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
	err = rows.Scan(&res.ID, &res.OdooPackageID, &res.Name, &res.DepartureDate, &res.ReturnDate, &res.PackageProgram, &res.OdooPackageProgramID, &res.ProgramDay, &res.Description, &res.Notes,
		&res.CreatedAt, &res.UpdatedAt, &res.Hotels, &res.Meals, &res.Busses, &res.Airlines, &res.Prices)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository TourPackageRepository) scanRow(row *sql.Row) (res models.TourPackage, err error) {
	err = row.Scan(&res.ID, &res.OdooPackageID, &res.Name, &res.DepartureDate, &res.ReturnDate, &res.PackageProgram, &res.OdooPackageProgramID, &res.ProgramDay, &res.Description, &res.Notes,
		&res.CreatedAt, &res.UpdatedAt, &res.Hotels, &res.Meals, &res.Busses, &res.Airlines, &res.Prices)
	if err != nil {
		return res, err
	}

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
	statement := `update "tour_packages" set name=$1, departure_date=$2, return_date=$3, description=$4, odoo_package_program_id=$5, package_program=$6, program_day=$7, notes=$8, updated_at=$9 where id=$10`
	_, err = tx.Exec(statement, input.Name, input.DepartureDate, input.ReturnDate, input.Description, input.OdooPackageProgramID, input.PackageProgram, input.ProgramDay, input.Notes, input.UpdatedAt, input.ID)
	if err != nil {
		return err
	}

	return nil
}

func (TourPackageRepository) Add(input models.TourPackage, tx *sql.Tx) (res string, err error) {
	statement := `insert into tour_packages (odoo_package_id,name,departure_date,return_date,description,partner_id,odoo_package_program_id,package_program,program_day,notes,created_at,
                  updated_at) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) returning "id"`
	err = tx.QueryRow(statement, input.OdooPackageID, input.Name, input.DepartureDate, input.ReturnDate, input.Description, input.PartnerID, input.OdooPackageProgramID, input.PackageProgram,
		input.ProgramDay, input.Notes, input.CreatedAt, input.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

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
