package requests

type CreatePassengerRequest struct {
	PackagePurchaseID string             `json:"package_purchase_id"`
	Registrant        RegistrantRequest  `json:"registrant"`
	Passengers        []PassengerRequest `json:"passengers"`
}

type PassengerRequest struct {
	IsRegistrant   bool   `json:"is_registrant"`
	Email          string `json:"email"`
	TypeOfIdentity string `json:"type_of_identity"`
	IdentityNumber string `json:"identity_number"`
	Name           string `json:"name"`
	Sex            string `json:"sex"`
	BirthDate      string `json:"birth_date"`
	BirthPlace     string `json:"birth_place"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	CityID         string `json:"city_id"`
	MaritalStatus  string `json:"marital_status"`
}

type RegistrantRequest struct {
	TypeOfIdentity string `json:"type_of_identity"`
	IdentityNumber string `json:"identity_number"`
	Name           string `json:"name"`
	Sex            string `json:"sex"`
	BirthDate      string `json:"birth_date"`
	BirthPlace     string `json:"birth_place"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	CityID         string `json:"city_id"`
	MaritalStatus  string `json:"marital_status"`
}
