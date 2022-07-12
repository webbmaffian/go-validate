package validate

import (
	"reflect"

	"github.com/webbmaffian/go-validate/utils"
	"github.com/webbmaffian/go-validate/validators"
)

type Validator func(value reflect.Value, parent reflect.Value, arg string, flags utils.Flags) error

type registeredValidator struct {
	validator  Validator
	acceptZero bool
}

var registeredValidators map[string]registeredValidator

func init() {
	registeredValidators = map[string]registeredValidator{
		"required": {validators.Required, true},
		"integer":  {validators.Integer, false},
		"regex":    {validators.Regex, false},
		"email":    {validators.Email, false},
		"oneof":    {validators.Oneof, false},
	}
}

func RegisterValidator(tag string, fn Validator, acceptZero ...bool) {
	registeredValidators[tag] = registeredValidator{fn, len(acceptZero) > 0 && acceptZero[0]}
}
