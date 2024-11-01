package essential

import (
	"errors"
	"reflect"
	"strings"
)

func IsStructFilled[T any](s T) error {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}

	if v.Kind() != reflect.Struct {
		return errors.New("struct expected, but got: " + v.Kind().String())
	}

	empty := []string{}

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsZero() {
			empty = append(empty, v.Type().Field(i).Name)
		}
	}

	if len(empty) > 0 {
		return errors.New("missing field: " + strings.Join(empty, ", "))
	}

	return nil
}
