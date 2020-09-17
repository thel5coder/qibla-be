package usecase

import (
	"errors"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/hashing"
	"qibla-backend/helpers/messages"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
	"time"
)

type JamaahUseCase struct {
	*UcContract
}

func (uc JamaahUseCase) ReadBy(column, value string) (res viewmodel.JamaahVm, err error) {
	repository := actions.NewUserRepository(uc.DB)
	isPinSet := false
	user, err := repository.ReadBy(column, value)
	if err != nil {
		return res, err
	}

	if user.PIN.String != "" {
		isPinSet = true
	}

	res = viewmodel.JamaahVm{
		ID:             user.ID,
		UserName:       user.UserName,
		Email:          user.Email,
		Name:           user.Name.String,
		MobilePhone:    user.MobilePhone.String,
		ProfilePicture: user.ProfilePicture.String,
		RoleID:         user.RoleID.String,
		RoleName:       user.RoleModel.Name,
		IsActive:       user.IsActive,
		IsPinSet:       isPinSet,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}

	return res, err
}

func (uc JamaahUseCase) Edit(input *requests.EditProfileRequest, ID string) (err error) {
	uc.TX,err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	user,err := uc.ReadBy("id",ID)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	userUc := UserUseCase{UcContract: uc.UcContract}
	isEmailExist, err := userUc.IsEmailExist(ID, input.Email)
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	if isEmailExist {
		return errors.New(messages.EmailAlreadyExist)
	}

	hashedPassword,_ := hashing.HashAndSalt(input.Password)
	err = userUc.Edit(ID, input.FullName, input.Email, input.Email, input.MobilePhone, user.RoleID, hashedPassword	, true, false)
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return err
}

func (uc JamaahUseCase) Add(name, roleSlug, email, password, mobilePhone string) (res string, err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	roleUc := RoleUseCase{UcContract: uc.UcContract}

	role, err := roleUc.ReadBy("slug", roleSlug)
	if err != nil {
		return res, err
	}

	encryptedPassword, _ := hashing.HashAndSalt(password)
	res, err = userUc.Add(name, email, email, mobilePhone, role.ID, encryptedPassword, true, false)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc JamaahUseCase) EditPassword(ID, password string) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	encryptedPassword, _ := hashing.HashAndSalt(password)
	_, err = repository.EditPassword(ID, encryptedPassword, now)
	if err != nil {
		return err
	}

	return nil
}

func (uc JamaahUseCase) EditPin(ID string, PIN string) (err error) {
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC().Format(time.RFC3339)

	encryptedPin, _ := hashing.HashAndSalt(PIN)
	_, err = repository.EditPIN(ID, encryptedPin, now)
	if err != nil {
		return err
	}

	return nil
}
