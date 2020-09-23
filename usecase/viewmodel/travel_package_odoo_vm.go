package viewmodel

type TravelPackageOdooVm struct {
	ID                 int64         `xmlrpc:"id,omptempty"`
	Image              string        `xmlrpc:"image,omptempty"`
	ArrivalDate        string        `xmlrpc:"arrival_date,omptempty"`
	ReturnDate         string        `xmlrpc:"return_date,omptempty"`
	CreateDate         string        `xmlrpc:"create_date,omptempty"`
	DisplayName        string        `xmlrpc:"display_name,omptempty"`
	Name               string        `xmlrpc:"name,omptempty"`
	JamaahList         []interface{} `xmlrpc:"jamaah_list,omptempty"`
	HotelIDS           []interface{} `xmlrpc:"hotel_ids,omptempty"`
	RoomRateIDS        []interface{} `xmlrpc:"room_rate_ids,omptempty"`
	PackageEquipmentID []interface{} `xmlrpc:"package_equipment_id,omptempty"`
	WebsiteDescription string        `xmlrpc:"website_description,omptempty"`
}
