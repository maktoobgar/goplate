package error_handlers

import (
	g "service/global"
	"service/pkg/errors"

	"service/i18n/i18n_interfaces"

	"github.com/kataras/iris/v12"
)

func NotFound(ctx iris.Context) {
	translate := ctx.Values().Get(g.TranslateKey).(i18n_interfaces.TranslatorI)
	panic(errors.New(errors.NotFoundStatus, translate.StatusCodes().PageNotFound(), "page not found"))
}
