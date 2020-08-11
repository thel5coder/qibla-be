package models

import "database/sql"

type MasterProduct struct {
	ID               string         `json:"id"`
	Slug             string         `json:"slug"`
	Name             string         `json:"name"`
	SubscriptionType string         `json:"subscription_type"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
	DeletedAt        sql.NullString `json:"deleted_at"`
}
