package validators

import (
	"errors"
	"reflect"

	"github.com/webbmaffian/go-validate/utils"
)

func Required(value, parent reflect.Value, arg string, opt *utils.Options) (err error) {
	if value.IsValid() {
		if !value.IsZero() {
			return
		}
	} else if opt.SkipNil {
		return
	}

	return errors.New("Required field")
}
