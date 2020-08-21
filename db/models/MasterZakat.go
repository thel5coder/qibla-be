package models

import "database/sql"

type MasterZakat struct {
	ID               string         `db:"id"`
	Slug             string         `db:"slug"`
	TypeZakat        string         `db:"type_zakat"`
	Name             string         `db:"name"`
	Description      string         `db:"description"`
	Amount           sql.NullInt32  `db:"amount"`
	CurrentGoldPrice sql.NullInt32  `db:"current_gold_price"`
	GoldNishab       sql.NullInt32  `db:"gold_nishab"`
	CreatedAt        string         `db:"created_at"`
	UpdatedAt        string         `db:"updated_at"`
	DeletedAt        sql.NullString `db:"deleted_at"`
}
