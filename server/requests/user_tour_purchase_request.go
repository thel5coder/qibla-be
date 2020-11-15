package requests

// UserTourPurchaseRequest ...
type UserTourPurchaseRequest struct {
	TourPackageID         string  `json:"tour_package_id" validate:"required"`
	PaymentType           string  `json:"payment_type" validate:"required,oneof=full installment"`
	CustomerName          string  `json:"customer_name"  validate:"required"`
	CustomerIdentityType  string  `json:"customer_identity_type" validate:"required"`
	IdentityNumber        string  `json:"identity_number" validate:"required"`
	FullName              string  `json:"full_name" validate:"required"`
	Sex                   string  `json:"sex" validate:"required,oneof=male female"`
	BirthDate             string  `json:"birth_date" validate:"required"`
	BirthPlace            string  `json:"birth_place" validate:"required"`
	PhoneNumber           string  `json:"phone_number" validate:"required"`
	CityID                string  `json:"city_id" validate:"required"`
	ContactID             string  `json:"contact_id"`
	MaritalStatus         string  `json:"marital_status" validate:"required"`
	CustomerAddress       string  `json:"customer_address" validate:"required"`
	OldUserTourPurchaseID string  `json:"old_user_tour_purchase_id"`
	CancelationFee        float64 `json:"cancelation_fee"`
	Total                 float64 `json:"total"`
}
