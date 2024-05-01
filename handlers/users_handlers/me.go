package users_handlers

import (
	g "service/global"
	"service/repositories"
	"service/utils"

	"github.com/kataras/iris/v12"
)

func Me(ctx iris.Context) {
	user := ctx.Value(g.UserKey).(*repositories.User)
	utils.SendJson(ctx, user)
}
