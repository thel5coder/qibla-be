package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IPromotionPositionRepository interface {
	BrowseByPromotionPlatformID(promotionPlatformID string) (data []models.PromotionPosition,err error)

	Add(promotionPlatformID, position string,tx *sql.Tx) (err error)

	Delete(promotionPlatformID string,tx *sql.Tx) (err error)
}
