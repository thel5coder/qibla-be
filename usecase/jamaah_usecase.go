package usecase

import (
	"errors"
	"qibla-backend/pkg/hashing"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/requests"
)

type JamaahUseCase struct {
	*UcContract
}

func (uc JamaahUseCase) Edit(input *requests.EditProfileRequest, ID string) (err error) {
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	userUc := UserUseCase{UcContract: uc.UcContract}
	user, err := userUc.ReadBy("u.id", ID)
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	isEmailExist, err := userUc.IsEmailExist(ID, input.Email)
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	if isEmailExist {
		return errors.New(messages.EmailAlreadyExist)
	}

	hashedPassword, _ := hashing.HashAndSalt(input.Password)
	err = userUc.Edit(ID, input.FullName, input.Email, input.Email, input.MobilePhone, user.Role.ID, hashedPassword, input.ProfilePictureID, true, false)
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
	res, err = userUc.Add(name, email, email, mobilePhone, role.ID, encryptedPassword, "", false, false)
	if err != nil {
		return res, err
	}

	return res, nil
}
