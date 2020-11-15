package contracts

import "qibla-backend/db/models"

type IOdooMasterPackageRepository interface {
	BrowseAll() (data []models.OdooMasterPackage,err error)
}
