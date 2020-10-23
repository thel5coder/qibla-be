package usecase

import (
	"errors"
	"fmt"
	"qibla-backend/helpers/hashing"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
)

type AdminUseCase struct {
	*UcContract
}

func (uc AdminUseCase) isExist(input *requests.UserRequest) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}

	//check is email exist
	isEmailExist, err := userUc.IsEmailExist("", input.Email)
	if err != nil {
		fmt.Println(4)
		return err
	}
	if isEmailExist {
		fmt.Println(5)
		return errors.New(messages.EmailAlreadyExist)
	}

	//check is username exist
	isUserNameExist, err := userUc.IsUserNameExist("", input.UserName)
	if err != nil {
		return err
	}
	if isUserNameExist {
		fmt.Println(6)
		return errors.New(messages.UserNameExist)
	}

	return nil
}

func (uc AdminUseCase) Edit(ID string, input *requests.UserRequest) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	password := ""

	//check is email or username exist
	err = uc.isExist(input)
	if err != nil {
		fmt.Println(1)
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
		fmt.Println(2)
		uc.TX.Rollback()

		return err
	}

	//edit menu permission user
	menuPermissionUserUc := MenuPermissionUserUseCase{UcContract: uc.UcContract}
	err = menuPermissionUserUc.Store(ID, input.MenuPermissions, input.DeletedMenuPermissions, uc.TX)
	if err != nil {
		fmt.Println(3)
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc AdminUseCase) Add(input *requests.UserRequest) (err error) {
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
	userID, err := userUc.Add(input.UserName, input.UserName, input.Email, "", input.RoleID, password, input.ProfilePictureID,true, true)
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
	uc.TX.Commit()

	return nil
}
