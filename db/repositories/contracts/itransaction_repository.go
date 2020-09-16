package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ITransactionRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.Transaction, count int, err error)

	ReadBy(column, value, operator string) (data models.Transaction, err error)

	Add(input viewmodel.TransactionVm, tx *sql.Tx) (res string,err error)

	EditDueDate(ID, dueDate, updatedAt string, dueDatePeriod int) (res string, err error)

	EditStatus(ID, paymentStatus, paidDate, updatedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)

	GetInvoiceCount(month int) (res int,err error)
}
