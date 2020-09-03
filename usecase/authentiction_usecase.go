package usecase

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/google"
	"qibla-backend/helpers/hashing"
	"qibla-backend/helpers/messages"
	"qibla-backend/helpers/str"
	"qibla-backend/server/requests"
	"qibla-backend/usecase/viewmodel"
)

type AuthenticationUseCase struct {
	*UcContract
}

func (uc AuthenticationUseCase) UpdateSessionLogin(ID string) (res string, err error) {
	value := uuid.NewV4().String()
	exp := os.Getenv("SESSION_EXP")
	key := "session-" + ID
	resSession := viewmodel.UserSessionVm{}
	resSession.Session = value

	uc.RedisClient.StoreToRedistWithExpired(key, resSession, exp)

	return value, err
}

func (uc AuthenticationUseCase) GenerateJwtToken(jwePayload, email, session string) (token, refreshToken, expTokenAt, expRefreshTokenAt string, err error) {
	token, expTokenAt, err = uc.JwtCred.GetToken(session, email, jwePayload)
	if err != nil {
		return token, refreshToken, expTokenAt, expRefreshTokenAt, err
	}

	refreshToken, expRefreshTokenAt, err = uc.JwtCred.GetRefreshToken(session, email, jwePayload)
	if err != nil {
		return token, refreshToken, expTokenAt, expRefreshTokenAt, err
	}

	return token, refreshToken, expTokenAt, expRefreshTokenAt, err
}

func (uc AuthenticationUseCase) IsPasswordValid(ID, password, hashedPassword string) (err error) {
	isValid := hashing.CheckHashString(password, hashedPassword)
	if !isValid {
		uc.RedisClient.RemoveFromRedis("session-" + ID)

		return errors.New(messages.CredentialDoNotMatch)
	}

	return err
}

func (uc AuthenticationUseCase) Login(username, password string) (res viewmodel.UserJwtTokenVm, err error) {
	repository := actions.NewUserRepository(uc.DB)
	userUc := UserUseCase{UcContract: uc.UcContract}
	isPinSet := false

	isExist, err := userUc.IsUserNameExist("", username)
	if err != nil {
		return res, err
	}
	if !isExist {
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	user, err := repository.ReadBy("username", username)
	if err != nil {
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	err = uc.IsPasswordValid(user.ID, password, user.Password)
	if err != nil {
		return res, err
	}

	jwePayload, _ := uc.Jwe.GenerateJwePayload(user.ID)
	session, _ := uc.UpdateSessionLogin(user.ID)
	token, refreshToken, tokenExpiredAt, refreshTokenExpiredAt, err := uc.GenerateJwtToken(jwePayload, username, session)
	if err != nil {
		return res, err
	}
	if user.PIN.String != "" {
		isPinSet = true
	}
	res = viewmodel.UserJwtTokenVm{
		Token:           token,
		ExpTime:         tokenExpiredAt,
		RefreshToken:    refreshToken,
		ExpRefreshToken: refreshTokenExpiredAt,
		IsPinSet:        isPinSet,
	}

	return res, nil
}

func (uc AuthenticationUseCase) RegisterByEmail(input *requests.RegisterByMailRequest) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}

	//check is email exist
	isEmailExist, err := userUc.IsEmailExist("", input.Email)
	if err != nil {
		return err
	}
	if isEmailExist {
		return errors.New(messages.EmailAlreadyExist)
	}

	//init db transaction
	uc.TX, err = uc.DB.Begin()
	if err != nil {
		uc.TX.Rollback()

		return err
	}

	//add user jamaah
	jamaahUc := JamaahUseCase{UcContract: uc.UcContract}
	err = jamaahUc.Add(input.Name, "jamaah", input.Email, input.Password)
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc AuthenticationUseCase) ForgotPassword(email string) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	jamaahUc := JamaahUseCase{UcContract: uc.UcContract}

	count, err := userUc.CountBy("", "email", email)
	if err != nil {
		fmt.Print(1)
		return err
	}

	if count > 0 {
		jamaah, err := jamaahUc.ReadBy("email", email)
		if err != nil {
			fmt.Print(1)
			return err
		}
		password := str.RandomString(6)
		err = jamaahUc.EditPassword(jamaah.ID, password)
		if err != nil {
			fmt.Print(1)
			return err
		}

		uc.GoMailConfig.SendGoMail(email, "New Password", `<h1>`+password+`</h1>`)
	}

	return nil
}

func (uc AuthenticationUseCase) RegisterByGmail(input *requests.RegisterByGmailRequest) (err error) {
	res ,err := google.GetGoogleProfile(input.Token)
	if err != nil {
		return errors.New(messages.CredentialDoNotMatch)
	}

	fmt.Print(res)

	return nil
}
