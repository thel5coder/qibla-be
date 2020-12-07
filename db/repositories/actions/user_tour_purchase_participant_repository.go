package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"qibla-backend/pkg/str"
	"qibla-backend/usecase/viewmodel"
	"time"
)

// UserTourPurchaseParticipantRepository ...
type UserTourPurchaseParticipantRepository struct {
	DB *sql.DB
}

// NewUserTourPurchaseParticipantRepository ...
func NewUserTourPurchaseParticipantRepository(DB *sql.DB) contracts.IUserTourPurchaseParticipantRepository {
	return &UserTourPurchaseParticipantRepository{DB: DB}
}

func (repository UserTourPurchaseParticipantRepository) scanRows(rows *sql.Rows) (d models.UserTourPurchaseParticipant, err error) {
	err = rows.Scan(
		&d.ID, &d.UserTourPurchaseID, &d.UserID, &d.IsNewJamaah, &d.IdentityType, &d.IdentityNumber,
		&d.FullName, &d.Sex, &d.BirthDate, &d.PhoneNumber, &d.CityID, &d.Address, &d.KkNumber,
		&d.PassportNumber, &d.PassportName, &d.ImmigrationOffice, &d.PassportValidityPeriod,
		&d.NationalIDFile, &d.KkFile, &d.BirthCertificate, &d.MarriageCertificate, &d.Photo3x4, &d.Photo4x6,
		&d.MeningitisFreeCertificate, &d.PassportFile, &d.IsDepart, &d.Status, &d.CreatedAt, &d.UpdatedAt,
		&d.DeletedAt, &d.User.Email, &d.User.Name,
	)

	return d, err
}

func (repository UserTourPurchaseParticipantRepository) scanRow(row *sql.Row) (d models.UserTourPurchaseParticipant, err error) {
	err = row.Scan(
		&d.ID, &d.UserTourPurchaseID, &d.UserID, &d.IsNewJamaah, &d.IdentityType, &d.IdentityNumber,
		&d.FullName, &d.Sex, &d.BirthDate, &d.PhoneNumber, &d.CityID, &d.Address, &d.KkNumber,
		&d.PassportNumber, &d.PassportName, &d.ImmigrationOffice, &d.PassportValidityPeriod,
		&d.NationalIDFile, &d.KkFile, &d.BirthCertificate, &d.MarriageCertificate, &d.Photo3x4, &d.Photo4x6,
		&d.MeningitisFreeCertificate, &d.PassportFile, &d.IsDepart, &d.Status, &d.CreatedAt, &d.UpdatedAt,
		&d.DeletedAt, &d.User.Email, &d.User.Name,
	)

	return d, err
}

// BrowseAll ...
func (repository UserTourPurchaseParticipantRepository) BrowseAll(userTourPurchaseID string) (data []models.UserTourPurchaseParticipant, err error) {
	statement := models.UserTourPurchaseParticipantSelect + ` WHERE def."deleted_at" IS NULL
	AND def."user_tour_purchase_id" = $1`
	rows, err := repository.DB.Query(statement, userTourPurchaseID)
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
func (repository UserTourPurchaseParticipantRepository) ReadBy(column, value string) (data models.UserTourPurchaseParticipant, err error) {
	statement := models.UserTourPurchaseParticipantSelect + ` WHERE ` + column + `=$1
	AND def."deleted_at" IS NULL`
	row := repository.DB.QueryRow(statement, value)
	data, err = repository.scanRow(row)
	if err != nil {
		return data, err
	}

	return data, err
}

// Add ...
func (UserTourPurchaseParticipantRepository) Add(input viewmodel.UserTourPurchaseParticipantVm, tx *sql.Tx) (res string, err error) {
	statement := `INSERT INTO "user_tour_purchase_participants" (
		"user_tour_purchase_id", "user_id", "is_new_jamaah", "identity_type", "identity_number", "full_name",
		"sex", "birth_date", "birth_place", "phone_number", "city_id", "address", "kk_number",
		"passport_number", "passport_name", "immigration_office", "passport_validity_period",
		"national_id_file", "kk_file", "birth_certificate", "marriage_certificate", "photo_3x4", "photo_4x6",
		"meningitis_free_certificate", "passport_file", "is_depart", "status", "created_at", "updated_at"
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29) returning "id"`
	err = tx.QueryRow(statement,
		input.UserTourPurchaseID, input.UserID, input.IsNewJamaah, input.IdentityType, input.IdentityNumber,
		input.FullName, input.Sex, input.BirthDate, input.PhoneNumber, str.EmptyString(input.CityID),
		input.Address, input.KkNumber, input.PassportNumber, input.PassportName, input.ImmigrationOffice,
		input.PassportValidityPeriod, str.EmptyString(input.NationalIDFile), str.EmptyString(input.KkFile),
		str.EmptyString(input.BirthCertificate), str.EmptyString(input.MarriageCertificate),
		str.EmptyString(input.Photo3x4), str.EmptyString(input.Photo4x6),
		str.EmptyString(input.MeningitisFreeCertificate), str.EmptyString(input.PassportFile),
		input.IsDepart, input.Status, datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339),
	).Scan(&res)

	return res, err
}

