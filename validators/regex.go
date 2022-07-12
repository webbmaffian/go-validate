package validators

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/webbmaffian/go-validate/utils"
)

var regexCache = make(map[string]*regexp.Regexp)

func getRegex(pattern string) (regex *regexp.Regexp, err error) {
	var exists bool

	regex, exists = regexCache[pattern]

	if !exists {
		regex, err = regexp.Compile(pattern)
		regexCache[pattern] = regex
	}

	return
}

func Regex(value, parent reflect.Value, arg string, flags utils.Flags) (err error) {
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	if str, ok := value.Interface().(string); ok {
		var regex *regexp.Regexp

		regex, err = getRegex(arg)

		if err != nil {
			return
		}

		if regex.MatchString(str) {
			return
		}
	}

	return errors.New("Must match the regex pattern: " + arg)
}
