package viewmodel

type OdooRoomVm struct {
	ID             int64  `xmlrpc:"id,omptempty"`
	DisplayName    string `xmlrpc:"display_name,omptempty"`
	NumberOfPerson int64  `xmlrpc:"no_of_person,omptempty"`
}
