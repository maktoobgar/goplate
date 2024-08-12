package admin_users_handlers

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

type ListUsersParams struct {
	OrderBy             string `json:"order_by" g:"choices=id&first_name&last_name&display_name&phone_number&email&gender&is_active&is_admin&created_at"`
	Search              string `json:"search"`
	Desc                bool   `json:"desc"`
	PerPage             int32  `json:"per_page" g:"min=1,max=50"`
	Page                int32  `json:"page" g:"min=1"`
	PhoneNumberVerified *bool  `json:"phone_number_verified"`
	EmailVerified       *bool  `json:"email_verified"`
	Gender              *int32 `json:"gender"`
	IsActive            *bool  `json:"is_active"`
	IsAdmin             *bool  `json:"is_admin"`
}

var ListUsersParamsValidator = validators.Generator.Validator(ListUsersParams{})

var DefaultListUsersParams = ListUsersParams{
	OrderBy:             "id",
	Search:              "",
	Desc:                false,
	PerPage:             10,
	Page:                1,
	PhoneNumberVerified: nil,
	EmailVerified:       nil,
	Gender:              nil,
	IsActive:            nil,
	IsAdmin:             nil,
}

func ListUsers(ctx iris.Context) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	user := *ctx.Value(g.UserKey).(*repositories.User)
	params := ctx.Value(g.Params).(*ListUsersParams)
	db := ctx.Values().Get(g.DbInstance).(*sql.DB)
	queries := repositories.New(db)
	_, _, _, _ = db, user, translator, queries

	usersRes, _ := queries.ListUsers(ctx, copier.Copy(&repositories.ListUsersParams{}, params))
	utils.SendPage(ctx, utils.GetTotalCount(usersRes), params.PerPage, params.Page, copier.ArrayCopy(make([]repositories.User, len(usersRes)), usersRes))
}
