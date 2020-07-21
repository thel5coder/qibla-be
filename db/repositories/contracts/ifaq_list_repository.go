package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
)

type IFaqListRepository interface {
	Browse(faqID string) (data []models.FaqList, err error)

	Add(faqID, question, answer, createdAt, updatedAt string, tx *sql.Tx) (err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	DeleteByFaqID(faqID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountBy(column, value string) (res int, err error)
}
