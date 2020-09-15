package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type ITransactionDetailRepository interface {
	BrowseByTransactionID(transactionID string) (data []models.TransactionDetails, err error)

	Add(transactionID, name string, fee, price float32, quantity int, tx *sql.Tx) (err error)

	Edit(transactionID, name string, fee, price float32, quantity int, tx *sql.Tx) (err error)

	DeleteByTransactionID(transactionID string, tx *sql.Tx) (err error)
}
