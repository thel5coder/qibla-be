package requests

type PartnerRegisterRequest struct {
	ContactID          string                `json:"contact_id"`
	UserName           string                `json:"user_name"`
	ProductID          string                `json:"product_id"`
	SubscriptionPeriod int                   `json:"subscription_period"`
	WebinarStatus      bool                  `json:"webinar_status"`
	WebsiteStatus      bool                  `json:"website_status"`
	ExtraProducts      []ExtraProductRequest `json:"extra_products"`
}
