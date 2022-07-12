package validate

import (
	"reflect"

	"github.com/webbmaffian/go-validate/utils"
)

type ValueGetter func(originalValue reflect.Value, opt *utils.Options) any

var registeredValueGetters map[string]ValueGetter

func init() {
	registeredValueGetters = map[string]ValueGetter{}
}

func RegisterValueGetter(fn ValueGetter, types ...any) {
	for _, t := range types {
		registeredValueGetters[reflect.TypeOf(t).Name()] = fn
	}
}

func getValue(originalValue reflect.Value, opt *utils.Options) reflect.Value {
	if valueGetter, exists := registeredValueGetters[originalValue.Type().Name()]; exists {
		res := valueGetter(originalValue, opt)

		if reflVal, ok := res.(reflect.Value); ok {
			return reflVal
		}

		return reflect.ValueOf(res)
	}

	return originalValue
}
