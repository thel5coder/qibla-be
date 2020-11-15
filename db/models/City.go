package models

type City struct {
	ID           string `db:"id"`
	Name         string `db:"name"`
	ProvinceID   string `db:"province_id"`
}
