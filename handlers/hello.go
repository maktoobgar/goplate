package handlers

import (
	g "service/global"
	"service/utils"

	"service/i18n/i18n_interfaces"

	"github.com/kataras/iris/v12"
)

func Hello(ctx iris.Context) {
	translate := ctx.Values().Get(g.TranslateKey).(i18n_interfaces.TranslatorI)
	utils.SendJson(ctx, map[string]string{
		"message": translate.HelloWorld() + " 🥳",
	})
}
