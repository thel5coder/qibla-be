package models

type OdooMasterPackage struct {
	ID                 int64  `db:"id"`
	Quota              int    `db:"quota"`
	PackageProgram     string `db:"package_program"`
	PackageProgramID   int64  `db:"package_program_id"`
	ReturnDate         string `db:"return_date"`
	DepartureDate      string `db:"departure_date"`
	WebsiteDescription string `db:"website_description"`
	Hotels             string `db:"hotels"`
	Busses             string `db:"busses"`
	Meals              string `db:"meals"`
	AirLines           string `db:"air_lines"`
	Days               string `db:"days"`
	Prices             string `db:"prices"`
}
