package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IUserTourPurchaseRoomRepository interface {
	Add(model models.UserTourPurchaseRoom, tx *sql.Tx) (err error)
}
