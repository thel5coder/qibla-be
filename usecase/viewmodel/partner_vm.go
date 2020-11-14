package viewmodel

type PartnerVm struct {
	ID                          string                  `json:"id"`
	UserID                      string                  `json:"user_id"`
	UserName                    string                  `json:"user_name"`
	ContractNumber              string                  `json:"contract_number"`
	WebinarStatus               bool                    `json:"webinar_status"`
	WebsiteStatus               bool                    `json:"website_status"`
	DomainSite                  string                  `json:"domain_site"`
	DomainErp                   string                  `json:"domain_erp"`
	Database                    string                  `json:"database"`
	DatabaseUsername            string                  `json:"database_username"`
	DatabasePassword            string                  `json:"database_password"`
	InvoicePublishDate          string                  `json:"invoice_publish_date"`
	DueDateAging                int                     `json:"due_date_aging"`
	IsActive                    bool                    `json:"is_active"`
	IsPaid                      bool                    `json:"is_paid"`
	Reason                      string                  `json:"reason"`
	SubscriptionPeriod          int                     `json:"subscription_period"`
	SubscriptionPeriodExpiredAt string                  `json:"subscription_period_expired_at"`
	IsSubscriptionPeriodExpired bool                    `json:"is_subscription_period_expired"`
	CreatedAt                   string                  `json:"created_at"`
	UpdatedAt                   string                  `json:"updated_at"`
	VerifiedAt                  string                  `json:"verified_at"`
	PaidAt                      string                  `json:"paid_at"`
	Contact                     ContactVm               `json:"contact"`
	Product                     SettingProductVm        `json:"product"`
	ExtraProduct                []PartnerExtraProductVm `json:"extra_product"`
}
