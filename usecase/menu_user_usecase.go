package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"strings"
)

type MenuUserUseCase struct {
	*UcContract
}

func (uc MenuUserUseCase) BrowseBy(column, value, operator string) (res []viewmodel.MenuUserVm, err error) {
	repository := actions.NewMenuUserRepository(uc.DB)
	menuUsers, err := repository.BrowseBy(column, value, operator)
	if err != nil {
		return res, err
	}

	for _, menuUser := range menuUsers {
		res = append(res, uc.buildBody(menuUser))
	}

	return res, err
}

func (uc MenuUserUseCase) Add(userID, menuID string) (err error) {
	repository := actions.NewMenuUserRepository(uc.DB)
	err = repository.Add(userID, menuID, uc.TX)
	if err != nil {
		return err
	}

	return nil
}

func (uc MenuUserUseCase) Delete(userID string) (err error) {
	repository := actions.NewMenuUserRepository(uc.DB)
	err = repository.Delete(userID, uc.TX)
	if err != nil {
		return err
	}

	return nil
}

func (uc MenuUserUseCase) Store(userID string, inputs []requests.MenuUserRequest) (err error) {
	menuUserPermissionUc := MenuUserPermissionUseCase{UcContract: uc.UcContract}
	rows, err := uc.BrowseBy("mu.user_id", userID, "=")
	if err != nil {
		return err
	}

	if len(rows) > 0 {
		err = uc.Delete(userID)
		if err != nil {
			return err
		}
	}

	for _, input := range inputs {
		err = uc.Add(userID, input.MenuID)
		if err != nil {
			return err
		}

		err = menuUserPermissionUc.Store(input.MenuID, input.MenuPermissions)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uc MenuUserUseCase) buildBody(model models.MenuUser) viewmodel.MenuUserVm {
	permissionStr := model.MenuPermissions
	permissionArr := strings.Split(permissionStr, ",")
	return viewmodel.MenuUserVm{
		ID:          model.ID,
		UserID:      model.UserID,
		MenuID:      model.MenuID,
		Permissions: permissionArr,
	}
}
