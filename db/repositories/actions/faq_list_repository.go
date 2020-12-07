package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
	"qibla-backend/pkg/datetime"
	"time"
)

type FaqListRepository struct {
	DB *sql.DB
}

func NewFaqListRepository(DB *sql.DB) contracts.IFaqListRepository {
	return &FaqListRepository{DB: DB}
}

func (repository FaqListRepository) Browse(faqID string) (data []models.FaqList, err error) {
	statement := `select * from "faq_lists" where "faq_id"=$1`
	rows, err := repository.DB.Query(statement, faqID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		dataTemp := models.FaqList{}

		err = rows.Scan(
			&dataTemp.ID,
			&dataTemp.FaqID,
			&dataTemp.Question,
			&dataTemp.Answer,
			&dataTemp.CreatedAt,
			&dataTemp.UpdatedAt,
			&dataTemp.DeletedAt,
		)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}

	return data, err
}

func (FaqListRepository) Add(faqID, question, answer, createdAt, updatedAt string, tx *sql.Tx) (err error) {
	statement := `insert into "faq_lists" ("faq_id","question","answer","created_at","updated_at") values($1,$2,$3,$4,$5)`
	_, err = tx.Exec(statement, faqID, question, answer, datetime.StrParseToTime(createdAt, time.RFC3339), datetime.StrParseToTime(updatedAt, time.RFC3339))

	return err
}

func (repository FaqListRepository) Delete(ID, updatedAt, deletedAt string) (res string, err error) {
	statement := `update "faq_lists" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), ID).Scan(&res)

	return res, err
}

func (FaqListRepository) DeleteByFaqID(faqID, updatedAt, deletedAt string, tx *sql.Tx) (err error) {
	statement := `update "faq_lists" set "updated_at"=$1, "deleted_at"=$2 where "faq_id"=$3`
	_, err = tx.Exec(statement, datetime.StrParseToTime(updatedAt, time.RFC3339), datetime.StrParseToTime(deletedAt, time.RFC3339), faqID)

	return err
}

func (repository FaqListRepository) CountBy(column, value string) (res int,err error) {
	statement := `select count("id") from "faq_lists" where `+column+`=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement,value).Scan(&res)

	return res,err
}
