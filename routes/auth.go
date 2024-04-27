package routes

import (
	"service/handlers/auth_handlers"
	"service/middlewares"

	"github.com/kataras/iris/v12/core/router"
)

func AuthHTTP(app router.Party) {
	api := app.Party("/auth")
	regValidator := middlewares.Validate(auth_handlers.RegisterValidator, auth_handlers.RegisterReq{})
	api.Post("/register", regValidator, auth_handlers.Register)
}
