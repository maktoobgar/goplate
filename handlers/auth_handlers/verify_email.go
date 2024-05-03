package auth_handlers

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/errors"
	"service/repositories"
	"service/utils"
	"service/validators"
	"time"

	"github.com/kataras/iris/v12"
)

type VerifyEmailReq struct {
	Code string `json:"code" g:"required,min=6"`
}

var VerifyEmailValidator = validators.Generator.Validator(VerifyEmailReq{})

func VerifyEmail(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := *ctx.Value(g.UserKey).(*repositories.User)
	req := ctx.Values().Get(g.RequestBody).(*VerifyEmailReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _, _, _ = db, req, user, translator, queries

	if user.EmailVerified {
		panic(errors.New(errors.InvalidStatus, translator.Auth().EmailIsAlreadyVerified(), "email is already verified"))
	}

	params := user.GetParams()
	if _, ok := params["emailExpire"]; !ok {
		panic(errors.New(errors.InvalidStatus, translator.Auth().FirstRequestForVerifyCode(), "no code generated before"))
	}

	if attempts := params["emailAttempts"]; int(attempts.(float64)) <= 0 {
		panic(errors.New(errors.InvalidStatus, translator.Auth().EmailVerifyCodeTooManyRequests(), "too many requests"))
	}

	emailAttempts := int(params["emailAttempts"].(float64))
	emailAttempts -= 1
	params["emailAttempts"] = emailAttempts
	user.SetParams(ctx, db, params)
	if code := params["emailCode"]; code != req.Code {
		if emailAttempts == 0 {
			panic(errors.New(errors.InvalidStatus, translator.Auth().EmailVerifyCodeTooManyRequests(), "too many requests"))
		}
		panic(errors.New(errors.InvalidStatus, translator.Auth().WrongCode(emailAttempts), "wrong code provided"))
	}

	now := time.Now()
	expireFloat := params["emailExpire"]
	expire := int64(expireFloat.(float64))
	if expire < now.Unix() {
		panic(errors.New(errors.InvalidStatus, translator.Auth().EmailVerifyCodeExpired(), "code expired"))
	}

	queries.ConfirmEmail(ctx, user.ID)
	delete(params, "emailCode")
	delete(params, "emailAttempts")
	delete(params, "emailExpire")
	delete(params, "emailRefresh")
	user.SetParams(ctx, db, params)
	utils.SendMessage(ctx, translator.Auth().EmailVerified())
}
