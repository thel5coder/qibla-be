package requests

type TourPurchaseCreatePassengerRequest struct {
	PackagePurchaseID string                         `json:"package_purchase_id"`
	Name              string                         `json:"name"`
	Passengers        []TourPurchasePassengerRequest `json:"passengers"`
}

type TourPurchasePassengerRequest struct {
	Name string `json:"name"`
}
