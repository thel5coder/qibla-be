package middleware

import (
	"errors"
	"fmt"
	jwt2 "qibla-backend/pkg/jwt"
	"qibla-backend/pkg/messages"
	"qibla-backend/server/handlers"
	"qibla-backend/usecase"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtVerify struct {
	*usecase.UcContract
}

func (jwtVerify JwtVerify) JWTWithConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (err error) {
		claims := &jwt2.CustomClaims{}
		apiHandler := handlers.Handler{UseCaseContract: jwtVerify.UcContract}

		tokenAuthHeader := ctx.Request().Header.Get("Authorization")
		if !strings.Contains(tokenAuthHeader, "Bearer") {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.AuthHeaderNotPresent))
		}

		tokenAuth := strings.Replace(tokenAuthHeader, "Bearer ", "", -1)
		_, err = jwt.ParseWithClaims(tokenAuth, claims, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			secret := jwtVerify.JwtConfig.SigningKey
			return secret, nil
		})
		if err != nil {
			return apiHandler.SendResponseUnauthorized(ctx, err)
		}

		if claims.ExpiresAt < time.Now().Unix() {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.ExpiredToken))
		}

		jweRes, err := jwtVerify.Jwe.Rollback(claims.Id)
		if err != nil {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.FailedLoadPayload))
		}
		if jweRes == nil {
			return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.FailedLoadPayload))
		}
		claims.Id = fmt.Sprintf("%v", jweRes["id"])
		jwtVerify.UcContract.UserID = claims.Id

		//sessionData := viewmodel.UserSessionVm{}
		//jwtVerify.UcContract.RedisClient.GetFromRedis("session-"+claims.Id, &sessionData)
		//if sessionData.Session != claims.Session {
		//	return apiHandler.SendResponseUnauthorized(ctx, errors.New(messages.InvalidSession))
		//}

		//set context
		ctx.Set("user", claims)

		return next(ctx)
	}
}
