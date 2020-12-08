package models

import "database/sql"

// UserZakat ...
type UserZakat struct {
	ID               string         `db:"id"`
	UserID           sql.NullString `db:"user_id"`
	User             User           `db:"user"`
	TransactionID    sql.NullString `db:"transaction_id"`
	Transaction      Transaction    `db:"transaction"`
	ContactID        sql.NullString `db:"contact_id"`
	Contact          Contact        `db:"contact"`
	MasterZakatID    sql.NullString `db:"master_zakat_id"`
	MasterZakat      MasterZakat    `db:"master_zakat_helper"`
	TypeZakat        sql.NullString `db:"type_zakat"`
	CurrentGoldPrice sql.NullInt32  `db:"current_gold_price"`
	GoldNishab       sql.NullInt32  `db:"gold_nishab"`
	Wealth           sql.NullInt32  `db:"wealth"`
	Total            sql.NullInt32  `db:"total"`
	CreatedAt        string         `db:"created_at"`
	UpdatedAt        string         `db:"updated_at"`
	DeletedAt        sql.NullString `db:"deleted_at"`
}

var (
	// UserZakatSelect ...
	UserZakatSelect = `SELECT uz."id", uz."user_id", uz."transaction_id", uz."contact_id",
	uz."master_zakat_id", uz."type_zakat", uz."current_gold_price", uz."gold_nishab",
	uz."wealth", uz."total", uz."created_at", uz."updated_at", uz."deleted_at",
	u."email", u."name", t."invoice_number" as transaction_invoice_number, t."payment_method_code", t."payment_status",
	t."due_date", t."va_number", t."bank_name" as transaction_bank_name, c."branch_name", c."travel_agent_name"
	FROM "user_zakats" uz
	LEFT JOIN "users" u ON u."id" = uz."user_id"
	LEFT JOIN "transactions" t ON t."id" = uz."transaction_id"
	LEFT JOIN "contacts" c ON c."id" = uz."contact_id"`
)
