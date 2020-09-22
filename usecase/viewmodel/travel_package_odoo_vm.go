package viewmodel

type TravelPackageOdooVm struct {
	ID          int64         `xmlrpc:"id,omptempty"`
	DisplayName string        `xmlrpc:"display_name,omptempty"`
	Name        string        `xmlrpc:"name,omptempty"`
	JamaahList  []interface{} `xmlrpc:"jamaah_list,omptempty"`
}
