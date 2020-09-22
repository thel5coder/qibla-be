package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IPartnerRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.Partner, count int, err error)

	BrowseProfilePartner(search, order, sort string, limit, offset int) (data []models.Partner, count int, err error)

	ReadBy(column, value string) (data models.Partner, err error)

	Edit(input viewmodel.PartnerVm, tx *sql.Tx) (err error)

	EditVerified(input viewmodel.PartnerVm) (res string, err error)

	EditBoolStatus(ID, column, reason, updatedAt string, value bool) (res string, err error)

	EditPaymentStatus(ID, paidAt, updatedAt string, tx *sql.Tx) (err error)

	Add(input viewmodel.PartnerVm, tx *sql.Tx) (res string, err error)

	DeleteBy(column, value, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountBy(ID, column, value string) (res int, err error)
}
