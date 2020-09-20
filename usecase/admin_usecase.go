package usecase

import (
	"errors"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/hashing"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
)

type AdminUseCase struct {
	*UcContract
}

func (uc AdminUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.AdminVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewUserRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	users, count, err := repository.BrowseUserAdminPanel(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, user := range users {
		res = append(res, viewmodel.AdminVm{
			ID:       user.ID,
			UserName: user.UserName,
			Email:    user.Email.String,
			IsActive: user.IsActive,
			Role: viewmodel.RoleVm{
				ID:        user.RoleModel.ID,
				Name:      user.RoleModel.Name,
				CreatedAt: user.RoleModel.CreatedAt,
				UpdatedAt: user.RoleModel.UpdatedAt,
				DeletedAt: user.RoleModel.DeletedAt.String,
			},
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc AdminUseCase) ReadBy(column, value string) (res viewmodel.AdminVm, err error) {
	repository := actions.NewUserRepository(uc.DB)
	menuPermissionUserUc := MenuPermissionUserUseCase{UcContract: uc.UcContract}
	var permissions []viewmodel.MenuPermissionUserVm

	user, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	menuPermissionsUsers, err := menuPermissionUserUc.Browse(user.ID)
	if err != nil {
		return res, err
	}
	for _, menuPermissionsUser := range menuPermissionsUsers {
		permissions = append(permissions, menuPermissionsUser)
	}

	res = viewmodel.AdminVm{
		ID:          user.ID,
		UserName:    user.UserName,
		Email:       user.Email.String,
		MobilePhone: user.MobilePhone.String,
		IsActive:    user.IsActive,
		Role: viewmodel.RoleVm{
			ID:        user.RoleModel.ID,
			Name:      user.RoleModel.Name,
			CreatedAt: user.RoleModel.CreatedAt,
			UpdatedAt: user.RoleModel.UpdatedAt,
			DeletedAt: user.RoleModel.DeletedAt.String,
		},
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt.String,
		MenuPermissions: permissions,
	}

	return res, err
}

func (uc AdminUseCase) isExist(input *requests.AdminRequest) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}

	//check is email exist
	isEmailExist, err := userUc.IsEmailExist("", input.Email)
	if err != nil {
		return err
	}
	if isEmailExist {
		return errors.New(messages.EmailAlreadyExist)
	}

	//check is username exist
	isUserNameExist, err := userUc.IsUserNameExist("", input.UserName)
	if err != nil {
		return err
	}
	if isUserNameExist {
		return errors.New(messages.UserNameExist)
	}

	return nil
}

func (uc AdminUseCase) Edit(ID string, input *requests.AdminRequest) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	password := ""

	//check is email or username exist
	err = uc.isExist(input)
	if err != nil {
		return err
	}

	//init db transaction
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	//edit user
	userUc.TX = uc.TX
	if input.Password != "" {
		password, _ = hashing.HashAndSalt(input.Password)
	}
	err = userUc.Edit(ID, input.UserName, input.UserName, input.Email, "", input.RoleID, password, "", input.IsActive, true)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	//edit menu permission user
	menuPermissionUserUc := MenuPermissionUserUseCase{UcContract: uc.UcContract}
	err = menuPermissionUserUc.Store(ID, input.MenuPermissions, input.DeletedMenuPermissions, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc AdminUseCase) Add(input *requests.AdminRequest) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}

	//check is email or username exist
	err = uc.isExist(input)
	if err != nil {
		return err
	}

	//init db transaction
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	userUc.TX = uc.TX

	//add user admin
	password, _ := hashing.HashAndSalt(input.Password)
	userID, err := userUc.Add(input.UserName, input.UserName, input.Email, "", input.RoleID, password, true, true)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	//add menu permissions
	menuPermissionUserUc := MenuPermissionUserUseCase{UcContract: uc.UcContract}
	err = menuPermissionUserUc.Store(userID, input.MenuPermissions, input.DeletedMenuPermissions, uc.TX)
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc AdminUseCase) Delete(ID string) (err error) {
	//init db transaction
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	userUc := UserUseCase{UcContract: uc.UcContract}
	menuPermissionUserUc := MenuPermissionUserUseCase{UcContract: uc.UcContract}

	//validate is data exist
	count, err := userUc.CountByPk(ID)
	if err != nil {
		uc.TX.Rollback()

		return errors.New(messages.DataNotFound)
	}

	//if count delete user and user menu permission
	if count > 0 {
		//delete user
		err = userUc.Delete(ID)
		if err != nil {
			uc.TX.Rollback()

			return err
		}

		//delete user menu permission
		err = menuPermissionUserUc.DeleteByUser(ID, uc.TX)
		if err != nil {
			uc.TX.Rollback()

			return err
		}
	}
	uc.TX.Rollback()

	return nil
}

func (uc AdminUseCase) IsPasswordValid(userID, password string) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	user, err := repository.ReadBy("id", userID)
	if err != nil {
		return errors.New(messages.CredentialDoNotMatch)
	}

	isValid := hashing.CheckHashString(password, user.Password)
	if !isValid {
		return errors.New(messages.CredentialDoNotMatch)
	}

	return nil
}
