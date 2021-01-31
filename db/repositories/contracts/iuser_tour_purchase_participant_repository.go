package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

// IUserTourPurchaseParticipantRepository ...
type IUserTourPurchaseParticipantRepository interface {
	BrowseAll(userTourPurchaseID string) (data []models.UserTourPurchaseParticipant, err error)

	ReadBy(column, value string) (data models.UserTourPurchaseParticipant, err error)

	Edit(input viewmodel.UserTourPurchaseParticipantVm, tx *sql.Tx) (err error)

	EditStatus(input viewmodel.UserTourPurchaseParticipantVm, tx *sql.Tx) (err error)

	Add(model models.UserTourPurchaseParticipant, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountBy(ID, column, value string) (res int, err error)
}
