package requests

type SettingProductRequest struct {
	ProductID           string   `json:"product_id" validate:"required"`
	Price               int      `json:"price" validate:"required"`
	PriceUnit           string   `json:"price_unit"`
	MaintenancePrice    int32    `json:"maintenance_price"`
	Discount            int32    `json:"discount"`
	DiscountType        string   `json:"discount_type"`
	DiscountPeriodStart string   `json:"discount_period_start"`
	DiscountPeriodEnd   string   `json:"discount_period_end"`
	Description         string   `json:"description"`
	Sessions            string   `json:"sessions"`
	SettingPeriods      []int    `json:"setting_periods"`
	SettingFeatures     []string `json:"setting_features"`
}
