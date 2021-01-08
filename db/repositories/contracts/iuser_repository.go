package contracts

import (
	"database/sql"
	"qibla-backend/db/models"
	"qibla-backend/usecase/viewmodel"
)

type IUserRepository interface {
	Browse(isAdminPanel bool, search, order, sort string, limit, offset int) (data []models.User, count int, err error)

	ReadBy(column, value string) (data models.User, err error)

	Edit(input viewmodel.UserVm, password string, tx *sql.Tx) (err error)

	EditPIN(ID, pin, updatedAt string) (res string, err error)

	EditFingerPrintStatus(ID,updatedAt string, status bool) (res string,err error)

	EditPassword(ID, password, updatedAt string) (res string, err error)

	EditUserName(ID, userName, updatedAt string, tx *sql.Tx) (err error)

	EditFcmDeviceToken(ID, fcmDeviceToken, updatedAt string) (res string, err error)

	EditIsActiveStatus(email,updatedAt string,status bool) (res string,err error)

	Add(input viewmodel.UserVm, password string, tx *sql.Tx) (res string, err error)

	Delete(ID, updatedAt, deletedAt string, tx *sql.Tx) (err error)

	CountBy(ID, column, value string) (res int, err error)

	CountByPk(ID string) (res int, err error)
}
