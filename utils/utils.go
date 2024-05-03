package utils

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	g "service/global"
	"service/i18n/i18n_interfaces"
	"service/pkg/errors"
	"strings"
	"sync"

	"github.com/kataras/iris/v12"
)

func sendIfCtxNotCancelled[T any](ctx iris.Context, status int, value *T, sendEmpty ...bool) {
	if ctx.Err() == nil {
		writerLock := ctx.Values().Get(g.WriterLock).(*sync.Mutex)
		writerLock.Lock()
		defer writerLock.Unlock()

		closedWriter := ctx.Values().Get(g.ClosedWriter).(bool)
		if !closedWriter {
			if reflectValue := reflect.ValueOf(value); reflectValue.IsValid() && reflectValue.Kind() == reflect.Struct {
				if method := reflectValue.MethodByName("Reformat"); method.IsValid() {
					method.Call([]reflect.Value{reflectValue})
				}
			}
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

	sendIfCtxNotCancelled(ctx, code, &data)
}

func SendJson[T any](ctx iris.Context, data T, status ...int) {
	code := 200
	if len(status) > 0 {
		code = status[0]
	}

	sendIfCtxNotCancelled(ctx, code, &data)
}

func SendEmpty(ctx iris.Context, status ...int) {
	code := 204
	if len(status) > 0 {
		code = status[0]
	}

	sendIfCtxNotCancelled(ctx, code, &struct{}{}, true)
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
	sendIfCtxNotCancelled(ctx, -1, &output)
}

func SendPage[T any](ctx iris.Context, dataCount int, perPage int, page int, data []T) {
	translator := ctx.Values().Get(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	pagesCount := CalculatePagesCount(dataCount, perPage)
	if page > pagesCount {
		panic(errors.New(errors.NotFoundStatus, translator.StatusCodes().PageNotFound(), fmt.Sprintf("page %d requested but we have %d pages", page, pagesCount)))
	}
	dataLen := len(data)

	for _, singleData := range data {
		singleDataPoint := &singleData
		if reflectValue := reflect.ValueOf(singleDataPoint); reflectValue.IsValid() && reflectValue.Elem().Kind() == reflect.Struct {
			if method := reflectValue.MethodByName("Reformat"); method.IsValid() {
				method.Call([]reflect.Value{})
			}
		}
	}

	sendIfCtxNotCancelled(ctx, -1, &map[string]any{
		"page":        page,
		"per_page":    perPage,
		"pages_count": pagesCount,
		"all_count":   dataCount,
		"count":       dataLen,
		"data":        data,
	})
}

func CalculatePagesCount(dataCount int, perPage int) int {
	pagesCount := -1
	if dataCount%int(perPage) == 0 {
		pagesCount = dataCount / perPage
	} else {
		pagesCount = (dataCount / perPage) + 1
	}

	// If there is no data, just return 1 page so that NotFound do not get returned
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

// Returns the content, then content type, then the extension
func GetFile(content string) ([]byte, string, string) {
	fileByte, _ := base64.StdEncoding.DecodeString(content)
	contentType := http.DetectContentType(fileByte)
	extension := ""
	extensionExtractor := strings.Split(contentType, "/")
	if len(extensionExtractor) > 1 {
		extension = extensionExtractor[1]
	}
	return fileByte, contentType, extension
}

func EncodeId(id int32) string {
	// Encode the 64 bit number
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(id))
	encoded := base64.StdEncoding.EncodeToString([]byte(b))

	// https://youtu.be/gocwRvLhDf8?t=75
	encodedId := strings.ReplaceAll(encoded[:11], "+", "-")
	encodedId = strings.ReplaceAll(encodedId, "/", "_")
	return encodedId
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateVerificationCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
