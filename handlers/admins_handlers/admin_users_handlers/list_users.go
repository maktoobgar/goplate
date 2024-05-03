package admin_users_handlers

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/repositories"
	"service/utils"

	"github.com/kataras/iris/v12"
)

func ListUsers(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := *ctx.Value(g.UserKey).(*repositories.User)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _, _ = db, user, translator, queries

	users, _ := queries.ListUsers(ctx)
	utils.SendPage(ctx, len(users), len(users), 1, users)
}
