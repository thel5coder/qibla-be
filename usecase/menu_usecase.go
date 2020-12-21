package usecase

import (
	"errors"
	"fmt"
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"strings"
	"time"
)

type MenuUseCase struct {
	*UcContract
}

//browse
func (uc MenuUseCase) Browse(parentID, search, order, sort string, page, limit int) (res []viewmodel.MenuVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewMenuRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	menus, count, err := repository.Browse(parentID, search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, menu := range menus {
		res = append(res, uc.buildBody(menu, true))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc MenuUseCase) BrowseAllBy(column, value, operator string, isActive bool) (res []viewmodel.MenuVm, err error) {
	repository := actions.NewMenuRepository(uc.DB)
	menus, err := repository.BrowseAllBy(column, value, operator, isActive)
	if err != nil {
		return res, err
	}

	for _, menu := range menus {
		res = append(res, uc.buildBody(menu, isActive))
	}

	return res, nil
}

func (uc MenuUseCase) browseChild(parentID string, isActive bool) (res []viewmodel.MenuVm) {
	menus, err := uc.BrowseAllBy("m.parent_id", parentID, "=", isActive)
	if err != nil {
		return res
	}

	for _, menu := range menus {
		menu.ChildMenus = uc.browseChild(menu.ID, isActive)
	}
	res = menus

	return res
}

func (uc MenuUseCase) ReadBy(column, value, operator string) (res viewmodel.MenuVm, err error) {
	repository := actions.NewMenuRepository(uc.DB)

	menu, err := repository.ReadBy(column, value, operator)
	if err != nil {
		return res, err
	}
	res = uc.buildBody(menu, false)

	return res, err
}

func (uc MenuUseCase) Edit(inputs *requests.EditMenuRequest) (err error) {
	repository := actions.NewMenuRepository(uc.DB)
	menuPermissionUc := MenuPermissionUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)
	var selectedMenuPermissionBody []viewmodel.MenuPermissionVm

	transaction, err := uc.DB.Begin()
	if err != nil {
		transaction.Rollback()

		return err
	}

	for _, input := range inputs.Menus {
		isExist, err := uc.isMenuExist(input.ID, "name", input.Name)
		if err != nil {
			transaction.Rollback()

			return err
		}
		if isExist {
			transaction.Rollback()

			return errors.New(messages.DataAlreadyExist)
		}

		body := viewmodel.MenuVm{
			ID:        input.ID,
			Name:      input.Name,
			Url:       input.Url,
			IsActive:  input.IsActive,
			UpdatedAt: now,
		}
		err = repository.Edit(body, transaction)
		if err != nil {
			transaction.Rollback()

			return err
		}
	}

	if len(inputs.SelectedPermissions) > 0 {
		for _, menuPermission := range inputs.SelectedPermissions {
			selectedMenuPermissionBody = append(selectedMenuPermissionBody, viewmodel.MenuPermissionVm{
				MenuID:     menuPermission.MenuID,
				ID:         menuPermission.ID,
				Permission: menuPermission.Permission,
			})
		}
	}
	err = menuPermissionUc.Store(selectedMenuPermissionBody, inputs.DeletedPermissions, transaction)
	if err != nil {
		transaction.Rollback()

		return err
	}
	transaction.Commit()

	return err
}

func (uc MenuUseCase) Add(inputs *requests.AddMenuRequest) (err error) {
	repository := actions.NewMenuRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	transaction, err := uc.DB.Begin()
	if err != nil {
		transaction.Rollback()

		return err
	}

	for _, input := range inputs.Menus {
		isExist, err := uc.isMenuExist("", "name", input.Name)
		if err != nil {
			transaction.Rollback()

			return err
		}
		if isExist {
			transaction.Rollback()

			return errors.New(messages.DataAlreadyExist)
		}

		menuID, err := uc.GetMenuID(input.ParentID)
		if err != nil {
			return err
		}
		fmt.Println(menuID)
		body := viewmodel.MenuVm{
			MenuID:    menuID,
			Name:      input.Name,
			Url:       input.Url,
			ParentID:  input.ParentID,
			IsActive:  true,
			CreatedAt: now,
			UpdatedAt: now,
		}
		_, err = repository.Add(body, transaction)
		if err != nil {
			transaction.Rollback()

			return err
		}
	}
	transaction.Commit()

	return err
}

func (uc MenuUseCase) Delete(ID string) (err error) {
	repository := actions.NewMenuRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	count, err := uc.countByPk(ID)
	if err != nil {
		return err
	}

	transaction, err := uc.DB.Begin()
	if err != nil {
		transaction.Rollback()

		return err
	}

	if count > 0 {
		err = repository.Delete(ID, now, now, transaction)
		if err != nil {
			transaction.Rollback()

			return err
		}

		count, _ = uc.countBy("", "parent_id", ID)
		if count > 0 {
			err = repository.DeleteChild(ID, now, now, transaction)
			if err != nil {
				transaction.Rollback()

				return err
			}
		}
	}
	transaction.Commit()

	return err
}

func (uc MenuUseCase) GetMenuID(parentID string) (res string, err error) {
	count, err := uc.countBy("", "parent_id", parentID)
	if err != nil {
		return res, err
	}

	res = fmt.Sprintf("%03d", count+1)

	return res, err
}

func (uc MenuUseCase) countBy(ID, column, value string) (res int, err error) {
	repository := actions.NewMenuRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc MenuUseCase) countByPk(ID string) (res int, err error) {
	repository := actions.NewMenuRepository(uc.DB)
	res, err = repository.CountByPk(ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc MenuUseCase) isMenuExist(ID, column, value string) (res bool, err error) {
	count, err := uc.countBy(ID, column, value)
	if err != nil {
		return res, err
	}

	return count > 0, err
}

func (uc MenuUseCase) buildBody(model models.Menu, isActive bool) viewmodel.MenuVm {
	var menuPermissions []viewmodel.SelectedMenuPermissionVm

	if model.Permissions.String != "" {
		menuPermissionArr := strings.Split(model.Permissions.String, ",")
		for _, permission := range menuPermissionArr {
			permissionArr := strings.Split(permission, ":")
			menuPermissions = append(menuPermissions, viewmodel.SelectedMenuPermissionVm{
				ID:         permissionArr[0],
				Permission: permissionArr[1],
			})
		}
	}

	return viewmodel.MenuVm{
		ID:              model.ID,
		MenuID:          model.MenuID,
		Name:            model.Name,
		Url:             model.Url,
		ParentID:        model.ParentID.String,
		IsActive:        model.IsActive,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		MenuPermissions: menuPermissions,
		ChildMenus:      uc.browseChild(model.ID, isActive),
	}
}
