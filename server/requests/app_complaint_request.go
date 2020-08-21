package requests

type AppComplaintRequest struct {
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	TicketNumber  string `json:"ticket_number"`
	ComplaintType string `json:"complaint_type"`
	Complaint     string `json:"complaint"`
	Solution      string `json:"solution"`
	Status        string `json:"status"`
}
