package usecase

import (
	"database/sql"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type MenuPermissionUserUseCase struct {
	*UcContract
}

func (uc MenuPermissionUserUseCase) Browse(userID string) (res []viewmodel.MenuPermissionUserVm, err error) {
	repository := actions.NewMenuUserPermissionRepository(uc.DB)
	menuPermissionUsers, err := repository.Browse(userID)
	if err != nil {
		return res, err
	}

	for _, menuPermissionUser := range menuPermissionUsers {
		res = append(res, viewmodel.MenuPermissionUserVm{
			MenuID:           menuPermissionUser.MenuID,
			MenuName:         menuPermissionUser.MenuName,
			MenuPermissionID: menuPermissionUser.MenuPermissionID,
			Permission:       menuPermissionUser.Permission,
		})
	}

	return res, err
}

func (uc MenuPermissionUserUseCase) Add(userID string, menuPermissionUsers []string, tx *sql.Tx) (err error) {
	repository := actions.NewMenuUserPermissionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	for _, menuPermissionUser := range menuPermissionUsers {
		err = repository.Add(userID, menuPermissionUser, now, now, tx)
	}

	return err
}

func (uc MenuPermissionUserUseCase) Delete(userID string, menuPermissionUsers []string, tx *sql.Tx) (err error) {
	repository := actions.NewMenuUserPermissionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	for _, menuPermissionUser := range menuPermissionUsers {
		err = repository.Delete(userID, menuPermissionUser, now, now, tx)
	}

	return err
}

func (uc MenuPermissionUserUseCase) Store(userID string, menuPermissionUsers []string, deletedMenuPermissions []string, tx *sql.Tx) (err error) {
	if len(menuPermissionUsers) > 0 {
		err = uc.Add(userID,menuPermissionUsers,tx)
		if err != nil {
			return err
		}
	}

	if len(deletedMenuPermissions)>0 {
		err = uc.Delete(userID, deletedMenuPermissions,tx)
		if err !=nil {
			return err
		}
	}

	return nil
}

func (uc MenuPermissionUserUseCase) DeleteByUser(userID string, tx *sql.Tx) (err error) {
	repository := actions.NewMenuUserPermissionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	err = repository.DeleteByUser(userID, now, now, tx)

	return err
}
