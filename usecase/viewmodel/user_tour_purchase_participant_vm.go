package viewmodel

// UserTourPurchaseParticipantVm ...
type UserTourPurchaseParticipantVm struct {
	ID                        string `json:"id"`
	UserTourPurchaseID        string `json:"user_tour_purchase_id"`
	UserID                    string `json:"user_id"`
	UserEmail                 string `json:"user_email"`
	UserName                  string `json:"user_name"`
	IsNewJamaah               bool   `json:"is_new_jamaah"`
	IdentityType              string `json:"identity_type"`
	IdentityNumber            string `json:"identity_number"`
	FullName                  string `json:"full_name"`
	Sex                       string `json:"sex"`
	BirthDate                 string `json:"birth_date"`
	BirthPlace                string `json:"birth_place"`
	PhoneNumber               string `json:"phone_number"`
	CityID                    string `json:"city_id"`
	Address                   string `json:"address"`
	KkNumber                  string `json:"kk_number"`
	PassportNumber            string `json:"passport_number"`
	PassportName              string `json:"passport_name"`
	ImmigrationOffice         string `json:"immigration_office"`
	PassportValidityPeriod    string `json:"passport_validity_period"`
	NationalIDFile            string `json:"national_id_file"`
	KkFile                    string `json:"kk_file"`
	BirthCertificate          string `json:"birth_certificate"`
	MarriageCertificate       string `json:"marriage_certificate"`
	Photo3x4                  string `json:"photo_3x4"`
	Photo4x6                  string `json:"photo_4x6"`
	MeningitisFreeCertificate string `json:"meningitis_free_certificate"`
	PassportFile              string `json:"passport_file"`
	IsDepart                  bool   `json:"is_depart"`
	Status                    string `json:"status"`
	CreatedAt                 string `json:"created_at"`
	UpdatedAt                 string `json:"updated_at"`
	DeletedAt                 string `json:"deleted_at"`
}
