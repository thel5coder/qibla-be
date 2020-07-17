package usecase

import (
	"errors"
	"fmt"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"strconv"
	"time"
)

type MenuUseCase struct {
	*UcContract
}

func (uc MenuUseCase) Browse(parentID, search, order, sort string, page, limit int) (res []viewmodel.MenuVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewMenuRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	menus, count, err := repository.Browse(parentID, search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, menu := range menus {
		res = append(res, viewmodel.MenuVm{
			ID:        menu.ID,
			MenuID:    menu.MenuID,
			Name:      menu.Name,
			Url:       menu.Url,
			ParentID:  menu.ParentID.String,
			IsActive:  menu.IsActive,
			CreatedAt: menu.CreatedAt,
			UpdatedAt: menu.UpdatedAt,
			DeletedAt: menu.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc MenuUseCase) ReadBy(column, value string) (res viewmodel.MenuVm, err error) {
	repository := actions.NewMenuRepository(uc.DB)
	menuPermissionUc := MenuPermissionUseCase{UcContract: uc.UcContract}
	var permissions []string

	menu, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	rootMenuPermissions, err := menuPermissionUc.Browse(menu.ID)
	if err != nil {
		return res, err
	}
	for _, menuPermission := range rootMenuPermissions {
		permissions = append(permissions, menuPermission.ID)
	}

	childMenus, _, err := uc.Browse(menu.ID, "", "", "", 0, 0)
	if err != nil {
		return res, err
	}
	for i:=0;i<len(childMenus);i++{
		childMenuPermissions,err := menuPermissionUc.Browse(childMenus[i].ID)
		if err != nil {
			return res,err
		}
		for _,childMenuPermission := range childMenuPermissions{
			childMenus[i].MenuPermissions = append(childMenus[i].MenuPermissions,childMenuPermission.ID)
		}
	}
	fmt.Println(childMenus)

	res = viewmodel.MenuVm{
		ID:              menu.ID,
		MenuID:          menu.MenuID,
		Name:            menu.Name,
		Url:             menu.Url,
		ParentID:        menu.ParentID.String,
		IsActive:        menu.IsActive,
		CreatedAt:       menu.CreatedAt,
		UpdatedAt:       menu.UpdatedAt,
		DeletedAt:       menu.DeletedAt.String,
		MenuPermissions: permissions,
		ChildMenus:      childMenus,
	}

	return res, err
}

func (uc MenuUseCase) ReadByPk(ID string) (res viewmodel.MenuVm, err error) {
	res, err = uc.ReadBy("id", ID)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc MenuUseCase) Edit(inputs *requests.EditMenuRequest) (err error) {
	repository := actions.NewMenuRepository(uc.DB)
	menuPermissionUc := MenuPermissionUseCase{UcContract: uc.UcContract}
	now := time.Now().UTC().Format(time.RFC3339)

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
		var selectedMenuPermissionBody []viewmodel.MenuPermissionVm
		for _, menuPermission := range inputs.SelectedPermissions {
			selectedMenuPermissionBody = append(selectedMenuPermissionBody, viewmodel.MenuPermissionVm{
				MenuID:     menuPermission.MenuID,
				ID:         menuPermission.ID,
				Permission: menuPermission.Permission,
			})
		}

		err = menuPermissionUc.Store(selectedMenuPermissionBody, inputs.DeletedPermissions, transaction)
		if err != nil {
			transaction.Rollback()

			return err
		}
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
			fmt.Println("isexist")
			transaction.Rollback()

			return err
		}
		if isExist {
			fmt.Println("already exist")
			transaction.Rollback()

			return errors.New(messages.DataAlreadyExist)
		}

		body := viewmodel.MenuVm{
			MenuID:    input.MenuID,
			Name:      input.Name,
			Url:       input.Url,
			ParentID:  input.ParentID,
			IsActive:  true,
			CreatedAt: now,
			UpdatedAt: now,
		}
		_, err = repository.Add(body, transaction)
		if err != nil {
			fmt.Print("add")
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

	if count > 9 {
		res = `0` + strconv.Itoa(count+1)
	} else if count > 99 {
		res = string(count)
	} else {
		res = `00` + strconv.Itoa(count+1)
	}

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
