package i18nInterfaces

import "reflect"

type TranslatorI interface {
	Galidator() TranslatorGalidatorI
	StatusCodes() TranslatorStatusCodesI
	Translate(key string, optionalInputs ...[]any) string
}

type TranslatorGalidatorI interface {
	Example() string
	Translate(key string, optionalInputs ...[]any) string
}

type TranslatorStatusCodesI interface {
	InternalServerError() string
	PageNotFound() string
	Translate(key string, optionalInputs ...[]any) string
}

func translate(instance any, key string, optionalInputs ...[]any) string {
	structType := reflect.TypeOf(instance)
	inputs := []any{}
	if len(optionalInputs) > 0 {
		inputs = optionalInputs[0]
	}

	// Iterate over all methods of the struct
	for i := 0; i < structType.NumMethod(); i++ {
		method := structType.Method(i)
		if method.Name == key {
			reflectValues := make([]reflect.Value, len(inputs))
			for i, v := range inputs {
				reflectValues[i] = reflect.ValueOf(v)
			}
			return method.Func.Call(reflectValues)[0].String()
		}
	}
	return key
}