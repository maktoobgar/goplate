package copier

import (
	"database/sql"
	"errors"
	"reflect"
	"time"

	"github.com/kataras/iris/v12"
)

func StructCheck(input any) {
	inputVar := reflect.ValueOf(input)
	if inputVar.Kind() == reflect.Ptr && inputVar.Elem().Kind() == reflect.Struct {
		return
	}
	panic(errors.New("copy strcuc check: input type is not a pointer to a struct"))
}

func Copy[T, T2 any](to *T, from *T2) T {
	StructCheck(to)
	StructCheck(from)

	fromValue := reflect.ValueOf(from).Elem()
	toValue := reflect.ValueOf(to).Elem()
	if fromValue := reflect.ValueOf(from); fromValue.IsValid() && fromValue.Elem().Kind() == reflect.Struct {
		if method := fromValue.MethodByName("Reformat"); method.IsValid() {
			method.Call([]reflect.Value{})
		}
	}

	for _, field := range reflect.VisibleFields(fromValue.Type()) {
		if field.IsExported() {
			toFieldValue := toValue.FieldByName(field.Name)
			fromFieldValue := fromValue.FieldByName(field.Name)
			for fromFieldValue.Kind() == reflect.Ptr {
				fromFieldValue = fromFieldValue.Elem()
			}
			for toFieldValue.Kind() == reflect.Ptr {
				toFieldValue = toFieldValue.Elem()
			}
			if !toFieldValue.IsValid() || !fromFieldValue.IsValid() {
				continue
			}
			toFieldValueType := toFieldValue.Type()
			fromFieldValueType := fromFieldValue.Type()
			if toFieldValue.IsValid() {
				if fromFieldValueType.String() == "time.Time" && toFieldValue.Kind() == reflect.Int64 {
					t := fromFieldValue.Interface().(time.Time)
					toFieldValue.Set(reflect.ValueOf(t.Unix()))
				} else if fromFieldValue.Kind() == reflect.Int64 && toFieldValueType.String() == "time.Time" {
					toFieldValue.Set(reflect.ValueOf(time.Unix(fromFieldValue.Int(), 0)))
				} else if toFieldValueType.String() == "sql.NullString" && fromFieldValueType.String() == "string" {
					toFieldValue.Set(reflect.ValueOf(sql.NullString{String: fromFieldValue.String(), Valid: true}))
				} else if toFieldValueType.String() == "sql.NullInt32" && fromFieldValueType.String() == "int32" {
					toFieldValue.Set(reflect.ValueOf(sql.NullInt32{Int32: int32(fromFieldValue.Int()), Valid: true}))
				} else if toFieldValueType.String() == "sql.NullTime" && fromFieldValueType.String() == "time.Time" {
					toFieldValue.Set(reflect.ValueOf(sql.NullTime{Time: fromFieldValue.Interface().(time.Time), Valid: true}))
				} else if toFieldValueType.String() == "sql.NullBool" && fromFieldValueType.String() == "bool" {
					toFieldValue.Set(reflect.ValueOf(sql.NullBool{Bool: fromFieldValue.Interface().(bool), Valid: true}))
				} else if fromFieldValueType.String() == "sql.NullString" && toFieldValueType.String() == "string" {
					toFieldValue.Set(reflect.ValueOf(fromFieldValue.Interface().(sql.NullString).String))
				} else if fromFieldValueType.String() == "sql.NullInt32" && toFieldValueType.String() == "int32" {
					toFieldValue.Set(reflect.ValueOf(fromFieldValue.Interface().(sql.NullInt32).Int32))
				} else if fromFieldValueType.String() == "sql.NullTime" && toFieldValueType.String() == "time.Time" {
					toFieldValue.Set(reflect.ValueOf(fromFieldValue.Interface().(sql.NullTime).Time))
				} else if fromFieldValueType.String() == "sql.NullBool" && toFieldValueType.String() == "bool" {
					toFieldValue.Set(reflect.ValueOf(fromFieldValue.Interface().(sql.NullBool).Bool))
				} else {
					toFieldValue.Set(fromFieldValue)
				}
			}
		}
	}

	return toValue.Interface().(T)
}

func ArrayCopy[T, T2 any](to []T, from []T2) []T {
	if len(to) != len(from) {
		panic("length of both 'from' and 'to' has to be the same")
	}

	for i := range from {
		Copy(&to[i], &from[i])
	}
	return to
}

func CastParams[T any](ctx iris.Context, params *T) {
	paramsValue := reflect.ValueOf(params).Elem()
	for i := 0; i < paramsValue.NumField(); i++ {
		field := paramsValue.Field(i)
		fieldType := field.Type()
		exactFieldType := fieldType
		fieldStructType := paramsValue.Type().Field(i)
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}
		var value any = nil
		if tag, ok := fieldStructType.Tag.Lookup("json"); ok && tag != "" {
			if !ctx.URLParamExists(tag) {
				continue
			}
			switch fieldType.Kind() {
			case reflect.Bool:
				value, _ = ctx.URLParamBool(tag)
			case reflect.Float64:
				value, _ = ctx.URLParamFloat64(tag)
			case reflect.Float32:
				value, _ = ctx.URLParamFloat64(tag)
				value = float32(value.(float64))
			case reflect.Int:
				value, _ = ctx.URLParamInt(tag)
			case reflect.Int8:
				value, _ = ctx.URLParamInt(tag)
				value = int8(value.(int))
			case reflect.Int16:
				value, _ = ctx.URLParamInt(tag)
				value = int16(value.(int))
			case reflect.Int32:
				value, _ = ctx.URLParamInt(tag)
				value = int32(value.(int))
			case reflect.Int64:
				value, _ = ctx.URLParamInt(tag)
				value = int64(value.(int))
			case reflect.String:
				value = ctx.URLParam(tag)
			}
		}
		if exactFieldType.Kind() == reflect.Ptr {
			valueValue := reflect.New(reflect.TypeOf(value))
			valueValue.Elem().Set(reflect.ValueOf(value))
			field.Set(valueValue)
		} else {
			field.Set(reflect.ValueOf(value))
		}
	}
}
