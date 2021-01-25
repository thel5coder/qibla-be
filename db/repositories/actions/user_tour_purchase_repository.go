package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

// UserTourPurchaseRepository ...
type UserTourPurchaseRepository struct {
	DB *sql.DB
}

// NewUserTourPurchaseRepository ...
func NewUserTourPurchaseRepository(DB *sql.DB) contracts.IUserTourPurchaseRepository {
	return &UserTourPurchaseRepository{DB: DB}
}

func (repository UserTourPurchaseRepository) scanRows(rows *sql.Rows) (d models.UserTourPurchase, err error) {

	return d, err
}

func (repository UserTourPurchaseRepository) scanRow(row *sql.Row) (d models.UserTourPurchase, err error) {

	return d, err
}

// Browse ...
func (repository UserTourPurchaseRepository) Browse(userID, status, order, sort string, limit, offset int) (data []models.UserTourPurchase, count int, err error) {
	//var conditionString string
	//if userID != "" {
	//	conditionString += ` AND LOWER(def."user_id") = '` + userID + `'`
	//}
	//if status == models.UserTourPurchaseFilterStatusUnpaid {
	//	conditionString += ` AND ((def."status" = '` + models.UserTourPurchaseStatusActive + `'
	//	AND COUNT(unpaid."id") > 0) OR (def."status" = '` + models.UserTourPurchaseStatusPending + `'))`
	//} else if status == models.UserTourPurchaseFilterStatusPaid {
	//	conditionString += ` AND def."status" = '` + models.UserTourPurchaseStatusActive + `'
	//	AND COUNT(unpaid."id") = 0`
	//} else if status == models.UserTourPurchaseFilterStatusFinish {
	//	conditionString += ` AND def."status" = '` + models.UserTourPurchaseStatusFinish + `'`
	//} else if status == models.UserTourPurchaseFilterStatusReschedule {
	//	conditionString += ` AND def."status" = '` + models.UserTourPurchaseStatusActive + `'
	//	AND def."old_user_tour_purchase_id" IS NOT NULL`
	//} else if status == models.UserTourPurchaseFilterStatusCancel {
	//	conditionString += ` AND def."status" = '` + models.UserTourPurchaseStatusCancel + `'`
	//}
	//
	//statement := models.UserTourPurchaseSelect + ` WHERE def."deleted_at" IS NULL ` + conditionString + `
	//	ORDER BY ` + order + ` ` + sort + ` LIMIT $1 OFFSET $2 ` + models.UserTourPurchaseGroup
	//rows, err := repository.DB.Query(statement, limit, offset)
	//if err != nil {
	//	return data, count, err
	//}
	//
	//for rows.Next() {
	//	d, err := repository.scanRows(rows)
	//	if err != nil {
	//		return data, count, err
	//	}
	//	data = append(data, d)
	//}
	//
	//statement = `SELECT COUNT(def."id") FROM "user_tour_purchases" def
	//LEFT JOIN "user_tour_purchase_transactions" utpt ON utpt."user_tour_purchase_id" = def."id"
	//LEFT JOIN "transactions" unpaid ON unpaid."id" = utpt."transaction_id" AND (unpaid."status" = 'pending' OR unpaid."status" = 'gagal')
	//WHERE def."deleted_at" IS NULL ` + conditionString
	//err = repository.DB.QueryRow(statement).Scan(&count)
	//if err != nil {
	//	return data, count, err
	//}

	return data, count, err
}

//// BrowseBy ...
//func (repository UserTourPurchaseRepository) BrowseBy(column, value, operator string) (data []models.UserTourPurchase, err error) {
//	//statement := models.UserTourPurchaseSelect + ` WHERE ` + column + `` + operator + `$1
//	//AND def."deleted_at" IS NULL ORDER BY def."id" DESC`
//	//rows, err := repository.DB.Query(statement, value)
//	//if err != nil {
//	//	return data, err
//	//}
//	//
//	//for rows.Next() {
//	//	d, err := repository.scanRows(rows)
//	//	if err != nil {
//	//		return data, err
//	//	}
//	//	data = append(data, d)
//	//}
//
//	return data, err
//}
//
//// BrowseAll ...
//func (repository UserTourPurchaseRepository) BrowseAll() (data []models.UserTourPurchase, err error) {
//	statement := models.UserTourPurchaseSelect + ` WHERE def."deleted_at" IS NULL`
//	rows, err := repository.DB.Query(statement)
//	if err != nil {
//		return data, err
//	}
//
//	for rows.Next() {
//		d, err := repository.scanRows(rows)
//		if err != nil {
//			return data, err
//		}
//		data = append(data, d)
//	}
//
//	return data, err
//}
//
//// ReadBy ...
//func (repository UserTourPurchaseRepository) ReadBy(column, value string) (data models.UserTourPurchase, err error) {
//	statement := models.UserTourPurchaseSelect + ` WHERE ` + column + `=$1
//	AND def."deleted_at" IS NULL`
//	row := repository.DB.QueryRow(statement, value)
//	data, err = repository.scanRow(row)
//	if err != nil {
//		return data, err
//	}
//
//	return data, err
//}

// Add ...
func (UserTourPurchaseRepository) Add(model models.UserTourPurchase, tx *sql.Tx) (res string, err error) {
	statement := `insert into user_tour_purchases (tour_package_id,user_id,created_at,updated_at) values($1,$2,$3,$4) returning id`
	err = tx.QueryRow(statement, model.TourPackageID, model.UserID, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// Edit ...
func (UserTourPurchaseRepository) Edit(model models.UserTourPurchase, tx *sql.Tx) (err error) {
	statement := `update user_tour_purchases set customer_identity_type=$1, identity_number=$2, full_name=$3, sex=$4, birth_date=$5, birth_place=$6, phone_number=$7, city_id=$8, marital_status=$9, 
                  customer_address=$10, updated_at=$11 where tour_package_id=$12`
	_, err = tx.Exec(statement, model.CustomerIdentityType.String, model.IdentityNumber.String, model.FullName.String, model.Sex.String, model.BirthDate.Time, model.BirthPlace.String, model.PhoneNumber.String,
		model.CityID.String, model.MaritalStatus.String, model.CustomerAddress.String, model.UpdatedAt, model.TourPackageID)
	if err != nil {
		return err
	}

	return nil
}
