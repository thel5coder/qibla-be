package models

import "database/sql"

type PartnerExtraProduct struct {
	ID        string  `db:"id"`
	PartnerID string  `db:"partner_id"`
	Product   Product `db:"product"`
}

type Product struct {
	ID               string         `db:"id"`
	Name             string         `db:"name"`
	SubscriptionType string         `db:"subscription_type"`
	Price            float32        `db:"price"`
	PriceUnit        string         `db:"price_unit"`
	Session          sql.NullString `db:"session"`
}
