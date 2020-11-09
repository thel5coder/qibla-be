package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IUserZakatRepository interface {
	Browse(search, createdAt, bankName, typeZakat, invoiceNumber, total, travelAgentName, order, sort string, limit, offset int) (data []models.UserZakat, count int, err error)

	BrowseBy(column, value, operator string) (data []models.UserZakat, err error)

	BrowseAll() (data []models.UserZakat, err error)

	ReadBy(column, value string) (data models.UserZakat, err error)

	Edit(input viewmodel.UserZakatVm, tx *sql.Tx) (err error)

	EditTransaction(input viewmodel.UserZakatVm, tx *sql.Tx) (err error)

	Add(input viewmodel.UserZakatVm, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountBy(ID, column, value string) (res int, err error)
}
