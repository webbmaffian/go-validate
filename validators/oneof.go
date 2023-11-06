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

func Oneof(value, parent reflect.Value, arg string, opt *utils.Options) (err error) {
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if v, ok := utils.String(value); ok {
		if v == "" {
			return
		}

		oneof := getOneof(arg)

		for _, s := range oneof {
			if s == v {
				return
			}
		}

		return errors.New("Must be one of: " + strings.Join(oneof, ", "))
	}

	if v, ok := utils.StringArray(value); ok {
		oneof := getOneof(arg)

	outer:
		for _, s := range v {
			if s == "" {
				continue outer
			}

			for _, str := range oneof {
				if s == str {
					continue outer
				}
			}

			return errors.New("Must be one of: " + strings.Join(oneof, ", "))
		}

		return
	}

	return nil
}
