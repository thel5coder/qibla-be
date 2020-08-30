package models

import "database/sql"

type CrmBoard struct {
	ID                string         `db:"id"`
	CrmStoryID        string         `db:"crm_story_id"`
	ContactID         string         `db:"contact_id"`
	Opportunity       float32        `db:"opportunity"`
	ProfitExpectation float32        `db:"profit_expectation"`
	Star              int            `db:"star"`
	CreatedAt         string         `db:"created_at"`
	UpdatedAt         string         `db:"updated_at"`
	DeletedAt         sql.NullString `db:"deleted_at"`
}
