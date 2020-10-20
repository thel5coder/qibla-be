package viewmodel

import "github.com/skilld-labs/go-odoo"

type OdooProductVm struct {
	CategID     []interface{} `xmlrpc:"categ_id,omptempty"`
	DisplayName string        `xmlrpc:"display_name,omptempty"`
	HotelOK     *odoo.Bool    `xmlrpc:"hotel_ok,omptempty"`
	ID          int64         `xmlrpc:"id,omptempty"`
	IsAirlines  *odoo.Bool    `xmlrpc:"is_airline,omptempty"`
	Rating      string        `xmlrpc:"rating,omptempty"`
	RatingCount int64         `xmlrpc:"rating_count,omptempty"`
	RatingIDs   []interface{} `xmlrpc:"rating_ids,omptempty"`
}
