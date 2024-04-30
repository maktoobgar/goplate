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

	loginWithPhoneValidator := middlewares.Validate(auth_handlers.LoginWithPhoneValidator, auth_handlers.LoginWithPhoneReq{})
	api.Post("/login_with_phone", loginWithPhoneValidator, auth_handlers.LoginWithPhone)

	loginWithEmailValidator := middlewares.Validate(auth_handlers.LoginWithEmailValidator, auth_handlers.LoginWithEmailReq{})
	api.Post("/login_with_email", loginWithEmailValidator, auth_handlers.LoginWithEmail)
}
