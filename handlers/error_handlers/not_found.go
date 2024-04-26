package error_handlers

import (
	g "service/global"
	"service/pkg/errors"

	i18nInterfaces "service/i18n/interfaces"

	"github.com/kataras/iris/v12"
)

func NotFound(ctx iris.Context) {
	translate := ctx.Values().Get(g.TranslateKey).(i18nInterfaces.TranslatorI)
	panic(errors.New(errors.NotFoundStatus, translate.StatusCodes().PageNotFound(), "page not found"))
}
