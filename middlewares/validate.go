package middlewares

import (
	"encoding/json"
	"io"
	"reflect"
	g "service/global"
	"service/pkg/errors"
	"service/utils"

	"service/i18n/i18n_interfaces"

	"github.com/golodash/galidator"
	"github.com/kataras/iris/v12"
)

// Parses and validates request body
func Validate(validator galidator.Validator, inputInstance any) iris.Handler {
	return func(ctx iris.Context) {
		translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
		req := reflect.New(reflect.TypeOf(inputInstance)).Interface()
		// Parse body and check for errors
		body := ctx.Request().Body
		bytes, err1 := io.ReadAll(body)
		err2 := json.Unmarshal(bytes, req)

		if err1 != nil {
			panic(errors.New(errors.InvalidStatus, translator.StatusCodes().BodyNotProvidedProperly(), err1.Error()))
		} else if err2 != nil {
			panic(errors.New(errors.InvalidStatus, translator.StatusCodes().BodyNotProvidedProperly(), err2.Error()))
		}

		// Validate and translate error messages if errors exist
		utils.Validate(ctx, req, validator)

		// If we come this far, data is valid, so record it in context
		ctx.Values().Set(g.RequestBody, req)

		ctx.Next()
	}
}
