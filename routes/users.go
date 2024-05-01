package routes

import (
	"service/handlers/users_handlers"
	"service/middlewares"

	"github.com/kataras/iris/v12/core/router"
)

func UsersHTTP(app router.Party) {
	api := app.Party("/", middlewares.Auth)
	// usersApi := api.Party("/api")

	api.Get("/me", users_handlers.Me)

	// updateMeValidator := middlewares.Validate(auth_handlers.UpdateMe, auth_handlers.UpdateMeReq{})
	// api.Put("/me", updateMeValidator, users_handlers.UpdateMe)
}
