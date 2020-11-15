package actions

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/contracts"
)

type OdooMasterPackageRepository struct {
	DB *sql.DB
}

func NewOdooMasterPackageRepository(DB *sql.DB) contracts.IOdooMasterPackageRepository {
	return &OdooMasterPackageRepository{DB: DB}
}

const (
	odooMasterPackageSelectStatement = `select m."id",m."package_equipment_id",pe."name",m."name",m."quota",m."arrival_date",m."return_date",m."notes",m."website_description",`
)

func (OdooMasterPackageRepository) BrowseAll() (data []models.OdooMasterPackage, err error) {
	return data, err
}
