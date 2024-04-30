package copier

import (
	"database/sql"
	"errors"
	"reflect"
	"time"
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
			if !toFieldValue.IsValid() {
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
					toFieldValue.Set(reflect.ValueOf(sql.NullString{String: fromFieldValue.String(), Valid: len(fromFieldValue.String()) > 0}))
				} else if toFieldValueType.String() == "sql.NullInt32" && fromFieldValueType.String() == "int32" {
					toFieldValue.Set(reflect.ValueOf(sql.NullInt32{Int32: int32(fromFieldValue.Int()), Valid: fromFieldValue.Int() != 0}))
				} else if toFieldValueType.String() == "sql.NullTime" && fromFieldValueType.String() == "time.Time" {
					toFieldValue.Set(reflect.ValueOf(sql.NullTime{Time: fromFieldValue.Interface().(time.Time), Valid: !fromFieldValue.Interface().(time.Time).IsZero()}))
				} else if fromFieldValueType.String() == "sql.NullString" && toFieldValueType.String() == "string" {
					toFieldValue.Set(reflect.ValueOf(fromFieldValue.Interface().(sql.NullString).String))
				} else if fromFieldValueType.String() == "sql.NullInt32" && toFieldValueType.String() == "int32" {
					toFieldValue.Set(reflect.ValueOf(fromFieldValue.Interface().(sql.NullInt32).Int32))
				} else if fromFieldValueType.String() == "sql.NullTime" && toFieldValueType.String() == "time.Time" {
					toFieldValue.Set(reflect.ValueOf(fromFieldValue.Interface().(sql.NullTime).Time))
				} else {
					toFieldValue.Set(fromFieldValue)
				}
			}
		}
	}

	return toValue.Interface().(T)
}
