package auth_handlers

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/copier"
	"service/pkg/errors"
	"service/repositories"
	"service/utils"

	"github.com/kataras/iris/v12"
)

type LoginWithEmailReq struct {
	Username string `json:"username" g:"required,email"`
	Password string `json:"password" g:"required"`
}

type LoginWithEmailRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var LoginWithEmailValidator = g.Galidator.Validator(LoginWithEmailReq{})

func LoginWithEmail(ctx iris.Context) {
	translator := ctx.Values().Get(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	req := ctx.Values().Get(g.RequestBody).(*LoginWithEmailReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)

	user, err := repositories.New(db).LoginUserWithEmail(ctx, sql.NullString{String: req.Username, Valid: true})
	if err != nil && err == sql.ErrNoRows {
		panic(errors.New(errors.InvalidStatus, translator.Auth().UserWithEmailNotFound(), err.Error()))
	}

	if !user.IsSamePassword(req.Password) {
		panic(errors.New(errors.InvalidStatus, translator.Auth().WrongPasswordWithEmailPassword(), "Password doesn't match with account"))
	}

	accessToken := user.GenerateToken()
	refreshToken := user.GenerateToken(true)
	accessToken, _ = repositories.New(db).CreateAccessToken(ctx, copier.Copy(&repositories.CreateAccessTokenParams{}, &accessToken))
	refreshToken, _ = repositories.New(db).CreateAccessToken(ctx, copier.Copy(&repositories.CreateAccessTokenParams{}, &refreshToken))

	utils.SendJson(ctx, LoginWithPhoneRes{AccessToken: accessToken.Token, RefreshToken: refreshToken.Token})
}
