package extra_middlewares

import (
	g "service/global"
	"service/i18n"

	i18nInterfaces "service/i18n/interfaces"

	"github.com/kataras/iris/v12"
)

func Translator(ctx iris.Context) {
	lang := ctx.GetHeader("Accept-Language")
	if lang == "" {
		lang = ctx.GetCookie("Accept-Language")
	}

	var translateFunc i18nInterfaces.TranslatorI = i18n.NewTranslator(lang)
	ctx.Values().Set(g.TranslateKey, translateFunc)
	ctx.Next()
}
