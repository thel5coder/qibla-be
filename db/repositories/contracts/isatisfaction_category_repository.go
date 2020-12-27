package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type ISatisfactionCategoryRepository interface {
	BrowseAllBy(filters map[string]interface{},order,sort string) (data []models.SatisfactionCategory, err error)

	ReadBy(column, value string) (data []models.SatisfactionCategory, err error)

	Edit(input viewmodel.SatisfactionCategoryVm, tx *sql.Tx) (err error)

	EditUpdatedAt(model models.SatisfactionCategory,tx *sql.Tx) (err error)

	Add(input viewmodel.SatisfactionCategoryVm, tx *sql.Tx) (err error)

	DeleteBy(column, value, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountBy(ID, column, value string) (res int, err error)

	CountByParentID(parentID, ID, slug string) (res int, err error)
}
