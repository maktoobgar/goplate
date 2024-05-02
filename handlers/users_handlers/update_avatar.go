package users_handlers

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/copier"
	"service/repositories"
	"service/utils"
	"service/validators"

	"github.com/kataras/iris/v12"
)

type UpdateAvatarReq struct {
	Avatar string `json:"avatar" g:"required,image_type"`
}

var UpdateAvatarValidator = validators.Generator.Validator(UpdateAvatarReq{})

func UpdateAvatar(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := *ctx.Value(g.UserKey).(*repositories.User)
	req := ctx.Values().Get(g.RequestBody).(*UpdateAvatarReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _, _, _ = db, req, user, translator, queries

	avatarBytes, _, avatarExtension := utils.GetFile(req.Avatar)
	encodedId := utils.EncodeId(user.ID)
	userAddress, _ := g.UsersMedia.GoTo(encodedId, true)
	fileName := "avatar." + avatarExtension
	userAddress.OverwriteFile(avatarBytes, fileName)
	user, _ = repositories.New(db).UpdateAvatar(ctx, repositories.UpdateAvatarParams{Avatar: sql.NullString{Valid: true, String: userAddress.GetHostAddress(fileName)}, ID: user.ID})

	utils.SendJson(ctx, copier.Copy(&MeRes{}, &user))
}
