package auth_handlers

import (
	"database/sql"
	g "service/global"
	"service/pkg/copier"
	"service/repositories"
	"service/utils"

	"github.com/kataras/iris/v12"
)

type RegisterReq struct {
	PhoneNumber string `json:"phone_number" g:"required,max=16,phone"`
	Email       string `json:"email" g:"required,max=64"`
	Password    string `json:"password" g:"required,min=3"`
}

var RegisterValidator = g.Galidator.Validator(RegisterReq{})

func Register(ctx iris.Context) {
	req := ctx.Values().Get(g.RequestBody).(*RegisterReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)

	user := &repositories.User{}
	copier.Copy(user, req)
	user.HashPassword()

	response := &repositories.RegisterUserParams{}
	repositories.New(db).RegisterUser(ctx, copier.Copy(response, user).(repositories.RegisterUserParams))

	utils.SendJson(ctx, user)
}
