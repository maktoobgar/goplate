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

	updateMeValidator := middlewares.Validate(users_handlers.UpdateMeValidator, users_handlers.UpdateMeReq{})
	api.Put("/me", updateMeValidator, users_handlers.UpdateMe)

	updateMePatchValidator := middlewares.Validate(users_handlers.UpdateMePatchValidator, users_handlers.UpdateMePatchReq{})
	api.Patch("/me", updateMePatchValidator, users_handlers.UpdateMePatch)
}
