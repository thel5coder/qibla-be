package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type OdooMasterPackageRepository struct {
	DB *sql.DB
}

func NewOdooMasterPackageRepository(DB *sql.DB) contracts.IOdooMasterPackageRepository {
	return &OdooMasterPackageRepository{DB: DB}
}

const (
	odooMasterPackageSelectStatement = `select mp."id",eppt."id",eppt."name",mp."name",mp."quota",mp."arrival_date",mp."return_date",mp."notes",mp."website_description",
                                        array_to_string(array_agg(DISTINCT h.id || ':' || hpt.id || ':' || coalesce(h.hotel_name,'') || ':' || coalesce(hpt.rating,'0')),','),
                                        array_to_string(array_agg(DISTINCT  m.id || ':' || mpt.id || ':' || coalesce(mpt.name,'')),','),
                                        array_to_string(array_agg(DISTINCT t.id || ':' || tpt.id || ':' || coalesce(tpt.name,'')),','),
                                        array_to_string(array_agg(DISTINCT al.id || ':' || ta.id || ':' || coalesce(ta.name,'')),','),
                                        array_to_string(array_agg(DISTINCT rr.id || ':' || coalesce(rr.room_name,'') || ':' || coalesce(rr.price_unit,0) || ':' || coalesce(rr.price_promo,0) || ':' || coalesce(r.no_of_person,0) || ':' 
                                       || coalesce(alc.name,'')),',')`
	odooMasterPackageJoinStatement = `inner join "equipment" e on e.id = mp.package_equipment_id
                                      inner join "product_product" ep on ep.id = e.product_id
    								  inner join "product_template" eppt on eppt.id = ep.product_tmpl_id

									  inner join "hotels" h on h.package_id = mp.id
									  inner join "product_product" hp on hp.id = h.hotel_id
									  inner join "product_template" hpt on hpt.id = hp.product_tmpl_id

									  inner join "transportation" t on t.package_id = mp.id
									  inner join "product_product" tp on tp.id = t.product_id
									  inner join "product_template" tpt on tpt.id = tp.product_tmpl_id

									  inner join "meals" m on m.package_id = mp.id
                                      inner join "product_product" mpp on mpp.id = m.product_id
                                      inner join "product_template" mpt on mpt.id = mpp.product_tmpl_id

                                      inner join "airlines" al on al.package_id = mp.id
                                      inner join "tour_airline" ta on ta.id = al.air_lines_id
									  inner join "room_rate" rr on rr.package_id = mp.id
									  inner join "room_room" r on r.id = rr.room_id
									  inner join "airlines_class" alc on alc.id = rr.airlines_class`
	odooMasterPackageGroupByStatement = `group by mp.id,eppt.id`
)

func (repository OdooMasterPackageRepository) scanRows(rows *sql.Rows) (res models.OdooMasterPackage, err error) {
	err = rows.Scan(&res.ID, &res.EquipmentPackageID, &res.EquipmentPackageName, &res.Name, &res.Quota, &res.DepartureDate, &res.ReturnDate, &res.Notes, &res.WebsiteDescription,
		&res.Hotels, &res.Meals, &res.Transportations, &res.Airlines, &res.RoomRates)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository OdooMasterPackageRepository) scanRow(row *sql.Row) (res models.OdooMasterPackage, err error) {
	err = row.Scan(&res.ID, &res.EquipmentPackageID, &res.EquipmentPackageName, &res.Name, &res.Quota, &res.DepartureDate, &res.ReturnDate, &res.Notes, &res.WebsiteDescription,
		&res.Hotels, &res.Meals, &res.Transportations, &res.Airlines, &res.RoomRates)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository OdooMasterPackageRepository) BrowseAll() (data []models.OdooMasterPackage, err error) {
	statement := odooMasterPackageSelectStatement + ` from "master_package" mp ` + odooMasterPackageJoinStatement + ` where mp.is_active=true ` + odooMasterPackageGroupByStatement + ` order by mp.id asc`
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

func (repository OdooMasterPackageRepository) ReadBy(column, value, operator string) (data models.OdooMasterPackage, err error) {
	statement := odooMasterPackageSelectStatement + ` from "master_package" mp ` + odooMasterPackageJoinStatement + ` where mp.is_active=true and ` + column + `` + operator + `$1`
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
