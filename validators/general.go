package validators

import (
	"errors"
	"reflect"

	"github.com/webbmaffian/go-validate/utils"
)

func Required(value, parent reflect.Value, arg string, flags utils.Flags) (err error) {
	if !value.IsValid() || value.IsZero() {
		err = errors.New("Required field")
	}

	return
}
