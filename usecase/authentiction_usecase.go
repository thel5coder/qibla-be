package usecase

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"qibla-backend/helpers/facebook"
	"qibla-backend/helpers/google"
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

func (uc AuthenticationUseCase) Login(username, password, fcmDeviceToken string) (res viewmodel.UserJwtTokenVm, err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	isPinSet := false

	isExist, err := userUc.IsUserNameExist("", username)
	if err != nil {
		return res, err
	}
	if !isExist {
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	user, err := userUc.ReadBy("u.username", username)
	if err != nil {
		fmt.Println(err.Error())
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	isPasswordValid, err := userUc.IsPasswordValid(user.ID, password)
	if err != nil {
		return res, err
	}
	if !isPasswordValid {
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	err = userUc.EditFcmDeviceToken(user.ID, fcmDeviceToken)
	if err != nil {
		return res, err
	}

	jwePayload, _ := uc.Jwe.GenerateJwePayload(user.ID)
	session, _ := uc.UpdateSessionLogin(user.ID)
	token, refreshToken, tokenExpiredAt, refreshTokenExpiredAt, err := uc.GenerateJwtToken(jwePayload, username, session)
	if err != nil {
		return res, err
	}
	if user.PIN != "" {
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
	_, err = jamaahUc.Add(input.Name, "guest", input.Email, input.Password, "")
	if err != nil {
		uc.TX.Rollback()

		return err
	}
	uc.TX.Commit()

	return nil
}

func (uc AuthenticationUseCase) ForgotPassword(email string) (err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}

	count, err := userUc.CountBy("", "email", email)
	if err != nil {
		return err
	}

	if count > 0 {
		user, err := userUc.ReadBy("u.email", email)
		if err != nil {
			return err
		}
		password := str.RandomString(6)
		err = userUc.EditPassword(user.ID, password)
		if err != nil {
			return err
		}

		uc.GoMailConfig.SendGoMail(email, "New Password", `<h1>`+password+`</h1>`)
	}

	return nil
}

func (uc AuthenticationUseCase) RegisterByOauth(input *requests.RegisterByOauthRequest) (res viewmodel.UserJwtTokenVm, err error) {
	var profile map[string]interface{}
	var email, name string

	if input.Type == "gmail" {
		profile, err = google.GetGoogleProfile(input.Token)
		if err != nil {
			fmt.Println(err.Error())
			return res, errors.New(messages.CredentialDoNotMatch)
		}
		email = profile["email"].(string)
		name = profile["name"].(string)
	} else {
		profile, err = facebook.GetFacebookProfile(input.Token)
		if err != nil {
			fmt.Println(err.Error())
			return res, err
		}
		fmt.Println(profile)
		email = profile["email"].(string)
		name = profile["name"].(string)
	}

	return uc.registerUserByOauth(email, name, input.FcmDeviceToken)
}

func (uc AuthenticationUseCase) registerUserByOauth(email, name, fcmDeviceToken string) (res viewmodel.UserJwtTokenVm, err error) {
	var user viewmodel.UserVm
	var userID string
	userUc := UserUseCase{UcContract: uc.UcContract}
	jamaahUc := JamaahUseCase{UcContract: uc.UcContract}

	//count email by email profile
	count, err := userUc.CountBy("", "email", email)
	if err != nil {
		fmt.Print(1)
		return res, err
	}
	if count > 0 {
		user, err = userUc.ReadBy("u.email", email)
		if err != nil {
			fmt.Println(2)
			return res, err
		}
		userID = user.ID
	} else {
		uc.TX, err = uc.DB.Begin()
		if err != nil {
			uc.TX.Rollback()

			return res, err
		}
		jamaahUc.TX = uc.TX

		//add user jamaah
		password := str.RandomString(6)
		userID, err = jamaahUc.Add(name, "guest", email, password, "")
		if err != nil {
			fmt.Println(3)
			uc.TX.Rollback()

			return res, err
		}
		uc.TX.Commit()
	}

	err = userUc.EditFcmDeviceToken(user.ID, fcmDeviceToken)
	if err != nil {
		return res, err
	}

	jwePayload, _ := uc.Jwe.GenerateJwePayload(userID)
	session, _ := uc.UpdateSessionLogin(userID)
	token, refreshToken, tokenExpiredAt, refreshTokenExpiredAt, err := uc.GenerateJwtToken(jwePayload, email, session)
	if err != nil {
		return res, err
	}

	res = viewmodel.UserJwtTokenVm{
		Token:           token,
		ExpTime:         tokenExpiredAt,
		RefreshToken:    refreshToken,
		ExpRefreshToken: refreshTokenExpiredAt,
		IsPinSet:        user.IsPINSet,
	}

	return res, nil
}
