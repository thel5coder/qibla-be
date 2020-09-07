package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IPartnerExtraProductRepository interface {
	BrowseByPartnerID(partnerID string) (data []models.PartnerExtraProduct, err error)

	ReadBy(column, value string) (data models.PartnerExtraProduct, err error)

	Add(partnerID, productID string, tx *sql.Tx) (err error)

	DeleteBy(column, value string, tx *sql.Tx) (err error)

	CountBy(column, value string) (res int, err error)
}
