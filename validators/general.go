package validators

import (
	"errors"
	"reflect"

	"github.com/webbmaffian/go-validate/utils"
)

func Required(value, parent reflect.Value, arg string, opt *utils.Options) (err error) {
	if opt.SkipNil {
		return
	}

	if value.IsValid() && !value.IsZero() {
		return
	}

	return errors.New("Required field")
}

func Forbidden(value, parent reflect.Value, arg string, opt *utils.Options) (err error) {
	if !value.IsValid() || value.IsZero() {
		return
	}

	return errors.New("Forbidden field")
}
