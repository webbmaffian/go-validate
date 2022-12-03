package validators

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/webbmaffian/go-validate/utils"
)

var emailRegex *regexp.Regexp

func init() {
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
}

func Email(value, parent reflect.Value, arg string, opt *utils.Options) (err error) {
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if str, ok := utils.String(value); ok {
		if emailRegex.MatchString(str) {
			return
		}
	}

	return errors.New("Must be a valid email address")
}
