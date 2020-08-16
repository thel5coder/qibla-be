package viewmodel

type SettingProductVm struct {
	ID                   string                  `json:"id"`
	ProductID            string                  `json:"product_id"`
	ProductName          string                  `json:"product_name"`
	Price                int                     `json:"price"`
	PriceUnit            string                  `json:"price_unit"`
	MaintenancePrice     int                     `json:"maintenance_price"`
	Discount             int                     `json:"discount"`
	DiscountType         string                  `json:"discount_type"`
	DiscountPeriodStart  string                  `json:"discount_period_start"`
	DiscountPeriodEnd    string                  `json:"discount_period_end"`
	DiscountPeriod       string                  `json:"discount_period"`
	Description          string                  `json:"description"`
	SubscriptionPeriod   string                  `json:"subscription_period"`
	SubscriptionFeature  string                  `json:"subscription_feature"`
	Sessions             string                  `json:"sessions"`
	CreatedAt            string                  `json:"created_at"`
	UpdatedAt            string                  `json:"updated_at"`
	DeletedAt            string                  `json:"deleted_at"`
	SubscriptionPeriods  []SubscriptionPeriodVm  `json:"subscription_periods"`
	subscriptionFeatures []SubscriptionFeatureVm `json:"subscription_features"`
}
