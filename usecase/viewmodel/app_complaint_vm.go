package viewmodel

type AppComplaintVm struct {
	ID            string `json:"id"`
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	TicketNumber  string `json:"ticket_number"`
	ComplaintType string `json:"complaint_type"`
	Complaint     string `json:"complaint"`
	Solution      string `json:"solution"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}
