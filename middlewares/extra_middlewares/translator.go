package extra_middlewares

import (
	g "service/global"

	"github.com/kataras/iris/v12"
)

func Translator(ctx iris.Context) {
	lang := ctx.GetHeader("Accept-Language")
	if lang == "" {
		lang = ctx.GetCookie("Accept-Language")
	}

	ctx.Values().Set(g.TranslateKey, g.Translator.TranslateFunction(lang, "en"))
	ctx.Next()
}
