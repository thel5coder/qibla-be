package models

import "database/sql"

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
	BirthDate                 sql.NullString `db:"birth_date"`
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
	CreatedAt                 string         `db:"created_at"`
	UpdatedAt                 string         `db:"updated_at"`
	DeletedAt                 sql.NullString `db:"deleted_at"`
}

var (
	// UserTourPurchaseParticipantStatusPending ...
	UserTourPurchaseParticipantStatusPending = "pending"
	// UserTourPurchaseParticipantStatusApproved ...
	UserTourPurchaseParticipantStatusApproved = "approved"
	// UserTourPurchaseParticipantStatusRejected ...
	UserTourPurchaseParticipantStatusRejected = "rejected"
	// UserTourPurchaseParticipantStatusWhitelist ...
	UserTourPurchaseParticipantStatusWhitelist = []string{
		UserTourPurchaseParticipantStatusPending, UserTourPurchaseParticipantStatusApproved,
		UserTourPurchaseParticipantStatusRejected,
	}

	// UserTourPurchaseParticipantSelect ...
	UserTourPurchaseParticipantSelect = `SELECT def."id", def."user_tour_purchase_id", def."user_id",
	def."is_new_jamaah", def."identity_type", def."identity_number", def."full_name", def."sex",
	def."birth_date", def."birth_place", def."phone_number", def."city_id", def."address",
	def."kk_number", def."passport_number", def."passport_name", def."immigration_office",
	def."passport_validity_period", def."national_id_file", def."kk_file", def."birth_certificate",
	def."marriage_certificate", def."photo_3x4", def."photo_4x6", def."meningitis_free_certificate",
	def."passport_file", def."is_depart", def."status", def."created_at", def."updated_at", def."deleted_at",
	u."email", u."name"
	FROM "user_tour_purchase_participants" def
	LEFT JOIN "users" u ON u."id" = def."user_id"`

	// UserTourPurchaseParticipantGroup ...
	UserTourPurchaseParticipantGroup = `GROUP BY def."id", u."email", u."name"`
)
