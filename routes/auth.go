package routes

import (
	"service/handlers/auth_handlers"
	"service/middlewares"

	"github.com/kataras/iris/v12/core/router"
)

func AuthHTTP(app router.Party) {
	authApi := app.Party("/auth")
	regValidator := middlewares.Validate(auth_handlers.RegisterValidator, auth_handlers.RegisterReq{})
	authApi.Post("/register", regValidator, auth_handlers.Register)

	loginWithPhoneValidator := middlewares.Validate(auth_handlers.LoginWithPhoneValidator, auth_handlers.LoginWithPhoneReq{})
	authApi.Post("/login_with_phone", loginWithPhoneValidator, auth_handlers.LoginWithPhone)

	loginWithEmailValidator := middlewares.Validate(auth_handlers.LoginWithEmailValidator, auth_handlers.LoginWithEmailReq{})
	authApi.Post("/login_with_email", loginWithEmailValidator, auth_handlers.LoginWithEmail)

	authApi.Post("/refresh_token", auth_handlers.RefreshToken)
}
