package viewmodel

// UserTourPurchaseVm ...
type UserTourPurchaseVm struct {
	ID                     string  `json:"id"`
	TourPackageID          string  `json:"tour_package_id"`
	PaymentType            string  `json:"payment_type"`
	CustomerName           string  `json:"customer_name"`
	CustomerIdentityType   string  `json:"customer_identity_type"`
	IdentityNumber         string  `json:"identity_number"`
	FullName               string  `json:"full_name"`
	Sex                    string  `json:"sex"`
	BirthDate              string  `json:"birth_date"`
	BirthPlace             string  `json:"birth_place"`
	PhoneNumber            string  `json:"phone_number"`
	CityID                 string  `json:"city_id"`
	MaritalStatus          string  `json:"marital_status"`
	CustomerAddress        string  `json:"customer_address"`
	UserID                 string  `json:"user_id"`
	UserEmail              string  `json:"user_email"`
	UserName               string  `json:"user_name"`
	ContactID              string  `json:"contact_id"`
	ContactBranchName      string  `json:"contact_branch_name"`
	ContactTravelAgentName string  `json:"contact_travel_agent_name"`
	OldUserTourPurchaseID  string  `json:"old_user_tour_purchase_id"`
	CancelationFee         float64 `json:"cancelation_fee"`
	Total                  float64 `json:"total"`
	Status                 string  `json:"status"`
	CreatedAt              string  `json:"created_at"`
	UpdatedAt              string  `json:"updated_at"`
	DeletedAt              string  `json:"deleted_at"`
}
