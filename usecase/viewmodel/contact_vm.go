package viewmodel

type ContactVm struct {
	ID                   string `json:"id"`
	BranchName           string `json:"branch_name"`
	TravelAgentName      string `json:"travel_agent_name"`
	Address              string `json:"address"`
	Longitude            string `json:"longitude"`
	Latitude             string `json:"latitude"`
	AreaCode             string `json:"area_code"`
	PhoneNumber          string `json:"phone_number"`
	SKNumber             string `json:"sk_number"`
	SKDate               string `json:"sk_date"`
	Accreditation        string `json:"accreditation"`
	AccreditationDate    string `json:"accreditation_date"`
	DirectorName         string `json:"director_name"`
	DirectorContact      string `json:"director_contact"`
	PicName              string `json:"pic_name"`
	PicContact           string `json:"pic_contact"`
	FileLogo             FileVm `json:"file_logo"`
	VirtualAccountNumber string `json:"virtual_account_number"`
	AccountNumber        string `json:"account_number"`
	AccountName          string `json:"account_name"`
	AccountBankName      string `json:"account_bank_name"`
	AccountBankCode      string `json:"account_bank_code"`
	Email                string `json:"email"`
	IsZakatPartner       bool   `json:"is_zakat_partner"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	DeletedAt            string `json:"deleted_at"`
}
