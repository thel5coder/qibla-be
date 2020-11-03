package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
)

type MenuUserPermissionUseCase struct {
	*UcContract
}

func (uc MenuUserPermissionUseCase) Browse(menuID string) (res []viewmodel.MenuUserPermissionVm, err error) {
	repository := actions.NewMenuUserPermissionRepository(uc.DB)
	menuPermissionUsers, err := repository.Browse(menuID)
	if err != nil {
		return res, err
	}

	for _, menuPermissionUser := range menuPermissionUsers {
		res = append(res, uc.buildBody(menuPermissionUser))
	}

	return res, err
}

func (uc MenuUserPermissionUseCase) Add(menuID,permissionID string) (err error) {
	repository := actions.NewMenuUserPermissionRepository(uc.DB)
	err = repository.Add(menuID,permissionID,uc.TX)
	if err != nil {
		return err
	}

	return nil
}

func (uc MenuUserPermissionUseCase) Delete(menuID string) (err error) {
	repository := actions.NewMenuUserPermissionRepository(uc.DB)
	err = repository.Delete(menuID,uc.TX)
	if err != nil {
		return err
	}

	return nil
}

func (uc MenuUserPermissionUseCase) Store(menuID string, menuPermissions []string) (err error) {
	rows,err := uc.Browse(menuID)
	if err != nil {
		return err
	}

	if len(rows) > 0 {
		err = uc.Delete(menuID)
		if err != nil {
			return err
		}
	}

	for _, menuPermission := range menuPermissions{
		err = uc.Add(menuID,menuPermission)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc MenuUserPermissionUseCase) buildBody(model models.MenuUserPermission) viewmodel.MenuUserPermissionVm{
	return viewmodel.MenuUserPermissionVm{
		MenuID:           model.MenuID,
		MenuPermissionID: model.MenuPermissionID,
	}
}
