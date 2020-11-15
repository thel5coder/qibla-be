package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"qibla-backend/helpers/str"
	"qibla-backend/usecase/viewmodel"
	"time"
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
	err = rows.Scan(
		&d.ID, &d.TourPackageID, &d.CustomerName, &d.CustomerIdentityType, &d.IdentityNumber,
		&d.FullName, &d.Sex, &d.BirthDate, &d.PhoneNumber, &d.CityID, &d.MaritalStatus,
		&d.CustomerAddress, &d.UserID, &d.ContactID, &d.OldUserTourPurchaseID, &d.CancelationFee,
		&d.Total, &d.Status, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.User.Email, &d.User.Name, &d.Contact.BranchName, &d.Contact.TravelAgentName,
	)

	return d, err
}

func (repository UserTourPurchaseRepository) scanRow(row *sql.Row) (d models.UserTourPurchase, err error) {
	err = row.Scan(
		&d.ID, &d.TourPackageID, &d.CustomerName, &d.CustomerIdentityType, &d.IdentityNumber,
		&d.FullName, &d.Sex, &d.BirthDate, &d.PhoneNumber, &d.CityID, &d.MaritalStatus,
		&d.CustomerAddress, &d.UserID, &d.ContactID, &d.OldUserTourPurchaseID, &d.CancelationFee,
		&d.Total, &d.Status, &d.CreatedAt, &d.UpdatedAt, &d.DeletedAt,
		&d.User.Email, &d.User.Name, &d.Contact.BranchName, &d.Contact.TravelAgentName,
	)

	return d, err
}

// Browse ...
func (repository UserTourPurchaseRepository) Browse(userID, status, order, sort string, limit, offset int) (data []models.UserTourPurchase, count int, err error) {
	var conditionString string
	if userID != "" {
		conditionString += ` AND LOWER(def."user_id") = '` + userID + `'`
	}
	if status == models.UserTourPurchaseFilterStatusUnpaid {
		conditionString += ` AND def."status" = '` + models.UserTourPurchaseStatusActive + `'
		AND COUNT(unpaid."id") > 0`
	} else if status == models.UserTourPurchaseFilterStatusPaid {
		conditionString += ` AND def."status" = '` + models.UserTourPurchaseStatusActive + `'
		AND COUNT(unpaid."id") = 0`
	} else if status == models.UserTourPurchaseFilterStatusFinish {
		conditionString += ` AND def."status" = '` + models.UserTourPurchaseStatusFinish + `'`
	} else if status == models.UserTourPurchaseFilterStatusReschedule {
		conditionString += ` AND def."status" = '` + models.UserTourPurchaseStatusActive + `'
		AND def."old_user_tour_purchase_id" IS NOT NULL`
	} else if status == models.UserTourPurchaseFilterStatusCancel {
		conditionString += ` AND def."status" = '` + models.UserTourPurchaseStatusCancel + `'`
	}

	statement := models.UserTourPurchaseSelect + ` WHERE def."deleted_at" IS NULL ` + conditionString + `
		ORDER BY ` + order + ` ` + sort + ` LIMIT $1 OFFSET $2 ` + models.UserTourPurchaseGroup
	rows, err := repository.DB.Query(statement, limit, offset)
	if err != nil {
		return data, count, err
	}

	for rows.Next() {
		d, err := repository.scanRows(rows)
		if err != nil {
			return data, count, err
		}
		data = append(data, d)
	}

	statement = `SELECT COUNT(def."id") FROM "user_tour_purchases" def
	LEFT JOIN "user_tour_purchase_transactions" utpt ON utpt."user_tour_purchase_id" = def."id"
	LEFT JOIN "transactions" unpaid ON unpaid."id" = utpt."transaction_id" AND (unpaid."status" = 'pending' OR unpaid."status" = 'gagal')
	WHERE def."deleted_at" IS NULL ` + conditionString
	err = repository.DB.QueryRow(statement).Scan(&count)
	if err != nil {
		return data, count, err
	}

	return data, count, err
}

