package admin_users_routes

import (
	"service/handlers/admins_handlers/admin_users_handlers"
	"service/middlewares"

	"github.com/kataras/iris/v12/core/router"
)

func AdminHTTP(app router.Party) {
	users := app.Party("/admin/users", middlewares.Auth)

	adminListUsers := middlewares.ParamsValidator(admin_users_handlers.ListUsersParams{}, admin_users_handlers.DefaultListUsersParams, admin_users_handlers.ListUsersParamsValidator)
	users.Get("/", adminListUsers, admin_users_handlers.ListUsers)
}
