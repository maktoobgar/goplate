package auth_handlers

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/errors"
	"service/repositories"
	"service/utils"
	"service/validators"

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

var LoginWithEmailValidator = validators.Generator.Validator(LoginWithEmailReq{})

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

	accessTokenObject, accessToken := user.GenerateAccessToken(ctx, db)

	_, refreshToken := user.GenerateRefreshToken(ctx, db, accessTokenObject.ID)

	utils.SendJson(ctx, LoginWithEmailRes{AccessToken: accessToken, RefreshToken: refreshToken})
}
