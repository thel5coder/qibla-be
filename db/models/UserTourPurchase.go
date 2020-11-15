package models

import "database/sql"

// UserTourPurchase ...
type UserTourPurchase struct {
	ID                   string          `db:"id"`
	TourPackageID        sql.NullString  `db:"tour_package_id"`
	CustomerName         sql.NullString  `db:"customer_name"`
	CustomerIdentityType sql.NullString  `db:"customer_identity_type"`
	IdentityNumber       sql.NullString  `db:"identity_number"`
	FullName             sql.NullString  `db:"full_name"`
	Sex                  sql.NullString  `db:"sex"`
	BirthDate            sql.NullString  `db:"birth_date"`
	BirthPlace           sql.NullString  `db:"birth_place"`
	PhoneNumber          sql.NullString  `db:"phone_number"`
	CityID               sql.NullString  `db:"city_id"`
	MaritalStatus        sql.NullString  `db:"marital_status"`
	CustomerAddress      sql.NullString  `db:"customer_address"`
	UserID               sql.NullString  `db:"user_id"`
	User                 User            `db:"user"`
	ContactID            sql.NullString  `db:"contact_id"`
	Contact              Contact         `db:"contact"`
	AirlineOdooID        sql.NullInt32   `db:"airline_odoo_id"`
	AirlineClass         sql.NullString  `db:"airlines_class"`
	AirlinePrice         sql.NullFloat64 `db:"airline_price"`
	CreatedAt            string          `db:"created_at"`
	UpdatedAt            string          `db:"updated_at"`
	DeletedAt            sql.NullString  `db:"deleted_at"`
}

var (
	// UserTourPurchaseSelect ...
	UserTourPurchaseSelect = `SELECT def."id", def."tour_package_id", def."customer_name",
	def."customer_identity_type", def."identity_number", def."full_name", def."sex",
	def."birth_date", def."birth_place", def."phone_number", def."city_id", def."marital_status"
	def."customer_address", def."user_id", def."contact_id", def."airline_odoo_id", def."airlines_class"
	def."airline_price", def."created_at", def."updated_at", def."deleted_at",
	u."email", u."name", c."branch_name", c."travel_agent_name"
	FROM "user_zakats" def
	LEFT JOIN "users" u ON u."id" = def."user_id"
	LEFT JOIN "contacts" c ON c."id" = def."contact_id"`
)
