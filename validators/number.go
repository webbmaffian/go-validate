package validators

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/webbmaffian/go-validate/utils"
)

func Integer(value, parent reflect.Value, arg string, opt *utils.Options) (err error) {
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if str, ok := value.Interface().(string); ok {
		if _, intErr := strconv.ParseInt(str, 10, 64); intErr == nil {
			return
		}
	}

	return errors.New("Must be an integer")
}
