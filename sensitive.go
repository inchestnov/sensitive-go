package sensitive

import (
	"errors"
	"reflect"
)

type Sensitive[T any] any

func Detach[T any](source T) (T, Sensitive[T], error) {
	var s Sensitive[T] = source

	var isPtr bool
	value := reflect.ValueOf(source)
	if value.Kind() == reflect.Ptr {
		isPtr = true

		cp := reflect.New(value.Elem().Type()).Elem()
		cp.Set(value.Elem())
		s = cp.Addr().Interface().(Sensitive[T])

		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return source, s, errors.New("only structs supported")
	}

	if !isPtr {
		// Access exported fields.
		value = reflect.ValueOf(&source).Elem()
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		if !field.IsExported() {
			continue
		}

		sensitiveTag := field.Tag.Get("sensitive")
		if sensitiveTag != "true" {
			continue
		}

		value.Field(i).Set(reflect.Zero(field.Type))
	}

	return source, s, nil
}

func Attach[T any](_ T, sensitive Sensitive[T]) (T, error) {
	return sensitive.(T), nil
}
