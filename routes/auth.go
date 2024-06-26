package routes

import (
	"service/dto"
	"service/handlers/auth_handlers"
	"service/middlewares"
	"service/pkg/api"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
)

func AuthHTTP(app router.Party) {
	authApi := app.Party("/auth")
	api.PreRoute = "/auth"

	regValidator := middlewares.Validate(auth_handlers.RegisterValidator, auth_handlers.RegisterReq{})
	api.Post(authApi, "/register", []context.Handler{
		regValidator, auth_handlers.Register,
	}, &auth_handlers.RegisterReq{}, &auth_handlers.RegisterRes{}, api.Setting{
		Summary: "Users register from here",
	})

	loginWithPhoneValidator := middlewares.Validate(auth_handlers.LoginWithPhoneValidator, auth_handlers.LoginWithPhoneReq{})
	api.Post(authApi, "/login_with_phone", []context.Handler{
		loginWithPhoneValidator, auth_handlers.LoginWithPhone,
	}, &auth_handlers.LoginWithPhoneReq{}, &auth_handlers.LoginWithPhoneRes{}, api.Setting{
		Summary: "Login with phone number",
	})

	loginWithEmailValidator := middlewares.Validate(auth_handlers.LoginWithEmailValidator, auth_handlers.LoginWithEmailReq{})
	api.Post(authApi, "/login_with_email", []context.Handler{
		loginWithEmailValidator, auth_handlers.LoginWithEmail,
	}, &auth_handlers.LoginWithEmailReq{}, &auth_handlers.LoginWithEmailRes{}, api.Setting{
		Summary: "Login with email",
	})

	api.Post[any](authApi, "/refresh_token", []context.Handler{
		auth_handlers.RefreshToken,
	}, nil, &auth_handlers.RefreshTokenRes{}, api.Setting{
		Summary:       "Refreshs token when refresh token is in the header instead of access token",
		Authorization: true,
	})

	api.Post[any, any](authApi, "/logout", []context.Handler{auth_handlers.Logout}, nil, nil, api.Setting{Summary: "Logs out"})

	api.Post[any](authApi, "/request_verify_email", []context.Handler{
		middlewares.Auth, auth_handlers.RequestVerifyEmail,
	}, nil, &dto.Message{}, api.Setting{
		Summary:       "Sends a request to send a key to the registered email to verify it later with another api",
		Authorization: true,
	})

	api.Post[any](authApi, "/request_verify_phone_number", []context.Handler{
		middlewares.Auth, auth_handlers.RequestVerifyPhoneNumber,
	}, nil, &dto.Message{}, api.Setting{
		Summary:       "Sends a request to send a key to the registered phone number to verify it later with another api",
		Authorization: true,
	})

	verifyEmailValidator := middlewares.Validate(auth_handlers.VerifyEmailValidator, auth_handlers.VerifyEmailReq{})
	api.Post(authApi, "/verify_email", []context.Handler{
		middlewares.Auth, verifyEmailValidator, auth_handlers.VerifyEmail,
	}, &auth_handlers.VerifyEmailReq{}, &dto.Message{}, api.Setting{
		Summary:       "Verifies sent key that has been sent to the registered email",
		Authorization: true,
	})

	verifyPhoneNumberValidator := middlewares.Validate(auth_handlers.VerifyPhoneNumberValidator, auth_handlers.VerifyPhoneNumberReq{})
	api.Post[any, any](authApi, "/verify_phone_number", []context.Handler{
		middlewares.Auth, verifyPhoneNumberValidator, auth_handlers.VerifyPhoneNumber,
	}, nil, nil, api.Setting{
		Summary:       "Verifies sent key that has been sent to the registered phone number",
		Authorization: true,
	})
}
