package middlewares

import (
	"reflect"
	g "service/global"
	"service/pkg/copier"
	"service/utils"

	"github.com/golodash/galidator"
	"github.com/kataras/iris/v12"
)

func ParamsValidator[T1, T2 any](params T1, defaultParams T2, validator galidator.Validator) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		params := reflect.New(reflect.ValueOf(params).Type()).Elem().Interface().(T1)
		copier.Copy(&params, &defaultParams)
		copier.CastParams(ctx, &params)
		utils.Validate(ctx, params, validator)

		ctx.Values().Set(g.Params, &params)
		ctx.Next()
	}
}
