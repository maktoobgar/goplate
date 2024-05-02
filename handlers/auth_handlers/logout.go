package auth_handlers

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/errors"
	"service/repositories"
	"service/static_models"
	"service/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

func Logout(ctx iris.Context) {
	refreshToken := ctx.GetHeader(g.RefreshToken)
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _ = db, translator, queries

	accessTokenAndId := refreshRe.FindStringSubmatch(refreshToken)
	if len(accessTokenAndId) == 0 {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().InvalidToken(), "user didn't provide a refresh token"))
	}
	refreshToken = accessTokenAndId[1]

	claims := &static_models.Claims{}
	tokenState, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return g.SecretKeyBytes, nil
	})
	if err != nil {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().InvalidToken(), err.Error()))
	}
	if !tokenState.Valid {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().InvalidToken(), "token is invalid"))
	}
	if claims.Type != static_models.RefreshTokenType {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().InvalidToken(), "token is not refresh token"))
	}
	if claims.ExpiresAt < time.Now().Unix() {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().InvalidToken(), "token is expired"))
	}
	if _, err := queries.GetToken(ctx, claims.Id); utils.IsErrorNotFound(err) {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().InvalidToken(), "token not found in db"))
	}

	_, err = queries.GetUserById(ctx, claims.UserId)
	if err != nil {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().InvalidToken(), "user not found in db"))
	}

	queries.DeleteToken(ctx, claims.Id)
	queries.DeleteToken(ctx, claims.AccessTokenId)

	utils.SendEmpty(ctx)
}
