package models

import "database/sql"

type TermConditions struct {
	ID          string         `db:"id"`
	TermName    string         `db:"term_name"`
	TermType    string         `db:"term_type"`
	Description string         `db:"description"`
	CreatedAt   string         `db:"created_at"`
	UpdatedAt   string         `db:"updated_at"`
	DeletedAt   sql.NullString `db:"deleted_at"`
}