// Edit ...
func (UserTourPurchaseParticipantRepository) Edit(input viewmodel.UserTourPurchaseParticipantVm, tx *sql.Tx) (err error) {
	statement := `UPDATE "user_tour_purchase_participants" set "user_tour_purchase_id" = $1, "user_id" = $2,
	"is_new_jamaah" = $3, "identity_type" = $4, "identity_number" = $5, "full_name" = $6,
	"sex" = $7, "birth_date" = $8, "birth_place" = $9, "phone_number" = $10, "city_id" = $11, "address" = $12,
	"kk_number" = $13, "passport_number" = $14, "passport_name" = $15, "immigration_office" = $16,
	"passport_validity_period" = $17, "national_id_file" = $18, "kk_file" = $19, "birth_certificate" = $20,
	"marriage_certificate" = $21, "photo_3x4" = $22, "photo_4x6" = $23, "meningitis_free_certificate" = $23,
	"passport_file" = $24, "is_depart" = $25, "status" = $26, "updated_at" = $27 WHERE "id" = $28
	AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		input.UserTourPurchaseID, input.UserID, input.IsNewJamaah, input.IdentityType, input.IdentityNumber,
		input.FullName, input.Sex, input.BirthDate, input.PhoneNumber, str.EmptyString(input.CityID),
		input.Address, input.KkNumber, input.PassportNumber, input.PassportName, input.ImmigrationOffice,
		input.PassportValidityPeriod, str.EmptyString(input.NationalIDFile), str.EmptyString(input.KkFile),
		str.EmptyString(input.BirthCertificate), str.EmptyString(input.MarriageCertificate),
		str.EmptyString(input.Photo3x4), str.EmptyString(input.Photo4x6),
		str.EmptyString(input.MeningitisFreeCertificate), str.EmptyString(input.PassportFile),
		input.IsDepart, input.Status, datetime.StrParseToTime(input.CreatedAt, time.RFC3339),
		datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID,
	)

	return err
}

// EditStatus ...
func (UserTourPurchaseParticipantRepository) EditStatus(input viewmodel.UserTourPurchaseParticipantVm, tx *sql.Tx) (err error) {
	statement := `UPDATE "user_tour_purchase_participants" set "status"=$1, "updated_at"=$2
	WHERE "id"=$3 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement,
		input.Status, datetime.StrParseToTime(input.UpdatedAt, time.RFC3339), input.ID,
	)

	return err
}

// Delete ...
func (UserTourPurchaseParticipantRepository) Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `UPDATE "user_tour_purchase_participants" SET "updated_at"=$1,"deleted_at"=$2
	WHERE "id"=$3 AND "deleted_at" IS NULL`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID)

	return err
}

// CountBy ...
func (repository UserTourPurchaseParticipantRepository) CountBy(ID, column, value string) (res int, err error) {
	if ID == "" {
		statement := `SELECT count("id") FROM "user_tour_purchase_participants" WHERE ` + column + `=$1 AND "deleted_at" IS NULL`
		err = repository.DB.QueryRow(statement, value).Scan(&res)
	} else {
		statement := `SELECT count("id") FROM "user_tour_purchase_participants" WHERE (` + column + `=$1 AND "deleted_at" IS NULL) AND "id"<>$2`
		err = repository.DB.QueryRow(statement, value, ID).Scan(&res)
	}

	return res, err
}
