package models

import "database/sql"

type Faq struct {
	ID                   string         `db:"id"`
	WebContentCategoryID string         `db:"web_content_category_id"`
	FaqCategoryName      string         `db:"faq_category_name"`
	CreatedAt            string         `db:"created_at"`
	UpdatedAt            string         `db:"updated_at"`
	DeletedAt            sql.NullString `db:"deleted_at"`
	FaqListID            string         `json:"faq_list_id"`
	Question             string         `db:"question"`
	Answer               string         `db:"answer"`
}
