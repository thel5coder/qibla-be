package requests

// UserTourPurchaseRequest ...
type CreatePurchaseRequest struct {
	RoomRates []RoomRateRequest `json:"room_rates"`
}

type RoomRateRequest struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type RoomRatePurchaseRequest struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
	Price    int64  `json:"price"`
}
