package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

// IDisbursementRepository ...
type IDisbursementRepository interface {
	Browse(search, contactTravelAgentName, contactBranchName, total, startPeriod, endPeriod, contactAccountBankName, status, disburseAt, order, sort string, limit, offset int) (data []models.Disbursement, count int, err error)

	BrowseBy(column, value, operator string) (data []models.Disbursement, err error)

	BrowseAll() (data []models.Disbursement, err error)

	ReadBy(column, value string) (data models.Disbursement, err error)

	Edit(input viewmodel.DisbursementVm, tx *sql.Tx) (err error)

	Add(input viewmodel.DisbursementVm, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountBy(ID, column, value string) (res int, err error)
}
