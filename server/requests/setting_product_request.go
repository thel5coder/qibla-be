package requests

type SettingProductRequest struct {
	ProductID            string   `json:"product_id"`
	Price                int      `json:"price"`
	PriceUnit            string   `json:"price_unit"`
	MaintenancePrice     string   `json:"maintenance_price"`
	Discount             int      `json:"discount"`
	DiscountType         string   `json:"discount_type"`
	DiscountPeriodStart  string   `json:"discount_period_start"`
	DiscountPeriodEnd    string   `json:"discount_period_end"`
	Description          string   `json:"description"`
	Sessions             string   `json:"sessions"`
	SubscriptionPeriods  []int    `json:"subscription_periods"`
	SubscriptionFeatures []string `json:"subscription_features"`
}
