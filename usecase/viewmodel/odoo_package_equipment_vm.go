package viewmodel

type OdooPackageEquipmentVm struct {
	ID          int64  `xmlrpc:"id,omptempty"`
	DisplayName string `xmlrpc:"display_name,omptempty"`
	ProductID   interface{}  `xmlrpc:"product_id,omptempty"`
}
