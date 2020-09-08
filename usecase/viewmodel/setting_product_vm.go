package viewmodel

type SettingProductVm struct {
	ID                    string                    `json:"id"`
	ProductID             string                    `json:"product_id"`
	ProductName           string                    `json:"product_name"`
	ProductType           string                    `json:"product_type"`
	Price                 int                       `json:"price"`
	PriceUnit             string                    `json:"price_unit"`
	MaintenancePrice      int32                     `json:"maintenance_price"`
	Discount              int32                     `json:"discount"`
	DiscountType          string                    `json:"discount_type"`
	DiscountPeriodStart   string                    `json:"discount_period_start"`
	DiscountPeriodEnd     string                    `json:"discount_period_end"`
	DiscountPeriod        string                    `json:"discount_period"`
	Description           string                    `json:"description"`
	Sessions              string                    `json:"sessions"`
	CreatedAt             string                    `json:"created_at"`
	UpdatedAt             string                    `json:"updated_at"`
	DeletedAt             string                    `json:"deleted_at"`
	SettingProductPeriods []SettingProductPeriodVm  `json:"setting_product_periods"`
	SettingProductFeature []SettingProductFeatureVm `json:"setting_product_feature"`
}
