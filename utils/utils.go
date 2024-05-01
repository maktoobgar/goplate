package utils

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/errors"
	"sync"

	"github.com/golodash/galidator"
	"github.com/kataras/iris/v12"
)

func sendIfCtxNotCancelled(ctx iris.Context, status int, value any, sendEmpty ...bool) {
	if ctx.Err() == nil {
		writerLock := ctx.Values().Get(g.WriterLock).(*sync.Mutex)
		writerLock.Lock()
		defer writerLock.Unlock()

		closedWriter := ctx.Values().Get(g.ClosedWriter).(bool)
		if !closedWriter {
			if status != -1 {
				ctx.StatusCode(status)
			}
			if len(sendEmpty) == 0 || !sendEmpty[0] {
				ctx.JSON(value)
			}
		}

		closedWriter = true
		ctx.Values().Set(g.ClosedWriter, closedWriter)
	}
}

func SendJsonMessage(ctx iris.Context, message string, data map[string]any, status ...int) {
	code := 200
	if len(status) > 0 {
		code = status[0]
	}
	data["message"] = message

	sendIfCtxNotCancelled(ctx, code, data)
}

func SendJson(ctx iris.Context, data any, status ...int) {
	code := 200
	if len(status) > 0 {
		code = status[0]
	}

	sendIfCtxNotCancelled(ctx, code, data)
}

func SendEmpty(ctx iris.Context, status ...int) {
	code := 200
	if len(status) > 0 {
		code = status[0]
	}

	sendIfCtxNotCancelled(ctx, code, nil, true)
}

func Panic500(err error) {
	panic(errors.New(errors.UnexpectedStatus, "InternalServerError", err.Error(), nil))
}

func SendMessage(ctx iris.Context, message string, data ...map[string]any) {
	var output = map[string]any{}
	if len(data) != 0 {
		output = data[0]
	}

	output["message"] = message
	sendIfCtxNotCancelled(ctx, -1, output)
}

func SendPage(ctx iris.Context, dataCount int64, perPage int, page int, data any) {
	translator := ctx.Values().Get(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	pagesCount := CalculatePagesCount(dataCount, perPage)
	if page > pagesCount {
		panic(errors.New(errors.NotFoundStatus, translator.StatusCodes().PageNotFound(), fmt.Sprintf("page %d requested but we have %d pages", page, pagesCount)))
	}
	dataValue := reflect.ValueOf(data)
	if dataValue.Type().Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}

	sendIfCtxNotCancelled(ctx, -1, map[string]any{
		"page":        page,
		"per_page":    perPage,
		"pages_count": pagesCount,
		"all_count":   dataCount,
		"count":       dataValue.Len(),
		"data":        data,
	})
}

func CalculatePagesCount(dataCount int64, perPage int) int {
	pagesCount := int64(-1)
	if dataCount%int64(perPage) == 0 {
		pagesCount = dataCount / int64(perPage)
	} else {
		pagesCount = (dataCount / int64(perPage)) + 1
	}

	// If there is no date, just return 1 page so that NotFound do not get returned
	if int(pagesCount) == 0 {
		return 1
	}
	return int(pagesCount)
}

func Min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	} else if v2 < v1 {
		return v2
	} else {
		return v1
	}
}

func Validate(data any, validator galidator.Validator, translator func(key string, optionalInputs ...[]any) string) {
	if errs := validator.Validate(data, func(s string) string { return translator(s) }); errs != nil {
		panic(errors.New(errors.InvalidStatus, "BodyNotProvidedProperly", "", errs))
	}
}

func PrettyJsonBytes(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		return ""
	}
	return prettyJSON.String()
}

func CastParams(params interface{}, ctx iris.Context, defaultValues ...interface{}) {
	paramsValue := reflect.ValueOf(params)
	paramsType := reflect.TypeOf(params)
	for paramsType.Kind() == reflect.Ptr {
		paramsValue = paramsValue.Elem()
		paramsType = paramsValue.Type()
	}

	var defaults interface{} = nil
	var defaultsValue reflect.Value
	if len(defaultValues) > 0 {
		defaults = defaultValues[0]
		defaultsValue = reflect.ValueOf(defaults)
	}

	for i := 0; i < paramsValue.NumField(); i++ {
		elemType := paramsType.Field(i)
		elemValue := paramsValue.Field(i)
		isPointer := false
		exactElemType := elemType.Type
		if exactElemType.Kind() == reflect.Ptr {
			isPointer = true
			exactElemType = elemType.Type.Elem()
		}
		if elemType.IsExported() {
			var tag = elemType.Tag.Get("json")
			var value interface{} = nil
			switch exactElemType.Kind() {
			case reflect.Bool:
				if defaults != nil {
					value = ctx.URLParamBoolDefault(tag, defaultsValue.FieldByName(elemType.Name).Bool())
					break
				}
				if urlParamDoesNotExist := ctx.URLParamExists(tag); urlParamDoesNotExist {
					value, _ = ctx.URLParamBool(tag)
				}
			case reflect.Float64:
				if defaults != nil {
					value = ctx.URLParamFloat64Default(tag, defaultsValue.FieldByName(elemType.Name).Float())
					break
				}
				if urlParamDoesNotExist := ctx.URLParamExists(tag); urlParamDoesNotExist {
					value, _ = ctx.URLParamFloat64(tag)
				}
			case reflect.Int:
				if defaults != nil {
					value = ctx.URLParamIntDefault(tag, int(defaultsValue.FieldByName(elemType.Name).Int()))
					break
				}
				if urlParamDoesNotExist := ctx.URLParamExists(tag); urlParamDoesNotExist {
					value, _ = ctx.URLParamInt(tag)
				}
			case reflect.Int64:
				if defaults != nil {
					value = ctx.URLParamInt64Default(tag, defaultsValue.FieldByName(elemType.Name).Int())
					break
				}
				if urlParamDoesNotExist := ctx.URLParamExists(tag); urlParamDoesNotExist {
					value, _ = ctx.URLParamInt64(tag)
				}
			case reflect.String:
				if defaults != nil {
					value = ctx.URLParamDefault(tag, defaultsValue.FieldByName(elemType.Name).String())
					break
				}
				if urlParamDoesNotExist := ctx.URLParamExists(tag); urlParamDoesNotExist {
					value = ctx.URLParam(tag)
				}
			}
			valueValue := reflect.ValueOf(value)
			if valueValue.IsValid() {
				if isPointer {
					pointerToValue := reflect.New(exactElemType)
					pointerToValue.Elem().Set(reflect.ValueOf(value))
					elemValue.Set(pointerToValue)
				} else {
					elemValue.Set(valueValue)
				}
			}
		}
	}
}

func IsErrorNotFound(err error) bool {
	return err != nil && err == sql.ErrNoRows
}
