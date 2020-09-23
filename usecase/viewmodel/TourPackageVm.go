package viewmodel

type TourPackageVm struct {
	ID            string `json:"id"`
	OdooPackageID string `json:"odoo_package_id"`
	Detail        string `json:"detail"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type TourDetailOdooVm struct {
	ID          int64         `xmlrpc:"id,omptempty"`
	ArrivalDate string        `xmlrpc:"arrival_date,omptempty"`
	CreateDate  string        `xmlrpc:"create_date,omptempty"`
	DisplayName string        `xmlrpc:"display_name,omptempty"`
	Name        string        `xmlrpc:"name,omptempty"`
	JamaahList  []interface{} `xmlrpc:"jamaah_list,omptempty"`

}
