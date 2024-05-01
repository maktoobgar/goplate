package middlewares

import (
	"database/sql"
	"regexp"
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

var authRe, _ = regexp.Compile(`^Bearer (eyJ[a-zA-Z0-9_\-]+?\.[a-zA-Z0-9_\-]+?\.[a-zA-Z0-9_\-]+?$)`)

func Auth(ctx iris.Context) {
	accessToken := ctx.GetHeader(g.AccessToken)
	translator := ctx.Values().Get(g.TranslatorKey).(*i18n_interfaces.Translator)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)

	accessTokenAndId := authRe.FindStringSubmatch(accessToken)
	if len(accessTokenAndId) == 0 {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().Unauthorized(), "user didn't provide a token"))
	}
	accessToken = accessTokenAndId[1]

	claims := &static_models.Claims{}
	tokenState, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return g.SecretKeyBytes, nil
	})
	if err != nil {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().Unauthorized(), err.Error()))
	}
	if !tokenState.Valid {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().Unauthorized(), "token is invalid"))
	}
	if claims.Type != static_models.AccessTokenType {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().Unauthorized(), "token is not access token"))
	}
	if claims.ExpiresAt < time.Now().Unix() {
		panic(errors.New(errors.UnauthorizedStatus, translator.Auth().Unauthorized(), "token is expired"))
	}

	user, err := repositories.New(db).GetUserWithTokenId(ctx, claims.Id)
	if utils.IsErrorNotFound(err) {
		panic(errors.New(errors.NotFoundStatus, translator.Users().UserNotFound(), "user not found"))
	}

	ctx.Values().Set(g.UserKey, &user)
	ctx.Next()
}
