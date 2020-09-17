package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/helpers/datetime"
	"time"
)

type TransactionHistoryRepository struct{
	DB *sql.DB
}

func NewTransactionHistoryRepository(DB *sql.DB) contracts.ITransactionHistoryRepository{
	return &TransactionHistoryRepository{DB: DB}
}

func (repository TransactionHistoryRepository) ReadBy(column, value string) (data models.TransactionHistory, err error) {
	statement := `select * from "transaction_histories" where `+column+`=$1`
	err = repository.DB.QueryRow(statement,value).Scan(
		&data.ID,
		&data.TrxID,
		&data.Status,
		&data.Response,
		&data.CreatedAt,
		&data.UpdatedAt,
		)

	return data,err
}

func (repository TransactionHistoryRepository) EditByTrxID(trxID, status, response, updatedAt string) (res string,err error) {
	statement := `update "transaction_histories" set "status"=$1, "response"=$2, "updated_at"=$3 where "trx_id"=$4 returning "id"`
	err = repository.DB.QueryRow(statement,status,response,datetime.StrParseToTime(updatedAt,time.RFC3339),trxID).Scan(&res)

	return res,err
}

func (TransactionHistoryRepository) Add(trxID, status, response, createdAt, updatedAt string, tx *sql.Tx) (err error) {
	statement := `insert into "transaction_histories" ("trx_id","status","response","created_at","updated_at") values($1,$2,$3,$4,$5) returning "id"`
	_,err = tx.Exec(statement,trxID,status,response,datetime.StrParseToTime(createdAt,time.RFC3339),datetime.StrParseToTime(updatedAt,time.RFC3339))

	return err
}

