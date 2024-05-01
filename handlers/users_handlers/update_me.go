package users_handlers

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/repositories"

	"github.com/kataras/iris/v12"
)

type UpdateMeReq struct {
}

var UpdateMeValidator = g.Galidator.Validator(UpdateMeReq{})

func UpdateMe(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := ctx.Value(g.UserKey).(*repositories.User)
	req := ctx.Values().Get(g.RequestBody).(*UpdateMeReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	_, _, _, _ = db, req, user, translator
}
