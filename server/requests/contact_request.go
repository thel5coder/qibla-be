package requests

type ContactRequest struct {
	BranchName           string `json:"branch_name"`
	TravelAgentName      string `json:"travel_agent_name"`
	Address              string `json:"address" validate:"required"`
	Longitude            string `json:"longitude"`
	Latitude             string `json:"latitude"`
	AreaCode             string `json:"area_code" validate:"required"`
	PhoneNumber          string `json:"phone_number" validate:"required"`
	SKNumber             string `json:"sk_number"`
	SKDate               string `json:"sk_date"`
	Accreditation        string `json:"accreditation"`
	AccreditationDate    string `json:"accreditation_date"`
	DirectorName         string `json:"director_name" validate:"required"`
	DirectorContact      string `json:"director_contact"`
	PicName              string `json:"pic_name"`
	PicContact           string `json:"pic_contact"`
	Logo                 string `json:"logo" validate:"required"`
	VirtualAccountNumber string `json:"virtual_account_number"`
	AccountNumber        string `json:"account_number"`
	AccountName          string `json:"account_name"`
	AccountBankName      string `json:"account_bank_name"`
	AccountBankCode      string `json:"account_bank_code"`
	Email                string `json:"email" validate:"required"`
	IsZakatPartner       bool   `json:"is_zakat_partner"`
}
