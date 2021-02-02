package viewmodel

type UserTourPurchaseRespVm struct {
	UserTourPurchaseID string            `json:"user_tour_purchase_id"`
	RoomRates          []RoomRatesRespVM `json:"room_rates"`
}

type RoomRatesRespVM struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
