package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type UserTourPurchaseRoomRepository struct{
	DB *sql.DB
}

func NewUserTourPurchaseRoomRepository(DB *sql.DB) contracts.IUserTourPurchaseRoomRepository{
	return &UserTourPurchaseRoomRepository{DB: DB}
}

func (UserTourPurchaseRoomRepository) Add(model models.UserTourPurchaseRoom, tx *sql.Tx) (err error) {
	statement := `insert into user_tour_purchase_rooms (user_tour_purchase_id,tour_package_price_id,price,quantity,created_at,updated_at) values($1,$2,$3,$4,$5,$6)`
	_,err = tx.Exec(statement,model.UserTourPurchaseID,model.TourPackagePriceID,model.Price,model.Quantity,model.CreatedAt,model.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

