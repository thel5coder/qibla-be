package usecase

import (
	"qibla-backend/db/models"
	"qibla-backend/db/repositories/actions"
	functionCaller "qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/hashing"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type UserUseCase struct {
	*UcContract
}

func (uc UserUseCase) Browse(isAdminPanel bool, search, order, sort string, page, limit int) (res []viewmodel.UserVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewUserRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	users, count, err := repository.Browse(isAdminPanel, search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, user := range users {
		res = append(res, uc.buildBody(user))
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc UserUseCase) ReadBy(column, value string) (res viewmodel.UserVm, err error) {
	repository := actions.NewUserRepository(uc.DB)
	user, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}
	res = uc.buildBody(user)

	return res, err
}

func (uc UserUseCase) Edit(ID, name, userName, email, mobilePhone, roleID, password, profilePicture string, isActive, isAdminPanel bool) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.UserVm{
		ID:           ID,
		Name:         name,
		UserName:     userName,
		Email:        email,
		MobilePhone:  mobilePhone,
		IsActive:     isActive,
		IsAdminPanel: isAdminPanel,
		Role:         viewmodel.RoleVm{ID: roleID},
		File:         viewmodel.FileVm{ID: profilePicture},
		UpdatedAt:    now.Format(time.RFC3339),
	}
	err = repository.Edit(body, password, uc.TX)
	if err != nil {
		return err
	}

	return err
}

func (uc UserUseCase) EditUserName(ID, userName string) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	err = repository.EditUserName(ID, userName, now, uc.TX)

	return err
}

func (uc UserUseCase) EditPassword(ID, password string) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	encryptedPassword, _ := hashing.HashAndSalt(password)
	_, err = repository.EditPassword(ID, encryptedPassword, now)
	if err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) EditPin(ID, PIN string) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	encryptedPin, _ := hashing.HashAndSalt(PIN)
	_, err = repository.EditPIN(ID, encryptedPin, now)
	if err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) EditFingerPrintStatus(ID string, fingerPrintStatus bool) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)
	_, err = repository.EditFingerPrintStatus(ID, now, fingerPrintStatus)
	if err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) EditFcmDeviceToken(ID, fcmDeviceToken string) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	_, err = repository.EditFcmDeviceToken(ID, fcmDeviceToken, now)
	if err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) EditIsActiveStatus(email string, isActive bool) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	_, err = repository.EditIsActiveStatus(email, now, isActive)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(), functionCaller.PrintFuncName(),"query-user-editIsActiveStatus")
		return err
	}

	return nil
}

func (uc UserUseCase) Add(name, userName, email, mobilePhone, roleID, password, profilePicture string, isActive, isAdminPanel bool) (res string, err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.UserVm{
		Name:         name,
		UserName:     userName,
		Email:        email,
		MobilePhone:  mobilePhone,
		IsActive:     isActive,
		Role:         viewmodel.RoleVm{ID: roleID},
		File:         viewmodel.FileVm{ID: profilePicture},
		IsAdminPanel: isAdminPanel,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	res, err = repository.Add(body, password, uc.TX)
	if err != nil {
		return res, err
	}

	return res, err
}

func (uc UserUseCase) Delete(ID string) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC()

	err = repository.Delete(ID, now.Format(time.RFC3339), now.Format(time.RFC3339), uc.TX)
	if err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) CountBy(ID, column, value string) (res int, err error) {
	repository := actions.NewUserRepository(uc.DB)
	res, err = repository.CountBy(ID, column, value)

	return res, err
}

func (uc UserUseCase) CountByPk(ID string) (res int, err error) {
	repository := actions.NewUserRepository(uc.DB)
	res, err = repository.CountByPk(ID)

	return res, err
}

func (uc UserUseCase) IsUserNameExist(ID, userName string) (res bool, err error) {
	count, err := uc.CountBy(ID, "username", userName)
	if err != nil {
		return res, err
	}

	return count > 0, err
}

func (uc UserUseCase) IsEmailExist(ID, email string) (res bool, err error) {
	count, err := uc.CountBy(ID, "email", email)
	if err != nil {
		return res, err
	}

	return count > 0, err
}

func (uc UserUseCase) IsPasswordValid(ID, password string) (res bool, err error) {
	repository := actions.NewUserRepository(uc.DB)
	user, err := repository.ReadBy("u.id", ID)
	if err != nil {
		return res, err
	}
	res = hashing.CheckHashString(password, user.Password)

	return res, err
}

func (uc UserUseCase) buildBody(model models.User) (res viewmodel.UserVm) {
	var isPINSet = false
	menuPermissionUserUc := MenuUserPermissionUseCase{UcContract: uc.UcContract}
	var permissions []viewmodel.MenuUserPermissionVm

	menuPermissionsUsers, _ := menuPermissionUserUc.Browse(model.ID)
	for _, menuPermissionsUser := range menuPermissionsUsers {
		permissions = append(permissions, menuPermissionsUser)
	}

	fileUc := FileUseCase{UcContract: uc.UcContract}
	file, _ := fileUc.ReadByPk(model.ProfilePictureID.String)

	if model.PIN.String != "" {
		isPINSet = true
	}
	res = viewmodel.UserVm{
		ID:               model.ID,
		UserName:         model.UserName,
		Name:             model.Name.String,
		Email:            model.Email.String,
		MobilePhone:      model.MobilePhone.String,
		PIN:              model.PIN.String,
		IsActive:         model.IsActive,
		IsAdminPanel:     model.IsAdminPanel.Bool,
		IsPINSet:         isPINSet,
		IsFingerPrintSet: model.IsFingerprintSet.Bool,
		OdooUserID:       model.OdooUserID.Int32,
		FcmDeviceToken:   model.FcmDeviceToken.String,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
		Role: viewmodel.RoleVm{
			ID:   model.RoleModel.ID,
			Name: model.RoleModel.Name,
			Slug: model.RoleModel.Slug,
		},
		File: file,
	}

	return res
}
