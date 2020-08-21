package models

import "database/sql"

type AppComplaint struct {
	ID            string         `db:"id"`
	FullName      string         `db:"full_name"`
	Email         string         `db:"email"`
	TicketNumber  string         `db:"ticket_number"`
	ComplaintType string         `db:"complaint_type"`
	Complaint     string         `db:"complaint"`
	Solution      sql.NullString `db:"solution"`
	Status        string         `db:"status"`
	CreatedAt     string         `db:"created_at"`
	UpdatedAt     string         `db:"updated_at"`
	DeletedAt     sql.NullString `db:"deleted_at"`
}
