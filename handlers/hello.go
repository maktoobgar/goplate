package handlers

import (
	"service/utils"

	"github.com/kataras/iris/v12"
)

func Hello(ctx iris.Context) {
	utils.SendJson(ctx, map[string]string{
		"message": "Hello World 🥳",
	})
}
