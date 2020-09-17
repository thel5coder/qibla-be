package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type TransactionDetailRepository struct{
	DB *sql.DB
}

func NewTransactionDetailRepository(DB *sql.DB) contracts.ITransactionDetailRepository{
	return &TransactionDetailRepository{DB: DB}
}

func (repository TransactionDetailRepository) BrowseByTransactionID(transactionID string) (data []models.TransactionDetails, err error) {
	statement :=`select * from "transaction_details" where "transaction_id"=$1`
	rows,err := repository.DB.Query(statement,transactionID)
	if err != nil {
		return data,err
	}

	for rows.Next(){
		dataTemp := models.TransactionDetails{}
		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.Name,
			&dataTemp.Fee,
			&dataTemp.Price,
			&dataTemp.TransactionID,
			&dataTemp.Quantity,
			)
		if err != nil {
			return data,err
		}
		data = append(data,dataTemp)
	}

	return data,err
}

func (TransactionDetailRepository) Add(transactionID, name string, fee, price float32, quantity int, tx *sql.Tx) (err error) {
	statement := `insert into "transaction_details" ("transaction_id","name","fee","price","quantity") values($1,$2,$3,$4,$5)`
	_,err = tx.Exec(
		statement,
		transactionID,
		name,
		fee,
		price,
		quantity,
		)

	return err
}

func (TransactionDetailRepository) Edit(transactionID, name string, fee, price float32, quantity int, tx *sql.Tx) (err error) {
	panic("implement me")
}

func (TransactionDetailRepository) DeleteByTransactionID(transactionID string, tx *sql.Tx) (err error) {
	statement := `delete from "transaction_details" where "transaction_id"=$1`
	_,err = tx.Exec(statement,transactionID)

	return err
}
