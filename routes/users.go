package routes

import (
	"service/handlers/users_handlers"
	"service/middlewares"
	"service/pkg/api"

	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
)

func UsersHTTP(app router.Party) {
	users := app.Party("/", middlewares.Auth)
	api.PreRoute = ""

	api.Get(users, "/me", []context.Handler{users_handlers.Me}, &users_handlers.MeRes{}, api.Setting{
		Summary:       "Returns user's data",
		Authorization: true,
	})

	updateMeValidator := middlewares.Validate(users_handlers.UpdateMeValidator, users_handlers.UpdateMeReq{})
	api.Put(users, "/me", []context.Handler{
		updateMeValidator, users_handlers.UpdateMe,
	}, &users_handlers.UpdateMeReq{}, &users_handlers.MeRes{}, api.Setting{
		Summary:       "Updates user's data and returns updated version in response",
		Authorization: true,
	})

	updateMePatchValidator := middlewares.Validate(users_handlers.UpdateMePatchValidator, users_handlers.UpdateMePatchReq{})
	api.Patch(users, "/me", []context.Handler{
		updateMePatchValidator, users_handlers.UpdateMePatch,
	}, &users_handlers.UpdateMePatchReq{}, &users_handlers.MeRes{}, api.Setting{
		Summary:       "Updates user's data and returns updated version in response",
		Authorization: true,
	})

	updateAvatarValidator := middlewares.Validate(users_handlers.UpdateAvatarValidator, users_handlers.UpdateAvatarReq{})
	api.Put(users, "/me/avatar", []context.Handler{
		updateAvatarValidator, users_handlers.UpdateAvatar,
	}, &users_handlers.UpdateAvatarReq{}, &users_handlers.MeRes{}, api.Setting{
		Summary:       "Updates user's avatar and returns user's data",
		Authorization: true,
	})
}
