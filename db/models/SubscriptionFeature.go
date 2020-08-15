package models

type SubscriptionFeature struct {
	ID               string `db:"id"`
	SettingProductID string `db:"setting_product_id"`
	FeatureName      string `db:"feature_name"`
}
