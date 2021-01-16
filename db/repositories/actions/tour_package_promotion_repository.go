package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type TourPackagePromotionRepository struct {
	DB *sql.DB
}

func NewTourPackagePromotionRepository(DB *sql.DB) contracts.ITourPackagePromotionRepository {
	return &TourPackagePromotionRepository{DB: DB}
}

const (
	tourPackagePromotionSelectStatement = `select tpp.id,tpp.tour_package_id,tpp.created_at,tpp.updated_at,tpp.deleted_at,tp.odoo_package_id,tp.name,tp.departure_date,tp.return_date,tp.description,tp.created_at,tp.updated_at,tp.deleted_at,
                                           tp.package_type,tp.program_day,tp.notes,tp.image,tp.departure_airport,tp.destination_airport,tp.return_departure_airport, tp.return_destination_airport,tp.quota,
                                           c.id,c.travel_agent_name,c.branch_name,c.area_code,c.phone_number,
                                           array_to_string(array_agg(tph.id || ':' || tph.name || ':' || tph.rating || ':' || coalesce(tph.location,'')),','),
                                           array_to_string(array_agg(tpr.id || ':' || tpr.room_type ||':'|| tpr.room_capacity ||':'|| tpr.price ||':'|| tpr.promo_price ||':'|| tpr.airline_class),',')`
	tourPackagePromotionJoinStatement = `inner join promotion_purchases pp on pp.tour_package_promotion_id = tpp.id and pp.deleted_at is null
                                         inner join tour_packages tp on tp.id = tpp.tour_package_id and tp.deleted_at is null
                                         inner join partners p on p.id = tp.partner_id and p.deleted_at is null
                                         inner join contacts c on c.id = p.contact_id and c.deleted_at is null
                                         inner join tour_package_hotels tph on tph.tour_package_id = tp.id and tph.deleted_at is null
                                         inner join tour_package_prices tpr on tpr.tour_package_id = tp.id and tpr.deleted_at is null`
	tourPackagePromotionWhereStatement   = "where tpp.deleted_at is null"
	tourPackagePromotionGroupByStatement = `group by tpp.id, pp.id, tp.id, p.id, c.id`
)

func (repository TourPackagePromotionRepository) scanRows(rows *sql.Rows) (res models.TourPackagePromotion, err error) {
	err = rows.Scan(&res.ID, &res.TourPackageID, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.TourPackage.OdooPackageID, &res.TourPackage.Name, &res.TourPackage.DepartureDate, &res.TourPackage.ReturnDate, &res.TourPackage.Description, &res.TourPackage.CreatedAt, &res.TourPackage.UpdatedAt, &res.TourPackage.DeletedAt,
		&res.TourPackage.PackageType, &res.TourPackage.ProgramDay, &res.TourPackage.Notes, &res.TourPackage.Image, &res.TourPackage.DepartureAirport, &res.TourPackage.DestinationAirport, &res.TourPackage.ReturnDepartureAirport, &res.TourPackage.ReturnDestinationAirport, &res.TourPackage.Quota,
		&res.TravelAgentID, &res.TravelAgentName, &res.Branch, &res.AreaCode, &res.Phone, &res.Hotels, &res.RoomRates)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository TourPackagePromotionRepository) scanRow(row *sql.Row) (res models.TourPackagePromotion, err error) {
	err = row.Scan(&res.ID, &res.TourPackageID, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.TourPackage.OdooPackageID, &res.TourPackage.Name, &res.TourPackage.DepartureDate, &res.TourPackage.ReturnDate, &res.TourPackage.Description, &res.TourPackage.CreatedAt, &res.TourPackage.UpdatedAt, &res.TourPackage.DeletedAt,
		&res.TourPackage.PackageType, &res.TourPackage.ProgramDay, &res.TourPackage.Notes, &res.TourPackage.Image, &res.TourPackage.DepartureAirport, &res.TourPackage.DestinationAirport, &res.TourPackage.ReturnDepartureAirport, &res.TourPackage.ReturnDestinationAirport, &res.TourPackage.Quota,
		&res.TravelAgentID, &res.TravelAgentName, &res.Branch, &res.AreaCode, &res.Phone, &res.Hotels, &res.RoomRates)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repository TourPackagePromotionRepository) Browse(filters map[string]interface{}, order, sort string, limit, offset int) (data []models.TourPackagePromotion, count int, err error) {
	statement := tourPackagePromotionSelectStatement + ` from tour_package_promotions tpp ` + tourPackagePromotionJoinStatement + ` ` + tourPackagePromotionWhereStatement + ` ` + tourPackagePromotionGroupByStatement + ` order by tp.created_at asc limit $1 offset $2`
	rows, err := repository.DB.Query(statement, limit, offset)
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

	statement = `select count(distinct tpp.id ) from tour_package_promotions tpp ` + tourPackagePromotionJoinStatement + ` ` + tourPackagePromotionWhereStatement + ` ` + tourPackagePromotionGroupByStatement
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, nil
}

func (repository TourPackagePromotionRepository) ReadBy(column, value, operator string) (data models.TourPackagePromotion, err error) {
	statement := tourPackagePromotionSelectStatement + ` from tour_package_promotions tpp ` + tourPackagePromotionJoinStatement + ` where ` + column + `` + operator + `$1 and tpp.deleted_at is null ` + tourPackagePromotionGroupByStatement
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, nil
}
