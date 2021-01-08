package usecase

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"os"
	"qibla-backend/pkg/facebook"
	functionCaller "qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/google"
	"qibla-backend/pkg/logruslogger"
	"qibla-backend/pkg/messages"
	"qibla-backend/pkg/str"
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

// Login ...
func (uc AuthenticationUseCase) Login(username, password, fcmDeviceToken string) (res viewmodel.UserJwtTokenVm, err error) {
	userUc := UserUseCase{UcContract: uc.UcContract}
	isPinSet := false

	isExist, err := userUc.IsUserNameExist("", username)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "uc-user-isUserNameExist")
		return res, err
	}
	if !isExist {
		logruslogger.Log(logruslogger.WarnLevel, messages.CredentialDoNotMatch, functionCaller.PrintFuncName(), "uc-user-isUserNameExist")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	user, err := userUc.ReadBy("u.username", username)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "uc-user-readByUserName")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	isPasswordValid, err := userUc.IsPasswordValid(user.ID, password)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "uc-user-isPasswordValid")
		return res, err
	}
	if !isPasswordValid {
		logruslogger.Log(logruslogger.WarnLevel, messages.CredentialDoNotMatch, functionCaller.PrintFuncName(), "uc-user-isPasswordValid")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	err = userUc.EditFcmDeviceToken(user.ID, fcmDeviceToken)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "uc-user-EditFcmDeviceToken")
		return res, err
	}

	jwePayload, _ := uc.Jwe.GenerateJwePayload(user.ID)
	session, _ := uc.UpdateSessionLogin(user.ID)
	token, refreshToken, tokenExpiredAt, refreshTokenExpiredAt, err := uc.GenerateJwtToken(jwePayload, username, session)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "uc-generateJwtToken")
		return res, err
	}
	if user.PIN != "" {
		isPinSet = true
	}
	res = viewmodel.UserJwtTokenVm{
		Token:            token,
		ExpTime:          tokenExpiredAt,
		RefreshToken:     refreshToken,
		ExpRefreshToken:  refreshTokenExpiredAt,
		IsPinSet:         isPinSet,
		IsFingerPrintSet: user.IsFingerPrintSet,
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

	err = uc.generateActivationKey(input.Email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "uc-authentication-generateActivationKey")
		return nil
	}

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
			return res, errors.New(messages.CredentialDoNotMatch)
		}
		email = profile["email"].(string)
		name = profile["name"].(string)
	} else {
		profile, err = facebook.GetFacebookProfile(input.Token)
		if err != nil {
			return res, err
		}
		email = profile["email"].(string)
		name = profile["name"].(string)
	}

	return uc.registerUserByOauth(email, name, input.FcmDeviceToken)
}

func (uc AuthenticationUseCase) registerUserByOauth(email, name, fcmDeviceToken string) (res viewmodel.UserJwtTokenVm, err error) {
	var user viewmodel.UserVm
	var userID string
	var isPinSet bool
	var isFingerPrint bool
	userUc := UserUseCase{UcContract: uc.UcContract}
	jamaahUc := JamaahUseCase{UcContract: uc.UcContract}

	//count email by email profile
	count, err := userUc.CountBy("", "email", email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "count-user")
		return res, err
	}
	if count > 0 {
		user, err = userUc.ReadBy("u.email", email)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "read-by-user")
			return res, err
		}
		userID = user.ID
		if user.PIN != "" {
			isPinSet = true
		}
		isFingerPrint = user.IsFingerPrintSet
	} else {
		uc.TX, err = uc.DB.Begin()
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "init-transaction")
			uc.TX.Rollback()

			return res, err
		}
		jamaahUc.TX = uc.TX

		//add user jamaah
		password := str.RandomString(6)
		userID, err = jamaahUc.Add(name, "guest", email, password, "")
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "add-jamaah")
			uc.TX.Rollback()

			return res, err
		}
		uc.TX.Commit()
		isPinSet = false
		isFingerPrint = user.IsFingerPrintSet
	}

	//edit fcm token
	err = userUc.EditFcmDeviceToken(userID, fcmDeviceToken)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "edit-fcm")
		return res, err
	}

	jwePayload, _ := uc.Jwe.GenerateJwePayload(userID)
	session, _ := uc.UpdateSessionLogin(userID)
	token, refreshToken, tokenExpiredAt, refreshTokenExpiredAt, err := uc.GenerateJwtToken(jwePayload, email, session)
	if err != nil {
		return res, err
	}

	res = viewmodel.UserJwtTokenVm{
		Token:            token,
		ExpTime:          tokenExpiredAt,
		RefreshToken:     refreshToken,
		ExpRefreshToken:  refreshTokenExpiredAt,
		IsPinSet:         isPinSet,
		IsFingerPrintSet: isFingerPrint,
	}

	return res, nil
}

//generate activation key
func (uc AuthenticationUseCase) generateActivationKey(email string) (err error) {
	code, err := uc.GenerateKeyFromAES(email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functionCaller.PrintFuncName(), "uc-contract-generateKeyFromAes")
		return err
	}
	url := os.Getenv("APP_FE_URL") + `/activation?code` + code
	body := `klik link ini untuk melakukan aktivasi ` + url
	uc.GoMailConfig.SendGoMail(email, "Aktivasi User Qibla", body)

	return nil
}

//activation user from code/key
func(uc AuthenticationUseCase) ActivationUserByCode(code string) (err error){
	callBack := map[string]interface{}{}
	err = uc.RedisClient.GetFromRedis("code-"+code,&callBack)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,messages.InvalidKey,functionCaller.PrintFuncName(),"uc-redis-getCodeActivation")
		return errors.New(messages.InvalidKey)
	}

	userUc := UserUseCase{UcContract:uc.UcContract}
	email := callBack["email"].(string)
	err = userUc.EditIsActiveStatus(email,true)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functionCaller.PrintFuncName(),"uc-user-editIsActiveStatus")
		return err
	}

	err = uc.RedisClient.RemoveFromRedis("code-"+code)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functionCaller.PrintFuncName(),"uc-redis-removeFromRedis")
	}

	return nil
}
