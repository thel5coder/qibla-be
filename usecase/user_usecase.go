package usecase

import (
	"qibla-backend/db/repositories/actions"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type UserUseCase struct {
	*UcContract
}

func (uc UserUseCase) Edit(ID, name, userName, email, mobilePhone, roleID, password string, isActive, isAdminPanel bool) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC()

	body := viewmodel.UserVm{
		ID:           ID,
		Name:         name,
		UserName:     userName,
		Email:        email,
		MobilePhone:  mobilePhone,
		IsActive:     isActive,
		RoleID:       roleID,
		IsAdminPanel: isAdminPanel,
		UpdatedAt:    now.Format(time.RFC3339),
	}
	err = repository.Edit(body, password, uc.TX)
	if err != nil {
		return err
	}

	return err
}

func (uc UserUseCase) Add(name, userName, email, mobilePhone, roleID, password string, isActive, isAdminPanel bool) (res string, err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	body := viewmodel.UserVm{
		Name:         name,
		UserName:     userName,
		Email:        email,
		MobilePhone:  mobilePhone,
		IsActive:     isActive,
		RoleID:       roleID,
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
