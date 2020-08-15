package models

type SubscriptionPeriod struct {
	ID               string `db:"id"`
	SettingProductID string `db:"setting_product_id"`
	Period           int    `db:"period"`
}
