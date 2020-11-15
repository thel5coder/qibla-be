package models

type OdooMasterPackage struct {
	ID                   string `json:"id"`
	EquipmentPackageID   int32  `json:"equipment_package_id"`
	EquipmentPackageName string `json:"equipment_package_name"`
	Name                 string `json:"name"`
	Quota                int32  `json:"quota"`
	DepartureDate        string `json:"departure_date"`
	ReturnDate           string `json:"return_date"`
	Notes                string `json:"notes"`
	WebsiteDescription   string `json:"website_description"`
	Hotels               string `json:"hotels"`
	Meals                string `json:"meals"`
	Transportations      string `json:"transportations"`
	Airlines             string `json:"airlines"`
	RoomRates            string `json:"room_rates"`
}
