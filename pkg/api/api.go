package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	goStrings "github.com/golodash/godash/strings"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
)

type Setting struct {
	Authorization bool
	Tag           string
	Description   string
	Summary       string
}

var (
	Generate       = false
	PreRoute       = ""
	allDefinitions = map[string]any{}
	allRoutes      = map[string]any{}

	defaultSetting = Setting{
		Authorization: false,
		Tag:           "General",
		Description:   "",
		Summary:       "",
	}
)

func getPackageName(temp interface{}) string {
	strs := strings.Split(strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")[0], "/")
	return strs[len(strs)-1]
}

func getTypeName(fieldType reflect.Type, fieldName string) string {
	switch fieldType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fieldName = "integer"
	case reflect.Array, reflect.Slice:
		fieldName = "array"
	}

	return fieldName
}

func getNumberName(fieldType reflect.Type) string {
	if fieldType.Kind() == reflect.Int || fieldType.Kind() == reflect.Int32 {
		return "int32"
	} else if fieldType.Kind() == reflect.Int8 {
		return "int8"
	} else if fieldType.Kind() == reflect.Int16 {
		return "int16"
	} else if fieldType.Kind() == reflect.Int64 {
		return "int64"
	}
	return "int32"
}

func getTypeInStruct[T any](input T) map[string]any {
	properties := map[string]any{}

	inputType := reflect.TypeOf(input)
	for i := 0; i < inputType.NumField(); i++ {
		property := map[string]any{}
		field := inputType.Field(i)
		tag := field.Tag.Get("json")

		fieldType := inputType.Field(i).Type
		if field.Type.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		typeString := getTypeName(fieldType, fieldType.String())
		isArray := fieldType.Kind() == reflect.Slice || fieldType.Kind() == reflect.Array
		isStruct := fieldType.Kind() == reflect.Struct

		if isStruct {
			property["$ref"] = fmt.Sprintf("#/definitions/%s", fieldType.Name())
		} else if isArray {
			items := map[string]any{}
			if fieldType.Elem().Kind() == reflect.Struct {
				items["$ref"] = fmt.Sprintf("#/definitions/%s", fieldType.Elem().Name())
			} else {
				items["type"] = getTypeName(fieldType.Elem(), fieldType.Elem().String())
			}
			if items["type"] == "integer" {
				items["format"] = getNumberName(fieldType.Elem())
			}

			property["type"] = "array"
			property["items"] = items
		} else {
			property["type"] = typeString
		}

		if property["type"] == "integer" {
			property["format"] = getNumberName(fieldType)
		}

		properties[tag] = property

	}

	return map[string]any{"type": "object", "properties": properties, "xml": map[string]any{"name": inputType.Name()}}
}

func getAllTypes[T any](input T) (map[string]any, string) {
	outputs := map[string]any{}
	inputType := reflect.TypeOf(input)
	outputs[inputType.Name()] = getTypeInStruct(input)
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		fieldType := field.Type
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}
		if fieldType.Kind() == reflect.Struct {
			internalOutput, _ := getAllTypes(reflect.ValueOf(input).Field(i).Interface())
			for key, value := range internalOutput {
				outputs[key] = value
			}
		}
		if fieldType.Kind() == reflect.Array || fieldType.Kind() == reflect.Slice {
			if fieldType.Elem().Kind() == reflect.Struct {
				internalOutput, _ := getAllTypes(reflect.New(fieldType.Elem()).Elem().Interface())
				for key, value := range internalOutput {
					outputs[key] = value
				}
			}
		}
	}

	return outputs, inputType.Name()
}

