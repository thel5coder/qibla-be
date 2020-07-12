package bootstrap

import (
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"qibla-backend/helpers/jwe"
	"qibla-backend/helpers/jwt"
	"qibla-backend/usecase"
)

type Bootstrap struct {
	E               *echo.Echo
	Db              *sql.DB
	UseCaseContract usecase.UcContract
	Jwe             jwe.Credential
	Validator       *validator.Validate
	Translator      ut.Translator
	JwtConfig       middleware.JWTConfig
	JwtCred         jwt.JwtCredential
}
