package handlers

import (
	g "service/global"
	"service/utils"

	"service/i18n/i18n_interfaces"

	"github.com/kataras/iris/v12"
)

func Hello(ctx iris.Context) {
	translator := ctx.Values().Get(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	utils.SendMessage(ctx, translator.HelloWorld()+" 🥳")
}
