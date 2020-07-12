package usecase

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"qibla-backend/db/repositories/actions"
	"qibla-backend/helpers/hashing"
	"qibla-backend/helpers/messages"
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

func (uc AuthenticationUseCase) IsPasswordValid(ID,password,hashedPassword string) (err error){
	isValid := hashing.CheckHashString(password,hashedPassword)
	if !isValid {
		uc.RedisClient.RemoveFromRedis("session-"+ID)

		return errors.New(messages.CredentialDoNotMatch)
	}

	return err
}

func (uc AuthenticationUseCase) Login(username, password string) (res viewmodel.UserJwtTokenVm, err error) {
	repository := actions.NewUserRepository(uc.DB)
	userUc := UserUseCase{UcContract: uc.UcContract}

	isExist, err := userUc.IsUserNameExist("", username)
	if err != nil {
		return res, err
	}
	if !isExist {
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	user, err := repository.ReadBy("username",username)
	if err != nil {
		fmt.Println(err)
		return res,errors.New(messages.CredentialDoNotMatch)
	}

	err = uc.IsPasswordValid(user.ID,password,user.Password)
	if err != nil {
		return res,err
	}

	jwePayload, _ := uc.Jwe.GenerateJwePayload(user.ID)
	session, _ := uc.UpdateSessionLogin(user.ID)
	token, refreshToken, tokenExpiredAt, refreshTokenExpiredAt, err := uc.GenerateJwtToken(jwePayload, username, session)
	if err != nil {
		return res, err
	}
	res = viewmodel.UserJwtTokenVm{
		Token:           token,
		ExpTime:         tokenExpiredAt,
		RefreshToken:    refreshToken,
		ExpRefreshToken: refreshTokenExpiredAt,
	}

	return res, nil
}
