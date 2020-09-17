package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type ITransactionHistoryRepository interface {
	ReadBy(column,value string) (data models.TransactionHistory,err error)

	EditByTrxID(trxID,status,response,updatedAt string) (res string,err error)

	Add(trxID,status,response,createdAt,updatedAt string,tx *sql.Tx) (err error)
}
