package viewmodel

type ProductTemplateVm struct {
	CategID     []interface{} `xmlrpc:"categ_id,omptempty"`
	DisplayName string        `xmlrpc:"display_name,omptempty"`
	HotelOK     bool          `xmlrpc:"hotel_ok,omptempty"`
	ID          int64         `xmlrpc:"id,omptempty"`
	IsAirlines  bool          `xmlrpc:"is_airline,omptempty"`
	Rating      interface{}   `xmlrpc:"rating,omptempty"`
	RatingCount int64         `xmlrpc:"rating_count,omptempty"`
}
