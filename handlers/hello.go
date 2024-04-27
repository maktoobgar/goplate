package handlers

import (
	g "service/global"
	"service/utils"

	i18nInterfaces "service/i18n/interfaces"

	"github.com/kataras/iris/v12"
)

func Hello(ctx iris.Context) {
	translate := ctx.Values().Get(g.TranslateKey).(i18nInterfaces.TranslatorI)
	utils.SendJson(ctx, map[string]string{
		"message": translate.HelloWorld() + " 🥳",
	})
}
