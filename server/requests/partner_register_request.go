package requests

type PartnerRegisterRequest struct {
	InvoiceNumber      string                `json:"invoice_number"`
	ContactID          string                `json:"contact_id"`
	UserName           string                `json:"user_name"`
	ProductID          string                `json:"product_id"`
	SubscriptionPeriod int                   `json:"subscription_period"`
	WebinarStatus      bool                  `json:"webinar_status"`
	WebsiteStatus      bool                  `json:"website_status"`
	PaymentMethodCode  int32                   `json:"payment_method_code"`
	BankName           string                `json:"bank_name"`
	VaNumber           string                `json:"va_number"`
	ExtraProducts      []ExtraProductRequest `json:"extra_products"`
}
