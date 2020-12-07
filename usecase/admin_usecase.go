package usecase

import (
	"errors"
	"qibla-backend/pkg/hashing"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
)

type AdminUseCase struct {
	*UcContract
}

func (uc AdminUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.AdminVm, pagination viewmodel.PaginationVm, err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	users, pagination, err := userUc.Browse(true, search, order, sort, page, limit)
	if err != nil {
		return res, pagination, err
	}

	for _, user := range users {
		temp, err := uc.buildBody(user)
		if err != nil {
			return res, pagination, err
		}

		res = append(res, temp)
	}

	return res, pagination, err
}

func (uc AdminUseCase) ReadBy(column, value, operator string) (res viewmodel.AdminVm, err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	user, err := userUc.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	res, err = uc.buildBody(user)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc AdminUseCase) isExist(ID string, input *requests.UserRequest) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}

	//check is email exist
	isEmailExist, err := userUc.IsEmailExist(ID, input.Email)
	if err != nil {
		return err
	}
	if isEmailExist {
		return errors.New(messages.EmailAlreadyExist)
	}

	//check is username exist
	isUserNameExist, err := userUc.IsUserNameExist(ID, input.UserName)
	if err != nil {
		return err
	}
	if isUserNameExist {
		return errors.New(messages.UserNameExist)
	}

	return nil
}

func (uc AdminUseCase) Edit(ID string, input *requests.UserRequest) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	password := ""

	//check is email or username exist
	err = uc.isExist(ID, input)
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

	//add menu user permissions
	menuUserUc := MenuUserUseCase{UcContract: uc.UcContract}
	err = menuUserUc.Store(ID, input.MenuUsers)
	if err != nil {
		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc AdminUseCase) Add(input *requests.UserRequest) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}

	//check is email or username exist
	err = uc.isExist("", input)
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
	userID, err := userUc.Add(input.UserName, input.UserName, input.Email, "", input.RoleID, password, input.ProfilePictureID, true, true)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	//add menu user permissions
	menuUserUc := MenuUserUseCase{UcContract: uc.UcContract}
	err = menuUserUc.Store(userID, input.MenuUsers)
	if err != nil {
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
	//menuPermissionUserUc := MenuUserPermissionUseCase{UcContract: uc.UcContract}

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

		////delete user menu permission
		//err = menuPermissionUserUc.DeleteByUser(ID, uc.TX)
		//if err != nil {
		//	uc.TX.Rollback()
		//
		//	return err
		//}
	}
	uc.TX.Commit()

	return nil
}

func (uc AdminUseCase) buildBody(userVm viewmodel.UserVm) (res viewmodel.AdminVm, err error) {
	menuUserUc := MenuUserUseCase{UcContract: uc.UcContract}
	menuUsers, err := menuUserUc.BrowseBy("user_id", userVm.ID, "=")
	if err != nil {
		return res, err
	}

	var accessMenuTemp []viewmodel.AdminUserAccessMenuVm
	for _, menuUser := range menuUsers {
		accessMenuTemp = append(accessMenuTemp, viewmodel.AdminUserAccessMenuVm{
			MenuID:     menuUser.MenuID,
			Permission: menuUser.Permissions,
		})
	}

	res = viewmodel.AdminVm{
		User:       userVm,
		MenuAccess: accessMenuTemp,
	}

	return res, nil
}
