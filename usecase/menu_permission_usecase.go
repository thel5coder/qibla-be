package usecase

import (
	"database/sql"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type MenuPermissionUseCase struct {
	*UcContract
}

func (uc MenuPermissionUseCase) Browse(menuID string) (res []viewmodel.MenuPermissionVm, err error) {
	repository := actions.NewMenuPermissionRepository(uc.DB)
	menuPermissions, err := repository.Browse(menuID)
	if err != nil {
		return res, err
	}

	for _, menuPermission := range menuPermissions {
		res = append(res, viewmodel.MenuPermissionVm{
			ID:         menuPermission.ID,
			Permission: menuPermission.Permission,
			CreatedAt:  menuPermission.CreatedAt,
			UpdatedAt:  menuPermission.UpdatedAt,
			DeletedAt:  menuPermission.DeletedAt.String,
		})
	}

	return res, err
}

func (uc MenuPermissionUseCase) Edit(ID, permission string, tx *sql.Tx) (err error) {
	repository := actions.NewMenuPermissionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	err = repository.Edit(ID, permission, now, tx)

	return err
}

func (uc MenuPermissionUseCase) Add(menuID,permission string,tx *sql.Tx) (err error){
	repository := actions.NewMenuPermissionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	err = repository.Add(menuID,permission,now,now,tx)

	return err
}

func (uc MenuPermissionUseCase) Store(selectedPermissions []viewmodel.MenuPermissionVm,deletedPermissions []string,tx *sql.Tx) (err error){
	for _,input := range selectedPermissions {
		err = uc.Add(input.MenuID,input.Permission,tx)
		if err != nil {
			return err
		}
	}

	if len(deletedPermissions)> 0{
		for _,menuPermission := range deletedPermissions {
			err = uc.Delete(menuPermission,tx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (uc MenuPermissionUseCase) Delete(ID string,tx *sql.Tx)(err error){
	repository := actions.NewMenuPermissionRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	err = repository.Delete(ID,now,now,tx)

	return err
}