func defineHandler[ReqT, ResT any](method, path string, handler context.Handler, request *ReqT, response *ResT, settings ...Setting) {
	var setting Setting = defaultSetting
	if len(settings) != 0 {
		setting = settings[0]
	}

	var reqDefinitions map[string]any
	mainReqStructKey := ""
	if request != nil {
		reqDefinitions, mainReqStructKey = getAllTypes(*request)
		for key, definition := range reqDefinitions {
			allDefinitions[key] = definition
		}
	}

	var resDefinitions map[string]any
	mainResStructKey := ""
	if response != nil {
		resDefinitions, mainResStructKey = getAllTypes(*response)
		for key, definition := range resDefinitions {
			allDefinitions[key] = definition
		}
	}

	route := map[string]any{}
	tag := strings.TrimSpace(strings.Replace(goStrings.StartCase(getPackageName(handler)), "Handlers", "", 1))
	if tag == "" && setting.Tag == "" {
		tag = defaultSetting.Tag
	} else if tag == "" {
		tag = setting.Tag
	}
	route["tags"] = []string{tag}
	if setting.Description != "" {
		route["description"] = setting.Description
	}
	if setting.Summary != "" {
		route["summary"] = setting.Summary
	}
	if setting.Authorization {
		route["security"] = []any{map[string]any{"Authorization": []any{}}}
	}
	if request != nil {
		route["parameters"] = []any{map[string]any{
			"in":       "body",
			"name":     "body",
			"required": (method == "post" || method == "patch"),
			"schema":   map[string]any{"$ref": fmt.Sprintf("#/definitions/%s", mainReqStructKey)},
		}}
	}
	route["produces"] = []string{"application/json"}
	route["responses"] = map[string]any{
		"200": map[string]any{
			"schema": map[string]any{"$ref": fmt.Sprintf("#/definitions/%s", mainResStructKey)},
		},
	}

	allRoutes[filepath.Join(PreRoute, path)] = map[string]any{method: route}
}

func Get[ResT any](app router.Party, path string, handler []context.Handler, response *ResT, settings ...Setting) {
	if !Generate {
		app.Get(path, handler...)
		return
	}

	defineHandler[any]("get", path, handler[len(handler)-1], nil, response, settings...)
}

func Delete(app router.Party, path string, handler []context.Handler, settings ...Setting) {
	if !Generate {
		app.Delete(path, handler...)
		return
	}

	defineHandler[any, any]("delete", path, handler[len(handler)-1], nil, nil, settings...)
}

func Post[ReqT, ResT any](app router.Party, path string, handler []context.Handler, request *ReqT, response *ResT, settings ...Setting) {
	if !Generate {
		app.Post(path, handler...)
		return
	}

	defineHandler("post", path, handler[len(handler)-1], request, response, settings...)
}

func Put[ReqT, ResT any](app router.Party, path string, handler []context.Handler, request *ReqT, response *ResT, settings ...Setting) {
	if !Generate {
		app.Put(path, handler...)
		return
	}

	defineHandler("put", path, handler[len(handler)-1], request, response, settings...)
}

func Patch[ReqT, ResT any](app router.Party, path string, handler []context.Handler, request *ReqT, response *ResT, settings ...Setting) {
	if !Generate {
		app.Patch(path, handler...)
		return
	}

	defineHandler("patch", path, handler[len(handler)-1], request, response, settings...)
}

func Run(app *iris.Application, ip, port, address string) *iris.Application {
	if Generate {
		data := map[string]any{}
		swaggerJsonAddress := filepath.Join(address, "swagger.json")
		swaggerJson, _ := os.ReadFile(swaggerJsonAddress)
		json.Unmarshal(swaggerJson, &data)
		data["definitions"] = allDefinitions
		data["paths"] = allRoutes
		data["schemes"] = []string{"http", "https"}
		data["securityDefinitions"] = map[string]any{"Authorization": map[string]any{"in": "header", "name": "Authorization", "type": "apiKey"}}
		dataBytes, _ := json.MarshalIndent(data, "", "	")
		file, _ := os.OpenFile(swaggerJsonAddress, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		defer file.Close()
		file.Write(dataBytes)
		return app
	}

	app.Listen(fmt.Sprintf("%s:%s", ip, port))
	return nil
}
