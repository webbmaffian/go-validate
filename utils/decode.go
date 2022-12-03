package utils

import (
	"database/sql/driver"
	"reflect"
)

type stringer interface {
	String() string
}

func String(value reflect.Value) (str string, ok bool) {
	if !value.CanInterface() {
		return
	}

	return decodeString(value.Interface())
}

func decodeString(val any) (str string, ok bool) {
	switch v := val.(type) {

	case string:
		return v, true

	case stringer:
		return v.String(), true

	case driver.Valuer:
		if value, err := v.Value(); err == nil {
			return decodeString(value)
		}

	}

	return "", false
}
