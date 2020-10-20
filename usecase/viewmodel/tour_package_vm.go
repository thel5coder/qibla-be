package viewmodel

type TourPackageVm struct {
	OdooID        int64                  `json:"odoo_id"`
	Name          string                 `json:"name"`
	Package       string                 `json:"package"`
	ProgramDays   string                 `json:"program_days"`
	DepartureDate string                 `json:"departure_date"`
	ReturnDate    string                 `json:"return_date"`
	Hotels        []TourPackageHotelVm   `json:"hotels"`
	Airlines      []TourPackageAirlineVm `json:"airlines"`
	Prices        []TourPackagePriceVm   `json:"prices"`
}

type TourPackageHotelVm struct {
	Name           string `json:"name"`
	FacilityRating int64  `json:"facility_rating"`
	Location       string `json:"location"`
}

type TourPackageAirlineVm struct {
	Name string `json:"name"`
}

type TourPackagePriceVm struct {
	RoomType     string  `json:"room_type"`
	RoomCapacity string  `json:"room_capacity"`
	Price        float32 `json:"price"`
	PromoPrice   float32 `json:"promo_price"`
	IsDefault    bool    `json:"is_default"`
}
