package models

import "database/sql"

// UserTourPurchase ...
type UserTourPurchase struct {
	ID                    string          `db:"id"`
	TourPackageID         sql.NullString  `db:"tour_package_id"`
	PaymentType           sql.NullString  `db:"payment_type"`
	CustomerName          sql.NullString  `db:"customer_name"`
	CustomerIdentityType  sql.NullString  `db:"customer_identity_type"`
	IdentityNumber        sql.NullString  `db:"identity_number"`
	FullName              sql.NullString  `db:"full_name"`
	Sex                   sql.NullString  `db:"sex"`
	BirthDate             sql.NullString  `db:"birth_date"`
	BirthPlace            sql.NullString  `db:"birth_place"`
	PhoneNumber           sql.NullString  `db:"phone_number"`
	CityID                sql.NullString  `db:"city_id"`
	MaritalStatus         sql.NullString  `db:"marital_status"`
	CustomerAddress       sql.NullString  `db:"customer_address"`
	UserID                sql.NullString  `db:"user_id"`
	User                  User            `db:"user"`
	ContactID             sql.NullString  `db:"contact_id"`
	Contact               Contact         `db:"contact"`
	OldUserTourPurchaseID sql.NullString  `db:"old_user_tour_purchase_id"`
	CancelationFee        sql.NullFloat64 `db:"cancelation_fee"`
	Total                 sql.NullFloat64 `db:"total"`
	Status                sql.NullString  `db:"status"`
	CreatedAt             string          `db:"created_at"`
	UpdatedAt             string          `db:"updated_at"`
	DeletedAt             sql.NullString  `db:"deleted_at"`
}

var (
	// UserTourPurchaseFilterStatusUnpaid ...
	UserTourPurchaseFilterStatusUnpaid = "unpaid"
	// UserTourPurchaseFilterStatusPaid ...
	UserTourPurchaseFilterStatusPaid = "paid"
	// UserTourPurchaseFilterStatusFinish ...
	UserTourPurchaseFilterStatusFinish = "finish"
	// UserTourPurchaseFilterStatusReschedule ...
	UserTourPurchaseFilterStatusReschedule = "reschedule"
	// UserTourPurchaseFilterStatusCancel ...
	UserTourPurchaseFilterStatusCancel = "cancel"
	// UserTourPurchaseFilterStatusWhitelist ...
	UserTourPurchaseFilterStatusWhitelist = []string{
		UserTourPurchaseFilterStatusUnpaid, UserTourPurchaseFilterStatusPaid, UserTourPurchaseFilterStatusFinish,
		UserTourPurchaseFilterStatusReschedule, UserTourPurchaseFilterStatusCancel,
	}

	// UserTourPurchaseStatusPending ...
	UserTourPurchaseStatusPending = "pending"
	// UserTourPurchaseStatusActive ...
	UserTourPurchaseStatusActive = "active"
	// UserTourPurchaseStatusFinish ...
	UserTourPurchaseStatusFinish = "finish"
	// UserTourPurchaseStatusCancel ...
	UserTourPurchaseStatusCancel = "cancel"
	// UserTourPurchaseStatusWhitelist ...
	UserTourPurchaseStatusWhitelist = []string{
		UserTourPurchaseStatusPending, UserTourPurchaseStatusActive,
		UserTourPurchaseStatusFinish, UserTourPurchaseStatusCancel,
	}

	// UserTourPurchasePaymentTypeFull ...
	UserTourPurchasePaymentTypeFull = "full"
	// UserTourPurchasePaymentTypeInstallment ...
	UserTourPurchasePaymentTypeInstallment = "installment"

	// UserTourPurchaseSelect ...
	UserTourPurchaseSelect = `SELECT def."id", def."tour_package_id", def."payment_type", def."customer_name",
	def."customer_identity_type", def."identity_number", def."full_name", def."sex",
	def."birth_date", def."birth_place", def."phone_number", def."city_id", def."marital_status",
	def."customer_address", def."user_id", def."contact_id", def."old_user_tour_purchase_id",
	def."cancelation_fee", def."total", def."status", def."created_at", def."updated_at", def."deleted_at",
	u."email", u."name", c."branch_name", c."travel_agent_name"
	FROM "user_tour_purchases" def
	LEFT JOIN "users" u ON u."id" = def."user_id"
	LEFT JOIN "contacts" c ON c."id" = def."contact_id"
	LEFT JOIN "user_tour_purchase_transactions" utpt ON utpt."user_tour_purchase_id" = def."id"
	LEFT JOIN "transactions" unpaid ON unpaid."id" = utpt."transaction_id" AND (unpaid."status" = 'pending' OR unpaid."status" = 'gagal')`

	// UserTourPurchaseGroup ...
	UserTourPurchaseGroup = `GROUP BY def."id", u."email", u."name", c."branch_name", c."travel_agent_name"`
)
