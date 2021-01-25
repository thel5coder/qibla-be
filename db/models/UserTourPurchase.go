package models

import (
	"database/sql"
	"time"
)

// UserTourPurchase ...
type UserTourPurchase struct {
	ID                   string         `db:"id"`
	TourPackageID        string         `db:"tour_package_id"`
	PaymentType          sql.NullString `db:"payment_type"`
	CustomerIdentityType sql.NullString `db:"customer_identity_type"`
	IdentityNumber       sql.NullString `db:"identity_number"`
	FullName             sql.NullString `db:"full_name"`
	Sex                  sql.NullString `db:"sex"`
	BirthDate            sql.NullTime   `db:"birth_date"`
	BirthPlace           sql.NullString `db:"birth_place"`
	PhoneNumber          sql.NullString `db:"phone_number"`
	CityID               sql.NullString `db:"city_id"`
	MaritalStatus        sql.NullString `db:"marital_status"`
	CustomerAddress      sql.NullString `db:"customer_address"`
	UserID               string         `db:"user_id"`
	User                 User           `db:"user"`
	CreatedAt            time.Time      `db:"created_at"`
	UpdatedAt            time.Time      `db:"updated_at"`
	DeletedAt            sql.NullTime   `db:"deleted_at"`
}