// BrowseBy ...
func (repository UserTourPurchaseRepository) BrowseBy(column, value, operator string) (data []models.UserTourPurchase, err error) {
	statement := models.UserTourPurchaseSelect + ` WHERE ` + column + `` + operator + `$1
	AND def."deleted_at" IS NULL ORDER BY def."id" DESC`
	rows, err := repository.DB.Query(statement, value)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		d, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}

	return data, err
}

// BrowseAll ...
func (repository UserTourPurchaseRepository) BrowseAll() (data []models.UserTourPurchase, err error) {
	statement := models.UserTourPurchaseSelect + ` WHERE def."deleted_at" IS NULL`
	rows, err := repository.DB.Query(statement)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		d, err := repository.scanRows(rows)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}

	return data, err
}

// ReadBy ...
func (repository UserTourPurchaseRepository) ReadBy(column, value string) (data models.UserTourPurchase, err error) {
	statement := models.UserTourPurchaseSelect + ` WHERE ` + column + `=$1
	AND def."deleted_at" IS NULL`
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, err
}

// Add ...
func (UserTourPurchaseRepository) Add(input viewmodel.UserTourPurchaseVm, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO "user_tour_purchases" (
		"tour_package_id", "payment_type", "customer_name", "customer_identity_type", "identity_number",
		"full_name", "sex", "birth_date", "birth_place", "phone_number", "city_id",
		"marital_status" "customer_address", "user_id", "contact_id", "old_user_tour_purchase_id",
		"cancelation_fee", "total", "status", "created_at", "updated_at"
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$12,$13,$14,$15,$16,$17,$18,$19,$20) returning "id"`
	err = tx.QueryRow(statement,
		str.EmptyString(input.TourPackageID), input.PaymentType, input.CustomerName,
		input.CustomerIdentityType, input.IdentityNumber, input.FullName,
		input.Sex, input.BirthDate, input.BirthPlace, input.PhoneNumber,
		input.CityID, input.MaritalStatus, input.CustomerAddress, input.UserID,
		input.ContactID, input.OldUserTourPurchaseID, input.CancelationFee, input.Total,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339), datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

// Edit ...
func (UserTourPurchaseRepository) Edit(input viewmodel.UserTourPurchaseVm, tx *sql.Tx) (err error) {
	statement := `UPDATE "user_tour_purchases" set "user_id"=$1,"transaction_id"=$2,"contact_id"=$3,
		"master_zakat_id"=$4,"type_zakat"=$5,"current_gold_price"=$6, "gold_nishab"=$7,"wealth"=$8,
		"total"=$9, "updated_at"=$10 WHERE "id"=$11 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		str.EmptyString(input.TourPackageID), input.PaymentType, input.CustomerName,
		input.CustomerIdentityType, input.IdentityNumber, input.FullName,
		input.Sex, input.BirthDate, input.BirthPlace, input.PhoneNumber,
		input.CityID, input.MaritalStatus, input.CustomerAddress, input.UserID,
		input.ContactID, input.OldUserTourPurchaseID, input.CancelationFee, input.Total,
		datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID,
	)

	return err
}

// EditStatus ...
func (UserTourPurchaseRepository) EditStatus(input viewmodel.UserTourPurchaseVm, tx *sql.Tx) (err error) {
	statement := `UPDATE "user_tour_purchases" set "status"=$1, "updated_at"=$2
	WHERE "id"=$3 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		input.Status, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID,
	)

	return err
}

// Delete ...
func (UserTourPurchaseRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `UPDATE "user_tour_purchases" SET "updated_at"=$1,"deleted_at"=$2
	WHERE "id"=$3 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return err
}

// CountBy ...
func (repository UserTourPurchaseRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `SELECT count("id") FROM "user_tour_purchases" WHERE ` + column + `=$1 AND "deleted_at" IS NULL`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `SELECT count("id") FROM "user_tour_purchases" WHERE (` + column + `=$1 AND "deleted_at" IS NULL) AND "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
