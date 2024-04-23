package extra_middlewares

import (
	g "service/global"
	"service/pkg/translator"

	"github.com/kataras/iris/v12"
)

func Translator(ctx iris.Context) {
	lang := ctx.GetHeader("Accept-Language")
	if lang == "" {
		lang = ctx.GetCookie("Accept-Language")
	}

	// TODO: Fix
	var translateFunc translator.TranslatorFunc = func(value string) string { return value }
	ctx.Values().Set(g.TranslateKey, translateFunc)
	ctx.Next()
}
