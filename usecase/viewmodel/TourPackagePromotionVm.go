package viewmodel

type TourPackagePromotionVm struct {
	ID                 string                            `json:"id"`
	TravelAgent        TourPackagePromotionTravelAgentVm `json:"travel_agent"`
	Name               string                            `json:"name"`
	ProgramDay         int                               `json:"program_day"`
	DepartureDate      string                            `json:"departure_date"`
	ReturnDate         string                            `json:"return_date"`
	DepartureAirport   string                            `json:"departure_airport"`
	DestinationAirport string                            `json:"destination_airport"`
	Description        string                            `json:"description"`
	Image              string                            `json:"image"`
	PackageType        string                            `json:"package_type"`
	Hotels             []TourPackagePromotionHotelVm       `json:"hotels"`
	RoomRates          []TourPackagePromotionRoomRateVm    `json:"room_rates"`
}

type TourPackagePromotionRoomRateVm struct {
	ID           string `json:"id"`
	Room         string `json:"room"`
	AirlineClass string `json:"airline_class"`
	Price        int64  `json:"price"`
	PromoPrice   int64  `json:"promo_price"`
	RoomCapacity int    `json:"room_capacity"`
}

type TourPackagePromotionHotelVm struct {
	City string `json:"city"`
	Name string `json:"name"`
}

type TourPackagePromotionTravelAgentVm struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Branch string `json:"branch"`
	Phone  string `json:"phone"`
}
