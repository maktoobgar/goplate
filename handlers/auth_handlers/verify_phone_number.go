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

type VerifyPhoneNumberReq struct {
	Code string `json:"code" g:"required,min=6"`
}

var VerifyPhoneNumberValidator = validators.Generator.Validator(VerifyPhoneNumberReq{})

func VerifyPhoneNumber(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := *ctx.Value(g.UserKey).(*repositories.User)
	req := ctx.Values().Get(g.RequestBody).(*VerifyPhoneNumberReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _, _, _ = db, req, user, translator, queries

	if user.PhoneNumberVerified {
		panic(errors.New(errors.InvalidStatus, translator.Auth().PhoneNumberIsAlreadyVerified(), "phone number is already verified"))
	}

	params := user.GetParams()
	if _, ok := params["phoneExpire"]; !ok {
		panic(errors.New(errors.InvalidStatus, translator.Auth().FirstRequestForVerifyCode(), "no code generated before"))
	}

	if attempts := params["phoneAttempts"]; int(attempts.(float64)) <= 0 {
		panic(errors.New(errors.InvalidStatus, translator.Auth().PhoneNumberVerifyCodeTooManyRequests(), "too many requests"))
	}

	phoneAttempts := int(params["phoneAttempts"].(float64))
	phoneAttempts -= 1
	params["phoneAttempts"] = phoneAttempts
	user.SetParams(ctx, db, params)
	if code := params["phoneCode"]; code != req.Code {
		if phoneAttempts == 0 {
			panic(errors.New(errors.InvalidStatus, translator.Auth().PhoneNumberVerifyCodeTooManyRequests(), "too many requests"))
		}
		panic(errors.New(errors.InvalidStatus, translator.Auth().WrongCode(phoneAttempts), "wrong code provided"))
	}

	now := time.Now()
	expireFloat := params["phoneExpire"]
	expire := int64(expireFloat.(float64))
	if expire < now.Unix() {
		panic(errors.New(errors.InvalidStatus, translator.Auth().PhoneNumberVerifyCodeExpired(), "code expired"))
	}

	queries.ConfirmPhoneNumber(ctx, user.ID)
	delete(params, "phoneCode")
	delete(params, "phoneAttempts")
	delete(params, "phoneExpire")
	delete(params, "phoneRefresh")
	user.SetParams(ctx, db, params)
	utils.SendMessage(ctx, translator.Auth().PhoneNumberVerified())
}
