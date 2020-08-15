package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IPromotionPlatformRepository interface {
	BrowseByPromotionID(promotionID string) (data []models.PromotionPlatform, err error)

	Add(promotionID, platform string, tx *sql.Tx) (res string, err error)

	Delete(promotionID string, tx *sql.Tx) (err error)
}
