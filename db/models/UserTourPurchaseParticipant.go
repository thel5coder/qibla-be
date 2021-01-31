package models

import (
	"database/sql"
	"time"
)

// UserTourPurchaseParticipant ...
type UserTourPurchaseParticipant struct {
	ID                        string         `db:"id"`
	UserTourPurchaseID        sql.NullString `db:"user_tour_purchase_id"`
	UserID                    sql.NullString `db:"user_id"`
	User                      User           `db:"user"`
	IsNewJamaah               sql.NullBool   `db:"is_new_jamaah"`
	IdentityType              sql.NullString `db:"identity_type"`
	IdentityNumber            sql.NullString `db:"identity_number"`
	FullName                  sql.NullString `db:"full_name"`
	Sex                       sql.NullString `db:"sex"`
	BirthDate                 time.Time      `db:"birth_date"`
	BirthPlace                sql.NullString `db:"birth_place"`
	PhoneNumber               sql.NullString `db:"phone_number"`
	CityID                    sql.NullString `db:"city_id"`
	Address                   sql.NullString `db:"address"`
	KkNumber                  sql.NullString `db:"kk_number"`
	PassportNumber            sql.NullString `db:"passport_number"`
	PassportName              sql.NullString `db:"passport_name"`
	ImmigrationOffice         sql.NullString `db:"immigration_office"`
	PassportValidityPeriod    sql.NullString `db:"passport_validity_period"`
	NationalIDFile            sql.NullString `db:"national_id_file"`
	KkFile                    sql.NullString `db:"kk_file"`
	BirthCertificate          sql.NullString `db:"birth_certificate"`
	MarriageCertificate       sql.NullString `db:"marriage_certificate"`
	Photo3x4                  sql.NullString `db:"photo_3x4"`
	Photo4x6                  sql.NullString `db:"photo_4x6"`
	MeningitisFreeCertificate sql.NullString `db:"meningitis_free_certificate"`
	PassportFile              sql.NullString `db:"passport_file"`
	IsDepart                  sql.NullBool   `db:"is_depart"`
	Status                    sql.NullString `db:"status"`
	CreatedAt                 time.Time      `db:"created_at"`
	UpdatedAt                 time.Time      `db:"updated_at"`
	DeletedAt                 sql.NullTime   `db:"deleted_at"`
}
