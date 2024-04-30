package error_handlers

import (
	g "service/global"
	"service/pkg/errors"

	"service/i18n/i18n_interfaces"

	"github.com/kataras/iris/v12"
)

func NotFound(ctx iris.Context) {
	translator := ctx.Values().Get(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	panic(errors.New(errors.NotFoundStatus, translator.StatusCodes().PageNotFound(), "page not found"))
}
