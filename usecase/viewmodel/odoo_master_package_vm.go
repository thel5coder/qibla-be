package viewmodel

type OdooMasterPackageVm struct {
	OdooID          int32                               `json:"odoo_id"`
	Name            string                              `json:"name"`
	PackageType     string                              `json:"package_type"`
	PackageTypeID   int32                               `json:"package_type_id"`
	ProgramDays     int32                              `json:"program_days"`
	DepartureDate   string                              `json:"departure_date"`
	ReturnDate      string                              `json:"return_date"`
	Quota           int32                               `json:"quota"`
	Notes           string                              `json:"notes"`
	WebDescription  string                              `json:"web_description"`
	Hotels          []OdooMasterPackagePackageHotelVm   `json:"hotels"`
	Airlines        []OdooMasterPackageAirlineVm        `json:"airlines"`
	Meals           []OdooMasterPackageMealsVm          `json:"meals"`
	Transportations []OdooMasterPackageTransportationVm `json:"transportations"`
	RoomRates       []OdooMasterPackageRoomRate         `json:"room_rates"`
}

type OdooMasterPackagePackageHotelVm struct {
	OdooHotelID       int32  `json:"odoo_hotel_id"`
	ProductTemplateID int32  `json:"product_template_id"`
	Name              string `json:"name"`
	FacilityRating    int32  `json:"facility_rating"`
	Location          string `json:"location"`
}

type OdooMasterPackageAirlineVm struct {
	OdooTourAirlineID int32  `json:"odoo_tour_airline_id"`
	Name              string `json:"name"`
}

type OdooMasterPackageMealsVm struct {
	OdooMealID        int32  `json:"odoo_meal_id"`
	ProductTemplateID int32  `json:"product_template_id"`
	Name              string `json:"name"`
}

type OdooMasterPackageTransportationVm struct {
	OdooTransportationID  int32  `json:"odoo_transportation_id"`
	OdooProductTemplateID int32  `json:"odoo_product_template_id"`
	Name                  string `json:"name"`
}

type OdooMasterPackageRoomRate struct {
	RoomRateID   int32   `json:"room_rate_id"`
	RoomType     string  `json:"room_type"`
	Price        float32 `json:"price"`
	PricePromo   float32 `json:"promo_price"`
	RoomCapacity int32   `json:"room_capacity"`
	AirlineClass string  `json:"airline_class"`
}
