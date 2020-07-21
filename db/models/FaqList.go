package models

import "database/sql"

type FaqList struct {
	ID        string         `db:"id"`
	FaqID     string         `db:"faq_id"`
	Question  string         `db:"question"`
	Answer    string         `db:"answer"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}
