package viewmodel

type TourPackagePriceVm struct {
	ID            string `json:"id"`
	RoomType      string `json:"room_type"`
	RoomCapacity  int    `json:"room_capacity"`
	Price         int64  `json:"price"`
	PromoPrice    int64  `json:"promo_price"`
	AirLineClass  string `json:"air_line_class"`
	TourPackageID string `json:"tour_package_id"`
	RoomRateID    string `json:"room_rate_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
