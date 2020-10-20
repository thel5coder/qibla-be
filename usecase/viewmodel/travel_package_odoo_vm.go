package viewmodel

import "github.com/skilld-labs/go-odoo"

type TravelPackageOdooVm struct {
	IsActive            *odoo.Bool    `xmlrpc:"is_active,omptempty"`
	ID                  int64         `xmlrpc:"id,omptempty"`
	Image               string        `xmlrpc:"image,omptempty"`
	ArrivalDate         string        `xmlrpc:"arrival_date,omptempty"`
	ReturnDate          string        `xmlrpc:"return_date,omptempty"`
	CreateDate          string        `xmlrpc:"create_date,omptempty"`
	DisplayName         string        `xmlrpc:"display_name,omptempty"`
	Name                string        `xmlrpc:"name,omptempty"`
	RoomRateIDS         []interface{} `xmlrpc:"room_rate_ids,omptempty"`
	PackageEquipmentID  []interface{} `xmlrpc:"package_equipment_id,omptempty"`
	EquipmentPackageIDs []interface{} `xmlrpc:"equipment_package_ids,omptempty"`
	WebsiteDescription  string        `xmlrpc:"website_description,omptempty"`
}
