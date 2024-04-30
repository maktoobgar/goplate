package auth_handlers

import (
	"database/sql"
	g "service/global"
	"service/repositories"
	"service/utils"

	"github.com/kataras/iris/v12"
)

type LoginReq struct {
	Username string `json:"username" g:"required,phone"`
	Password string `json:"password" g:"required"`
}

func LoginWithPhone(ctx iris.Context) {
	req := ctx.Values().Get(g.RequestBody).(*LoginReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)

	user := repositories.NewUser(req)
	user.PhoneNumber = req.Username
	repositories.New(db)

	utils.SendJson(ctx, nil)
}
