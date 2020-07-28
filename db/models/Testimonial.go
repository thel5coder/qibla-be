package models

import "database/sql"

type Testimonial struct {
	ID                   string         `db:"id"`
	WebContentCategoryID string         `db:"web_content_category_id"`
	FileID               string         `db:"file_id"`
	Path                 sql.NullString `db:"path"`
	CustomerName         string         `db:"customer_name"`
	JobPosition          string         `db:"job_position"`
	Testimony            string         `db:"testimony"`
	Rating               int            `db:"rating"`
	IsActive             bool           `db:"is_active"`
	CreatedAt            string         `db:"created_at"`
	UpdatedAt            string         `db:"updated_at"`
	DeletedAt            sql.NullString `db:"deleted_at"`
}
