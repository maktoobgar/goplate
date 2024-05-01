package users_handlers

import (
	"database/sql"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/copier"
	"service/repositories"
	"service/utils"

	"github.com/golodash/galidator"
	"github.com/kataras/iris/v12"
)

type UpdateMeReq struct {
	FirstName   string `json:"first_name" g:"required"`
	LastName    string `json:"last_name" g:"required"`
	DisplayName string `json:"display_name" g:"required"`
	// 0 notdefined 1 male 2 female
	Gender int32 `json:"gender" g:"choices=0,1,2"`
}

var UpdateMeValidator = g.Galidator.ComplexValidator(galidator.Rules{
	"FirstName":   g.Galidator.R("first_name").Required(),
	"LastName":    g.Galidator.R("last_name").Required(),
	"DisplayName": g.Galidator.R("display_name").Required(),
	"Gender":      g.Galidator.R("gender").Choices(int32(0), int32(1), int32(2)),
})

func UpdateMe(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := *ctx.Value(g.UserKey).(*repositories.User)
	req := ctx.Values().Get(g.RequestBody).(*UpdateMeReq)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _, _, _ = db, req, user, translator, queries

	user, _ = queries.UpdateMe(ctx, copier.Copy(&repositories.UpdateMeParams{ID: user.ID}, req))
	utils.SendJson(ctx, copier.Copy(&MeRes{}, &user))
}
