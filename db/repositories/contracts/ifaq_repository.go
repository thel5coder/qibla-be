package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IFaqRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.Faq, count int, err error)

	ReadBy(column, value string) (data []models.Faq, err error)

	Edit(input viewmodel.FaqVm,tx *sql.Tx) (err error)

	Add(input viewmodel.FaqVm,webContentCategoryID string,tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string,tx *sql.Tx) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)

	CountByPk(ID string) (res int, err error)
}
