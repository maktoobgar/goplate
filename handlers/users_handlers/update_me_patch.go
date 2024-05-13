package users_handlers

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/copier"
	"service/repositories"
	"service/utils"
	"service/validators"

	"github.com/golodash/galidator/v2"
	"github.com/kataras/iris/v12"
)

type UpdateMePatchReq struct {
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	DisplayName *string `json:"display_name"`
	// 0 notdefined 1 male 2 female
	Gender *int32 `json:"gender" g:"choices=0,1,2"`
}

var UpdateMePatchValidator = validators.Generator.ComplexValidator(galidator.Rules{
	"Gender": validators.Generator.R("gender").Choices(int32(0), int32(1), int32(2)),
})

func UpdateMePatch(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := *ctx.Value(g.UserKey).(*repositories.User)
	req := ctx.Values().Get(g.RequestBody).(*UpdateMePatchReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _, _, _ = db, req, user, translator, queries

	updateParams := copier.Copy(&repositories.UpdateMeParams{ID: user.ID}, &user)
	user, _ = queries.UpdateMe(ctx, copier.Copy(&updateParams, req))
	utils.SendJson(ctx, copier.Copy(&MeRes{}, &user))
}
