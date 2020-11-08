package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

// IDisbursementDetailRepository ...
type IDisbursementDetailRepository interface {
	BrowseAll(disbursementID string) (data []models.DisbursementDetail, err error)

	Add(input viewmodel.DisbursementDetailVm, tx *sql.Tx) (res string, err error)

	Delete(disbursementID, updatedAt, deletedAt string, tx *sql.Tx) (err error)
}
