package validators

import (
	"errors"
	"reflect"
	"strings"

	"github.com/webbmaffian/go-validate/utils"
)

var oneofCache = make(map[string][]string)

func getOneof(str string) (oneof []string) {
	var exists bool

	oneof, exists = oneofCache[str]

	if !exists {
		oneof = strings.Split(str, " ")
		oneofCache[str] = oneof
	}

	return
}

func Oneof(value, parent reflect.Value, arg string, flags utils.Flags) (err error) {
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	switch v := value.Interface().(type) {

	case string:
		oneof := getOneof(arg)

		for _, s := range oneof {
			if s == v {
				return
			}
		}

		return errors.New("Must be one of: " + strings.Join(oneof, ", "))

	case []string:
		oneof := getOneof(arg)

	outer:
		for _, s := range oneof {
			for _, str := range v {
				if s == str {
					continue outer
				}

				return errors.New("Must be one of: " + strings.Join(oneof, ", "))
			}
		}

		return
	}

	return errors.New("Must be a string")
}
