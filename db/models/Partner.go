package models

import "database/sql"

type Partner struct {
	ID                          string         `db:"id"`
	UserID                      string         `db:"user_id"`
	UserName                    string         `db:"username"`
	ContactID                   string         `db:"contact_id"`
	ContractNumber              sql.NullString `db:"contract_number"`
	WebinarStatus               bool           `db:"webinar_status"`
	WebsiteStatus               bool           `db:"website_status"`
	DomainSite                  sql.NullString `db:"domain_site"`
	DomainErp                   sql.NullString `db:"domain_erp"`
	Database                    sql.NullString `db:"database"`
	DueDateAging                sql.NullInt32  `db:"due_date_aging"`
	VerifiedStatus              string         `db:"verified_status"`
	IsActive                    bool           `db:"is_active"`
	Reason                      sql.NullString `db:"reason"`
	ProductID                   string         `db:"product_id"`
	SubscriptionPeriod          int            `db:"subscription_period"`
	SubscriptionPeriodExpiredAt sql.NullString `db:"subscription_period_expired_at"`
	IsSubscriptionExpired sql.NullBool   `json:"is_subscription_expired"`
	CreatedAt             string         `db:"created_at"`
	UpdatedAt             string         `db:"updated_at"`
	DeletedAt             sql.NullString `db:"deleted_at"`
	PaidAt                sql.NullString `db:"paid_at"`
	VerifiedAt            sql.NullString `db:"verified_at"`
	IsPaid                bool           `db:"is_paid"`
	InvoicePublishDate    sql.NullString `db:"invoice_publish_date"`
	Contact               Contact        `db:"contacts"`
	DatabaseUsername      sql.NullString `db:"database_username"`
	DatabasePassword      sql.NullString `db:"database_password"`
}
