package contracts

import (
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ITermConditionRepository interface {
	Browse(search, order, sort string, limit, offset int) (data []models.TermConditions, count int, err error)

	ReadBy(column, value string) (data models.TermConditions, err error)

	Edit(input viewmodel.TermConditionVm) (res string, err error)

	Add(input viewmodel.TermConditionVm) (res string, err error)

	Delete(ID, updatedAt, deletedAt string) (res string, err error)

	CountBy(ID, column, value string) (res int, err error)

	CountByPk(ID string) (res int, err error)
}
