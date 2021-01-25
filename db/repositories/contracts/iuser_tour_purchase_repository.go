package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

// IUserTourPurchaseRepository ...
type IUserTourPurchaseRepository interface {
	//Browse(userID, status, order, sort string, limit, offset int) (data []models.UserTourPurchase, count int, err error)
	//
	//BrowseBy(column, value, operator string) (data []models.UserTourPurchase, err error)
	//
	//BrowseAll() (data []models.UserTourPurchase, err error)
	//
	//ReadBy(column, value string) (data models.UserTourPurchase, err error)
	//
	Edit(model models.UserTourPurchase, tx *sql.Tx) (err error)
	//
	//EditStatus(input viewmodel.UserTourPurchaseVm, tx *sql.Tx) (err error)

	Add(model models.UserTourPurchase, tx *sql.Tx) (res string, err error)

	//Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error)
	//
	//CountBy(ID, column, value string) (res int, err error)
}
