package extra_middlewares

import (
	g "service/global"
	"service/i18n"

	"service/i18n/i18n_interfaces"

	"github.com/kataras/iris/v12"
)

func Translator(ctx iris.Context) {
	lang := ctx.GetHeader("Accept-Language")
	if lang == "" {
		lang = ctx.GetCookie("Accept-Language")
	}

	var translateFunc i18n_interfaces.TranslatorI = i18n.NewTranslator(lang)
	ctx.Values().Set(g.TranslatorKey, translateFunc)
	ctx.Next()
}
