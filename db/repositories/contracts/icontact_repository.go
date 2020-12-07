package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IContactRepository interface {
	Browse(filters map[string]interface{}, order, sort string, limit, offset int) (data []models.Contact, count int, err error)

	BrowseAll(search string, isZakatPartner bool) (data []models.Contact, err error)

	BrowseAllZakatDisbursement() (data []models.Contact, err error)

	ReadBy(column, value string) (data models.Contact, err error)

	Edit(input viewmodel.ContactVm) (res string, err error)

	Add(input viewmodel.ContactVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)
}
