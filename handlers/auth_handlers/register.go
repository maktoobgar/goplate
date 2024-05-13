package auth_handlers

import (
	"database/sql"
	g "service/global"
	"service/pkg/copier"
	"service/repositories"
	"service/utils"
	"service/validators"

	"github.com/kataras/iris/v12"
)

type RegisterReq struct {
	PhoneNumber string `json:"phone_number" g:"required,max=16,phone"`
	Email       string `json:"email" g:"required,max=64,email_is_unique"`
	DisplayName string `json:"display_name" g:"required"`
	Password    string `json:"password" g:"required,min=3"`
}

type RegisterRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var RegisterValidator = validators.Generator.Validator(RegisterReq{})

func Register(ctx iris.Context) {
	req := ctx.Values().Get(g.RequestBody).(*RegisterReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)

	user := repositories.NewUser(req)

	user, _ = repositories.New(db).RegisterUser(ctx, copier.Copy(&repositories.RegisterUserParams{}, &user))

	accessTokenObject, accessToken := user.GenerateAccessToken(ctx, db)

	_, refreshToken := user.GenerateRefreshToken(ctx, db, accessTokenObject.ID)

	utils.SendJson(ctx, RegisterRes{AccessToken: accessToken, RefreshToken: refreshToken})
}
