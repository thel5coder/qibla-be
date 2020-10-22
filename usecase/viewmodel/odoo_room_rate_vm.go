package viewmodel

type OdooRoomRateVm struct {
	DisplayName string        `xmlrpc:"display_name,omptempty"`
	ID          int64         `xmlrpc:"id,omptempty"`
	PricePromo  float64       `xmlrpc:"price_promo,omptempty"`
	PriceUnit   float64       `xmlrpc:"price_unit,omptempty"`
	RoomID      []interface{} `xmlrpc:"room_id,omptempty"`
}
