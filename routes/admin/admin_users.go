package admin_users_routes

import (
	"service/handlers/admins_handlers/admin_users_handlers"
	"service/middlewares"

	"github.com/kataras/iris/v12/core/router"
)

func AuthHTTP(app router.Party) {
	users := app.Party("/admin/users", middlewares.Auth)
	users.Get("/", admin_users_handlers.ListUsers)
}
