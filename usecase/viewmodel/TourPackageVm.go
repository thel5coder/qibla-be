package viewmodel

type TourPackageVm struct {
	ID            string `json:"id"`
	OdooPackageID string `json:"odoo_package_id"`
	Detail        string `json:"detail"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type TourDetailOdooVm struct {

}
