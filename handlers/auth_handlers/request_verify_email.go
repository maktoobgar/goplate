package auth_handlers

import (
	"database/sql"
	"fmt"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/errors"
	"service/repositories"
	"service/utils"
	"time"

	"github.com/kataras/iris/v12"
)

func RequestVerifyEmail(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := *ctx.Value(g.UserKey).(*repositories.User)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _, _ = db, user, translator, queries

	if !user.Email.Valid || user.Email.String == "" {
		panic(errors.New(errors.InvalidStatus, translator.Auth().EmailNotFound(), "user has no email"))
	}

	if user.EmailVerified {
		panic(errors.New(errors.InvalidStatus, translator.Auth().EmailIsAlreadyVerified(), "email is already verified"))
	}

	params := user.GetParams()
	// Check if it is allowed to generate new code
	now := time.Now()
	if refreshTimeFloat, ok := params["emailRefresh"]; ok {
		refreshTime := int64(refreshTimeFloat.(float64))
		if refreshTime > now.Unix() {
			remainingTime := refreshTime - now.Unix()
			panic(errors.New(errors.TooManyRequests, translator.Auth().EmailVerifyCodeCoolDown(remainingTime), fmt.Sprintf("try %d seconds later", remainingTime)))
		}
	}

	generatedCode := utils.GenerateVerificationCode(6)
	params["emailCode"] = generatedCode
	params["emailAttempts"] = 10
	params["emailExpire"] = now.Add(time.Minute * 10).Unix()
	params["emailRefresh"] = now.Add(time.Minute * 3).Unix()
	user, _ = user.SetParams(ctx, db, params)

	if g.CFG.Debug {
		fmt.Println("emailCode: ", generatedCode)
	}

	utils.SendMessage(ctx, translator.Auth().EmailVerifyCodeSent())
}
