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

type LoginWithPhoneReq struct {
	Username string `json:"username" g:"required,phone"`
	Password string `json:"password" g:"required"`
}

type LoginWithPhoneRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var LoginWithPhoneValidator = validators.Generator.Validator(LoginWithPhoneReq{})

func LoginWithPhone(ctx iris.Context) {
	translator := ctx.Values().Get(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	req := ctx.Values().Get(g.RequestBody).(*LoginWithPhoneReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)

	user, err := repositories.New(db).LoginUserWithPhoneNumber(ctx, req.Username)
	if err != nil && err == sql.ErrNoRows {
		panic(errors.New(errors.InvalidStatus, translator.Auth().UserWithPhoneNumberNotFound(), err.Error()))
	}

	if !user.IsSamePassword(req.Password) {
		panic(errors.New(errors.InvalidStatus, translator.Auth().WrongPasswordWithPhoneNumberPassword(), "Password doesn't match with account"))
	}

	accessTokenObject, accessToken := user.GenerateAccessToken(ctx, db)

	_, refreshToken := user.GenerateRefreshToken(ctx, db, accessTokenObject.ID)

	utils.SendJson(ctx, LoginWithPhoneRes{AccessToken: accessToken, RefreshToken: refreshToken})
}
