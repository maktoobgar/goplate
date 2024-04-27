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

func Copy(to, from any) any {
	StructCheck(to)
	StructCheck(from)

	fromValue := reflect.ValueOf(from).Elem()
	toValue := reflect.ValueOf(to).Elem()
	for _, field := range reflect.VisibleFields(fromValue.Type()) {
		if field.IsExported() {
			fieldValue := toValue.FieldByName(field.Name)
			fromFielValue := fromValue.FieldByName(field.Name)
			for fromFielValue.Kind() == reflect.Ptr {
				fromFielValue = fromFielValue.Elem()
			}
			for fieldValue.Kind() == reflect.Ptr {
				fieldValue = fieldValue.Elem()
			}
			if fieldValue.IsValid() {
				if fromFielValue.Type().String() == "time.Time" && fieldValue.Kind() == reflect.Int64 {
					t := fromFielValue.Interface().(time.Time)
					fieldValue.Set(reflect.ValueOf(t.Unix()))
				} else if fieldValue.Type().String() == "sql.NullString" && fromFielValue.Type().String() == "string" {
					fieldValue.Set(reflect.ValueOf(sql.NullString{String: fromFielValue.String(), Valid: len(fromFielValue.String()) > 0}))
				} else if fromFielValue.Kind() == reflect.Int64 && fieldValue.Type().String() == "time.Time" {
					fieldValue.Set(reflect.ValueOf(time.Unix(fromFielValue.Int(), 0)))
				} else {
					fieldValue.Set(fromFielValue)
				}
			}
		}
	}

	return toValue.Interface()
}
